docker rmi -f $(docker images -f=reference="*:player" -q)
docker image prune -f
