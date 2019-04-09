# Build uche-data-sync in a stock Go builder container
FROM golang:latest as builder

ADD . $GOPATH/src/github.com/yuhu-tech/uche-data-sync
RUN mkdir /uche-data-sync && mkdir /uche-data-sync/bin && mkdir /uche-data-sync/conf && mkdir /var/log/uche-data-sync
RUN cd $GOPATH/src/github.com/yuhu-tech/uche-data-sync/cmd/collection && go build && mv collection /uche-data-sync/bin/
RUN cd $GOPATH/src/github.com/yuhu-tech/uche-data-sync/configs && cp collection.toml /uche-data-sync/conf/

CMD ["/uche-data-sync/bin/collection", "-conf", "/uche-data-sync/conf/collection.toml"]
