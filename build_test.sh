rm /tmp/rankings.db
make 

# TESTS:

unshare --fork --pid --mount-proc --user /bin/bash <<EOF
./server &> /tmp/gotest.log &
curl  "http://localhost:8090/leaderboard"; echo
curl -X POST "http://localhost:8090/leaderboard"; echo
EOF

printf "\nSERVER logs:\n"
cat /tmp/gotest.log
