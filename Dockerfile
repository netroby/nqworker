FROM golang

go get -u -v github.com/netroby/nqworker

CMD ["nqwoker"]
