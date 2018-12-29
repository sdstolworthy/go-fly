FROM golang

ADD . /go/src/github.com/sdstolworthy/go-fly

RUN /bin/sh -c "cd /go/src/github.com/sdstolworthy && go get ./... && cd seed && go run seed/*.go"

RUN go install github.com/sdstolworthy/go-fly

ENTRYPOINT env PORT=8080 /go/bin/go-fly

EXPOSE 8080