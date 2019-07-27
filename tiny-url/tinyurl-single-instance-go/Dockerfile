FROM mysql:8.0

RUN apt update
RUN apt-get install -y wget
RUN apt-get install -y net-tools
RUN apt-get install -y git
RUN apt-get install -y procps
RUN apt-get install -y curl


ENV MYSQL_ROOT_PASSWORD=demo

RUN wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
RUN tar -xzvf go1.12.7.linux-amd64.tar.gz
RUN mv go /usr/local

ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/noteanddata/system-design-demo/tiny-url
WORKDIR $GOPATH/src/github.com/noteanddata/system-design-demo/tiny-url

COPY . .

# to init a database 
COPY db/db-init.sql /docker-entrypoint-initdb.d

RUN go get -d ./...

COPY scripts/startup.sh /usr/local/bin/startup.sh
RUN chmod +x /usr/local/bin/startup.sh
#ENTRYPOINT ["/usr/local/bin/startup.sh"]

CMD /usr/local/bin/startup.sh

EXPOSE 8080



