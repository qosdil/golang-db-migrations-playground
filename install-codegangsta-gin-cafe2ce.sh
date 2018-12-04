echo "Installing codegangsta/gin..."
go get -v github.com/codegangsta/gin
cd $GOPATH/src/github.com/codegangsta/gin
git reset --hard cafe2ce
go install -i .
echo "Installing codegangsta/gin version \"cafe2ce\" was successful."
