FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates
COPY ./out/go-app /go-app

EXPOSE 50051 8080

ENTRYPOINT ["/go-app"]
