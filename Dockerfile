FROM alpine

ADD server /usr/bin/

ENTRYPOINT /usr/bin/server

