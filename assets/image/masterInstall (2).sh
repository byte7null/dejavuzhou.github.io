#!/usr/bin/env bash
if [ $# -lt 1 ] ; then
  echo "第一个参数时thanos打包tar.gz服务器的host eg: sudo ./masterInstall.sh dev.thanos.com:8081"
  exit 1;
fi
pkill -f 'thanos master start';
mv /usr/local/thanos/master.db /master.db;

rm -rf /usr/local/thanos;
mkdir -p /usr/local/thanos;
cd /usr/local/thanos;
curl "http://$1/custome_package/thanos/thanos_linux_amd64.tar.gz" -o minion.tar.gz -s;
tar -zxvf minion.tar.gz;
touch nohup.log;
chmod 777 nohup.log;
ls -al;

mv /master.db /usr/local/thanos/master.db;

nohup ./thanos master start >> nohup.log 2>&1 <&- &

sleep 3;
netstat -lntp | grep 8333;
/usr/local/thanos/thanos master ui;
/usr/local/thanos/thanos master;

echo "安装完成之后,请仔细检查config.toml文件, 如果有需要请vim编辑config.toml文件"
