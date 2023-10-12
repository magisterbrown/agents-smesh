docker build . -t rman
contid=$(docker run --rm -p 8090:8090 -v /var/run/docker.sock:/var/run/docker.sock -d  rman)
docker ps
curl "http://localhost:8090/leaderboard"
curl -X POST  -H "Content-Type: multipart/form-data"  -F "submission=@howto_submit/submission.tar" "http://localhost:8090/leaderboard"
docker logs $contid 
docker kill $contid
