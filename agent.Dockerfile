FROM golang:1.8-alpine

ENV YMIR_DIR /go/src/github.com/arlert/ymir

RUN set -ex \
    && apk add --no-cache make git

WORKDIR $YMIR_DIR
COPY . $YMIR_DIR 

# RUN make build_agent

ENTRYPOINT ["./entrypoint.sh"]

