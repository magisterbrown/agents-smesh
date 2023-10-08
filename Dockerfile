FROM docker:latest

RUN apk add --no-cache go
COPY ./mobytr.go go.mod go.sum .
COPY ./howto_submit/submission/Dockerfile ./howto_submit/subm.tar .
RUN --mount=type=cache,target=/root/go go build -gcflags '-N -l' -o ./rune mobytr.go

ENTRYPOINT ./rune
