go build -o cahce_server
./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=true &

#sleep 2
#echo ">>> start test"
#curl "http://localhost:9999/api?key=Tom" &
#sleep 1
#curl "http://localhost:9999/api?key=Tom" &
#sleep 1
#curl "http://localhost:9999/api?key=Tom" &
wait