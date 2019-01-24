FROM golang:alpine as build-env
RUN mkdir /go/src/github.com/gen1us2k/service-discovery/api -p
COPY api /go/src/github.com/gen1us2k/service-discovery/api
RUN cd /go/src/github.com/gen1us2k/service-discovery/api && go build -o api

# final stage
FROM alpine:latest
RUN echo "@edge http://nl.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories  && apk --no-cache add ca-certificates dumb-init@edge openssl curl
COPY --from=build-env /go/src/github.com/gen1us2k/service-discovery/api/api /api
RUN ["chmod", "+x", "/api"]
RUN ls -lah /
CMD ["/api"]
