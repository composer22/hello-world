#!/bin/bash
docker build -t composer22/hello-world_build .
docker run -v /var/run/docker.sock:/var/run/docker.sock -v $(which docker):$(which docker) -ti --name hello-world_build composer22/hello-world_build
docker rm hello-world_build
docker rmi composer22/hello-world_build
