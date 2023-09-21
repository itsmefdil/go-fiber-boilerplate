build:
	go build -o server main.go

run: build
	./server

watch:
	go run main.go