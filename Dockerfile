FROM golang:1.21.1

LABEL maintainer=""
LABEL version="1.0"
LABEL description="Docker image for building and running the go app"

WORKDIR /app
COPY . /app

RUN go get
RUN go build -o server main.go

CMD ["./server"]
ENV TZ=Asia/Jakarta
EXPOSE 3000