FROM golang:1.11 AS builder

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/pressly-goose-test
COPY . .

RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pressly-goose-test .

FROM scratch

WORKDIR /root
COPY --from=builder /go/src/pressly-goose-test/pressly-goose-test .
COPY --from=builder /go/src/pressly-goose-test/config.json .
CMD ["./pressly-goose-test"]
