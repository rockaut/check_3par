mkdir ~/src

git clone -b patch-1 https://github.com/rockaut/g2z.git ~/src/github.com/rockaut/g2z

docker run --rm -v "$PWD"/src:/go/src -w /go/src/github.com/rockaut/check_3par golang:latest go build -v
