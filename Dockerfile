FROM golangci/golangci-lint:latest-alpine

RUN \
  apk update && \
  apk add vim

ENV GOLINTUI_OPEN_COMMAND=vim

COPY golintui /usr/bin/
CMD ["golintui"]
