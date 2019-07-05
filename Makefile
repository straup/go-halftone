fmt:
	go fmt ./...

tools:
	go build -o bin/crop cmd/crop/main.go
	go build -o bin/halftone cmd/halftone/main.go
	go build -o bin/resize cmd/resize/main.go
	go build -o bin/transparency cmd/transparency/main.go
	go build -o bin/empty cmd/empty/main.go
