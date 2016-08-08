go build -o /tmp/main main.go
/tmp/main &
sleep 2
httperf --client=0/1 --server=localhost --port=8000 --uri=/file --num-conns=5000
