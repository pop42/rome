FROM golang:1.7.3
WORKDIR /go/src/github.com/bluefoxcode/rome/services/Caesar
COPY . /go/src/github.com/bluefoxcode/rome/
RUN go install
ENTRYPOINT ["Caesar"]
