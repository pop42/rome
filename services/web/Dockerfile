FROM golang:1.7.3
WORKDIR /go/src/github.com/bluefoxcode/rome/services/web
COPY . /go/src/github.com/bluefoxcode/rome/
RUN go install
ENTRYPOINT ["web"]