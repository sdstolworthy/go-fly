FROM golang

ADD . /go/src/github.com/sdstolworthy/go-fly

RUN cd /go/src/github.com/sdstolworthy/go-fly \
  && go get -v ./...

RUN go install github.com/sdstolworthy/go-fly

ENTRYPOINT /go/bin/go-fly

EXPOSE 8080