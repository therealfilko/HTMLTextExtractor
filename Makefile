default:
	go build -o extractor cmd/main.go

run:
	go run cmd/main.go

clean:
	rm extractor
