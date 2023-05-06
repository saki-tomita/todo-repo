FROM golang:1.19.1-alpine

RUN apk add --update --no-cache ca-certificates git

COPY ./todos /go/todos
COPY ./todos/go.mod /go/todos
WORKDIR /go/todos

RUN go mod download
COPY ./build .

RUN apk --no-cache add curl

EXPOSE 8081

CMD ["go", "run", "main.go"]