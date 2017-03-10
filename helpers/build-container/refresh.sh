docker-compose down

cd ~/src/github.com/rockaut/check_3par
git pull
docker run --rm -v /root/bin:/go/bin -v /root/src:/go/src -w /go/src/github.com/rockaut/check_3par golang:latest go build -buildmode=c-shared -v -o /go/bin/check_3par.so

cp ~/bin/check_3par.so /graph/volumes/zabbixcontainer_zabbix_agent_modules/_data/
cp ~/bin/check_3par.so /graph/volumes/zabbixcontainer_zabbix_modules/_data/

cd -

docker-compose up -d