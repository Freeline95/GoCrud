FROM golang:1.16-alpine as builder
COPY go.mod go.sum /go/src/github.com/Freeline95/GoCrud/
WORKDIR /go/src/github.com/Freeline95/GoCrud

RUN go mod download

COPY . /go/src/github.com/Freeline95/GoCrud

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/app/GoCrud github.com/Freeline95/GoCrud

ENTRYPOINT ["cmd/app/GoCrud"]