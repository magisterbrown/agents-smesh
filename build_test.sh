make 
./server &> /tmp/gotest.log &
sid=$!

# TESTS:

curl  "http://localhost:8090/leaderboard"
curl -X POST "http://localhost:8090/leaderboard"

echo "SERVER logs:"
cat /tmp/gotest.log
kill -9 $sid
