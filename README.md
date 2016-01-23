# hello-world
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/composer22/hello-world.svg?branch=master)](http://travis-ci.org/composer22/hello-world)
[![Current Release](https://img.shields.io/badge/release-v0.0.1-brightgreen.svg)](https://github.com/composer22/hello-world/releases/tag/v0.0.1)
[![Coverage Status](https://coveralls.io/repos/composer22/hello-world/badge.svg?branch=master)](https://coveralls.io/r/composer22/hello-world?branch=master)

A service for use in testing Docker service aware wireups written in [Go.](http://golang.org)

## About

This little app performs nothing except providing a "hello-world" endpoint and a
valid health check. It can be used to test service discovery in the larger context of
using registrators and routers along with Docker clustering solutions like swarm.

## Requirements

## CLI Usage

```
Description: hello-world is a server for testing service awareness.

Usage: hello-world

Server options:
    none

Common options:
    none

Example:

    hello-world
```
Please also see /docker dir for more information on running this service.

## HTTP API


Example cURL:

```
$ curl -i -H "Accept: application/json" \
-H "Content-Type: application/json" \
-X GET "http://0.0.0.0:8080/v1.0/health"

HTTP/1.1 200 OK
Content-Type: application/json;charset=utf-8
Date: Fri, 03 Apr 2015 17:29:17 +0000
X-Request-Id: DC8D9C2E-8161-4FC0-937F-4CA7037970D5
Content-Length: 0
```

Three API routes are provided for service measurement:

* http://localhost:8080/v1.0/health - GET: Is the server alive?
* http://localhost:8080/v1.0/hello - GET: return "hello World!"

## Building

This code currently requires version 1.42 or higher of Go.

You will need to install the following dependencies:
```
go get github.com/composer22/hello-world
```
Information on Golang installation, including pre-built binaries, is available at
<http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `hello-world` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

## Docker Images

A prebuilt docker image is available at (http://www.docker.com) [hello-world](https://registry.hub.docker.com/u/composer22/hello-world/)

If you have docker installed, run:
```
docker pull composer22/hello-world:latest

or

docker pull composer22/hello-world:<version>

if available.
```
See /docker directory README for more information on how to run it.

## License

(The MIT License)

Copyright (c) 2015 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
