FROM golang:1.23.2-bookworm

WORKDIR /usr/src/app

RUN git config --global --add safe.directory /usr/src/app

ENTRYPOINT [ "/bin/bash" ]
