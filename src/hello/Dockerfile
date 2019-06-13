FROM fedora:latest


LABEL ios.k8s.display-name="go-hello-world" \
    maintainer="Keith Tenzer <ktenzer@redhat.com>"

RUN dnf install -y golang

RUN mkdir -p /go/src/go-hello-world
RUN mkdir /app

WORKDIR /go/src/go-hello-world

ENV GOPATH=/go
ENV GOBIN=/app

RUN curl https://raw.githubusercontent.com/golang/dep/v0.5.1/install.sh | sh

COPY . /go/src/go-hello-world

RUN dep ensure

RUN go install go-hello-world/src/hello

RUN chown -R 1001:0 /app && \
    chmod -R 775 /app

RUN chmod -R 777 /tmp

RUN echo "1.0" > /etc/imageversion

USER 1001

WORKDIR /app

CMD /app/hello
