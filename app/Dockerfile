FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

RUN go get github.com/go-redis/redis

ADD . /go/src/myapp
WORKDIR /go/src/myapp
COPY . .

RUN go get myapp
RUN go install -v
CMD ["go", "run","url_working.go"]
ENTRYPOINT ["/go/bin/myapp"]
