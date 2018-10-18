FROM golang:1.11

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/pressly-goose-test
COPY . .

RUN dep ensure
RUN go build
CMD ["./pressly-goose-test"]
