default:
	go build -o extractor cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./tests/... -v

clean:
	rm extractor
