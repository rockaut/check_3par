mkdir ~/src
git clone https://github.com/rockaut/g2z.git ~/src/github.com/rockaut/g2z

docker run --rm -v "$PWD"/src:/src -w /src/github.com/rockaut/check_3par golang:latest go build -v
