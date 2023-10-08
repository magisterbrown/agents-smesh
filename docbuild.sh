docker build . -t rman
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock rman
