FROM ubuntu:14.04

ADD server /usr/bin/

ENTRYPOINT /usr/bin/server

