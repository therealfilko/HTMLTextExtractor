default:
	go build -o extractor cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./tests/... -v

test-count:
	go test ./tests/count_words_test.go -v

test-fetch-save:
	go test ./tests/fetch_and_save_test.go -v

clean:
	rm extractor
