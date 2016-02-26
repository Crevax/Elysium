FROM golang:1.6.0

WORKDIR /go/src/cjdavis.me/elysium
Add . /go/src/cjdavis.me/elysium

RUN go get github.com/tools/godep

RUN godep restore

RUN go install cjdavis.me/elysium

EXPOSE 8080
ENTRYPOINT ["/go/bin/elysium"]
