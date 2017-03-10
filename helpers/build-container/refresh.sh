cd ~/src/github.com/rockaut/check_3par
git pull
docker run --rm -v /root/bin:/go/bin -v /root/src:/go/src -w /go/src/github.com/rockaut/check_3par golang:latest go build -buildmode=c-shared -v -o /go/bin/check_3par.so

cd -