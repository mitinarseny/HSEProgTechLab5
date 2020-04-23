.PHONY: run
run:
	go run ./main.go

.PHONY: reports
reports: $(addprefix report.,json csv)
	cat $<
	cat $(word 2, $^)
