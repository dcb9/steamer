FROM ondrejmo/you-get

RUN pip3 install PySocks shadowsocks 
RUN pip3 install --upgrade you-get || exit 0

ENV GOLANG_VERSION 1.7.3
ENV GOLANG_SRC_URL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz
ENV GOLANG_SRC_SHA256 79430a0027a09b0b3ad57e214c4c1acfdd7af290961dd08d322818895af1ef44

# https://golang.org/issue/14851
COPY no-pic.patch /

RUN set -ex \
	&& apk add --no-cache --virtual .build-deps \
		bash \
		gcc \
		musl-dev \
		openssl \
		go \
	\
	&& export GOROOT_BOOTSTRAP="$(go env GOROOT)" \
	\
	&& wget -q "$GOLANG_SRC_URL" -O golang.tar.gz \
	&& echo "$GOLANG_SRC_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz \
	&& cd /usr/local/go/src \
	&& patch -p2 -i /no-pic.patch \
	&& ./make.bash \
	\
	&& rm -rf /*.patch \
	&& apk del .build-deps

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


RUN apk add --no-cache git

RUN go get github.com/gorilla/websocket

COPY . /go/src/github.com/dcb9/steamer/

RUN cd /go/src/github.com/dcb9/steamer \
  ; go get && go install

COPY run.sh /

EXPOSE 8080
ENTRYPOINT []
WORKDIR /go/src/github.com/dcb9/steamer
