# go-hello-world
A hello world microservice written in Golang.

## Usage
A Dockerfile is provided so you can just do a docker build. If you are running OpenShift simple run the template provided under the yaml folder.

```$ oc process -f yaml/deploy-go-hello-world.yaml```

If you are interested in building the code and setting up a development environment please follow the instructions in the blog below.
[Getting Started with Golang](https://keithtenzer.com/2019/06/13/the-coffee-is-getting-cold-its-time-to-go-getting-started-building-microservice/)
