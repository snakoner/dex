FROM golang:1.23

RUN go version
# ENV GOPATH=/
RUN mkdir /build
WORKDIR /build
COPY . .

RUN go mod download
RUN go build -o app ./cmd/app/*

EXPOSE 8000

CMD ["./app"]