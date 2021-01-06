build:
	go build ./...

test:
	go test ./...

coverage:
	go test -coverprofile fmtcoverage.html fmt