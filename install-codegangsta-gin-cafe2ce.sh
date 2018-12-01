echo "Installing codegangsta/gin..."
dir=$GOPATH/src/github.com/codegangsta/gin
rm -rf $dir
mkdir -p $dir
cd $dir
git clone git@github.com:codegangsta/gin.git .
git reset --hard cafe2ce
go install -i .
echo "Installing codegangsta/gin version \"cafe2ce\" was successful."
