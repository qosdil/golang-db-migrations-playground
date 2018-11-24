go get -v github.com/pressly/goose/cmd/goose
cd $GOPATH/src/github.com/pressly/goose
git checkout tags/v2.3.0
go install -i ./cmd/goose