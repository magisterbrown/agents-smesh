FROM docker:latest

RUN apk add --no-cache go make sqlite
COPY . /sources
WORKDIR /sources
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" make

ENTRYPOINT make run
