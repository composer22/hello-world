### [Dockerized] (http://www.docker.com) [hello-world](https://registry.hub.docker.com/u/composer22/hello-world/)

A docker image for hello-world. This is created as a single "static" executable using a lightweight image.

To make:

cd docker
./build.sh

Once it completes, you can run the server. For example:

docker run -d -p 8080:8080 --name hello-world composer22/hello-world
