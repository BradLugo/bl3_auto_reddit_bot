FROM golang

COPY . /go/src/github.com/matt1484/bl3_auto_vip
WORKDIR /go/src/github.com/matt1484/bl3_auto_vip

RUN go mod download && go mod verify
RUN go build -o ./main cmd/main.go && chmod +x ./main

CMD ./main