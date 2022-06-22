# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /main .

EXPOSE 8082

CMD ["./main"]