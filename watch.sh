PID=`lsof -i :8080 | awk '$1~/server/{print $2}'`
if [ "$PID" != "" ];then
  kill -9 $PID > /dev/null
fi
go run server.go $@

