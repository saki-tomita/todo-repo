FROM golang:1.19.1-alpine

RUN apk add --update --no-cache ca-certificates git

COPY ./todos /go/todos
COPY go.mod /go/todos
WORKDIR /go/todos

RUN go mod download
COPY . .

EXPOSE 8080
CMD ["go", "run", "main.go"]