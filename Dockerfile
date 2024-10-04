FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -v -o /go-app

FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates
COPY --from=0 /go-app /go-app

EXPOSE 50051 8080

ENTRYPOINT ["/go-app"]
