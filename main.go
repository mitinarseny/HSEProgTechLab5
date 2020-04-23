package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mitinarseny/HSEProgTechLab5/report"
)

func main() {
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-exitCh
		cancel()
	}()

	jr, err := report.NewReporter(report.JSONReporter)
	defer jr.Close()
	if err != nil {
		log.Fatalf("unable to start json reporter: %s", err)
	}
	cr, err := report.NewReporter(report.CSVReporter)
	defer cr.Close()
	if err != nil {
		log.Fatalf("unable to start csv reporter: %s", err)
	}
	events := startEvents(ctx)
	switch err := runReporters(ctx, events, jr, cr); err {
	case context.Canceled:
		log.Println("context cancelled")
	default:
		log.Fatalf("error while reporting: %s", err)
	}
}

func runReporters(ctx context.Context, events <-chan report.Event, reporters ...report.Reporter) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e := <-events:
			for _, r := range reporters {
				if err := r.Report(&e); err != nil {
					return err
				}
			}
		}
	}
}

func startEvents(ctx context.Context) <-chan report.Event {
	events := make(chan report.Event)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				n++
				fmt.Println(t)
				events <- report.Event{
					ID:        n,
					Timestamp: t,
				}
			}
		}
	}()
	return events
}
