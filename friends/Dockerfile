FROM golang:latest

ADD . /go/src/github.com/BearCloud/fa20-project-ujjy/friends

WORKDIR /go/src/github.com/BearCloud/fa20-project-ujjy/friends

RUN go mod download

RUN go build -o main .

EXPOSE 80

ENTRYPOINT [ "./main" ]
