# sample to build the image to access database from host (just for test)
# docker build . -t tinyurl --build-arg dbhost=host.docker.internal
# docker run -it -p 8080:8080 tinyurl


FROM golang:1.12
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git
    
RUN mkdir -p $GOPATH/src/github.com/noteanddata/system-design-demo/tiny-url
WORKDIR $GOPATH/src/github.com/noteanddata/system-design-demo/tiny-url

COPY . .

RUN go get -d ./...


ARG dbhost
ENV dbhost ${dbhost}
CMD ["bash", "-c", "go run tinyurl.go ${dbhost}"]

EXPOSE 8080



