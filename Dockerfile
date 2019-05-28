# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.12

WORKDIR /go/src/github.com/demo

COPY . .

RUN chmod a+x docker-entrypoint.sh
RUN go install .

ENTRYPOINT ["./docker-entrypoint.sh"]