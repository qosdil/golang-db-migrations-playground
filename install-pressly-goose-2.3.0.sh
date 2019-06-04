# Prerequisites
# - $GOPATH set up properly
# - curl
# - Git
# - Connection to github.com
# - Go (runs well on Go v1.11)
#
# Run: sh ./[this file]

# Set up Go environment
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg
mkdir -p $GOPATH/src

# Install Dep
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install goose v2.3.0
mkdir -p $GOPATH/src/github.com/pressly/goose
cd $GOPATH/src/github.com/pressly/goose
git clone git@github.com:pressly/goose.git .
git checkout tags/v2.3.0
dep ensure -v

# Create goose binary
go install -i ./cmd/goose
