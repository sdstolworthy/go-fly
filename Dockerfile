FROM golang

ADD . /go/src/github.com/sdstolworthy/go-fly

RUN cd /go/src/github.com/sdstolworthy/go-fly \
  && go get -v ./...

RUN cd /go/src/github.com/sdstolworthy/go-fly && go build -o gofly *.go

ENTRYPOINT /go/src/github.com/sdstolworthy/go-fly/gofly

EXPOSE 8080