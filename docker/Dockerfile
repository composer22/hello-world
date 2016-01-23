FROM google/golang:1.4

MAINTAINER Bryan-Kirk Reinhardt

RUN CGO_ENABLED=0 go get -a --installsuffix cgo -ldflags '-s' github.com/composer22/hello-world
COPY Dockerfile.final /gopath/bin/Dockerfile

CMD docker build -t composer22/hello-world /gopath/bin
