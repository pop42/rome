FROM golang:1.7.3
WORKDIR /go/src/github.com/bluefoxcode/rome/services/Notekeeper
COPY . /go/src/github.com/bluefoxcode/rome/
RUN go install
ENTRYPOINT ["Notekeeper"]