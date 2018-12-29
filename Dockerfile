FROM golang

ADD . /go/src/github.com/sdstolworthy/go-fly

RUN /bin/sh -c "cd /go/src/github.com/sdstolworthy && go get ./..."

RUN go install github.com/sdstolworthy/go-fly

ENTRYPOINT /go/bin/go-fly

EXPOSE 8080