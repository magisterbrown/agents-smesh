rm /tmp/rankings.db
make 

# TESTS:

unshare --map-root-user --fork --pid --mount-proc --user /bin/bash <<EOF
./server &> /tmp/gotest.log &
curl "http://localhost:8090/leaderboard"; echo
curl -X POST  -H "Content-Type: application/zip"  -F "file=@howto_submit/submission.zip" "http://localhost:8090/leaderboard"; echo
EOF

printf "\nSERVER logs:\n"
cat /tmp/gotest.log
