fmt:
	go fmt cmd/*.go
	go fmt flags/*.go
	go fmt halftone/*.go
	go fmt imaging/*.go
	go fmt pixel/*.go
	go fmt resize/*.go
	go fmt util/*.go

tools:
	go build -o bin/crop cmd/crop/main.go
	go build -o bin/halftone cmd/halftone/main.go
	go build -o bin/resize cmd/resize/main.go
	go build -o bin/transparency cmd/transparency/main.go
