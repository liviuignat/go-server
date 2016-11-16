FROM golang:1.7.3
# Add glide to the image

RUN wget https://github.com/Masterminds/glide/releases/download/0.10.1/glide-0.10.1-linux-amd64.tar.gz && \
    tar xvfz glide-0.10.1-linux-amd64.tar.gz -C /usr/local/bin --strip-components=1 linux-amd64/glide && \
    rm glide-0.10.1-linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

# go-wrapper is the unchanged one from golang Docker, may need to enhance this to incorporate running glide
# COPY scripts/go-wrapper /usr/local/bin/

# COPY scripts/go-build-help /usr/local/bin/

#---- below came from https://github.com/docker-library/golang/blob/ce284e14cdee73fbaa8fb680011a812f272eae2e/1.6/onbuild/Dockerfile

RUN mkdir -p /go/src/go-server
WORKDIR /go/src/go-server

COPY . /go/src/go-server

RUN glide install
RUN go build -o main .

ENV PORT=3000
ENV SERVICE_NAME=go-server

EXPOSE 3000

CMD ["/go/src/go-server/main"]
