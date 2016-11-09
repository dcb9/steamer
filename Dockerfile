FROM dcb9/steamer-base

RUN go get github.com/gorilla/websocket github.com/jteeuwen/go-bindata/...

COPY . /go/src/github.com/dcb9/steamer/

RUN cd /go/src/github.com/dcb9/steamer \
  ; go-bindata -o httpHandler/go-bindata.go -pkg httpHandler views \
  ; go get && go install

COPY run.sh /

EXPOSE 8080
ENTRYPOINT []
CMD ["/run.sh"]
