build:
	go build -o server main.go

run: build
	./server

swag:
	swag init
	
watch:
	go run main.go