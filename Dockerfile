FROM golang:1.21.0 AS base

WORKDIR /go/src/github.com/BryanL95/game
EXPOSE 9001

ENV LANG en_US.UTF-8
ENV LC_ALL=C
ENV LANGUAGE en_US.UTF-8

FROM base AS build

COPY . .

RUN cd infrastructure/server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/server .

ENV USER=game

RUN addgroup --gid 1001 --system $USER && adduser -u 1001 --system $USER --gid 1001
RUN chown $USER:$USER /go/bin/server -R

FROM scratch AS final

COPY --from=build  /go/bin/server /go/bin/server
COPY --from=build  /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER $USER:$USER

ENTRYPOINT ["/go/bin/server"]
