mkdir ~/src

git clone -b patch-1 https://github.com/rockaut/g2z.git ~/src/github.com/rockaut/g2z

docker volume create zabbixcontainer_zabbix_modules
docker volume create zabbixcontainer_zabbix_agent_modules

docker run --rm -v /root/bin:/go/bin -v /root/src:/go/src -w /go/src/github.com/rockaut/check_3par golang:latest go build -buildmode=c-shared -v -o /go/bin/check_3par.so

cp ~/bin/check_3par.so /graph/volumes/zabbixcontainer_zabbix_agent_modules/_data/
cp ~/bin/check_3par.so /graph/volumes/zabbixcontainer_zabbix_modules/_data/

docker run --rm -v /root/bin:/go/bin -v /root/src:/go/src -w /go/src/github.com/rockaut/check_3par/dummy golang:latest go build -buildmode=c-shared -v -o /go/bin/dummy.so

cp ~/bin/dummy.so /graph/volumes/zabbixcontainer_zabbix_agent_modules/_data/
cp ~/bin/dummy.so /graph/volumes/zabbixcontainer_zabbix_modules/_data/
