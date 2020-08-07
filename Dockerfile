FROM golangci/golangci-lint:latest

RUN \
  apt-get update && \
  apt-get install -y vim

ENV GOLINTUI_OPEN_COMMAND=vim

COPY golintui /usr/bin/
CMD ["golintui"]
