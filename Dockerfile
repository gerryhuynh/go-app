FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -v -o /go-app

# from absolutely nothing
# not normal, usually do from alpine/ubuntu; bc we need certificates
# FROM scratch
FROM ubuntu:latest

# just copy the binary; not everything else; because we run off the binary
# think go build
# then ./go-app

# ca certificates is bundle of valid authorities on internet; valid HTTPS certificates
# later try download from public site with from scratch vs. from ubuntu
# multistage
# first stage is builder for binary
# second stage is runtime (output)
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=0 /go-app /go-app

EXPOSE 50051 8080

ENTRYPOINT ["/go-app"]
CMD []

