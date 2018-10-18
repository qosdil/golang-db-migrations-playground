FROM golang:1.11

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/pressly-goose-test
COPY . .

RUN dep ensure
RUN go build
RUN chmod +x ./pressly-goose-test
CMD ["./pressly-goose-test"]
