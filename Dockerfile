FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o advert ./cmd/main.go

CMD ["./advert"]