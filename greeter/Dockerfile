FROM golang

RUN go get -u -v github.com/kardianos/govendor

RUN ls -l

ENV GOPATH /usr/local/app/src/

RUN ls -l $GOPATH

COPY ./vendor/vendor.json $GOPATH/vendor/vendor.json

RUN ls -l $GOPATH

RUN cd $GOPATH/src && govendor sync

