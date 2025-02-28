FROM golang:1.24

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN  ls
RUN go mod tidy

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
