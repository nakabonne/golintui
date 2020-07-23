FROM golang:1.14

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.29.0

COPY golintui /usr/bin/
CMD ["golintui"]
