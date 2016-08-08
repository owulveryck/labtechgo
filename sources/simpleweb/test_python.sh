python -m SimpleHTTPServer > /dev/null 2>/dev/null&
httperf --client=0/1 --server=localhost --port=8000 --uri=/file --num-conns=5000
