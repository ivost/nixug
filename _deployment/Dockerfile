FROM golang:alpine

LABEL maintainer="Ivo Stoyanov <ivostoy@gmail.com>"

EXPOSE 8080

ENV GO111MODULE=on

RUN apk --no-cache add ca-certificates git && \
    cd /root && \
    git clone --depth 1 https://github.com/ivost/nixug.git && \
    cd /root/nixug && \
    go build cmd/service/nixug.go

COPY config.json.template /root/nixug/config.json
COPY run.sh /root/

# todo: 2-stage build to reduce size

ENTRYPOINT ["/root/run.sh"]


