
# Test web-app to use with Pluralsight courses and Docker Deep Dive book
# Linux x64
FROM willzegers/dotfiles
ARG USER=will.zegers
ARG VERSION="1.15.3"

USER root
WORKDIR /tmp
ADD https://golang.org/dl/go${VERSION}.linux-amd64.tar.gz .
RUN tar -C /usr/local -xzf go${VERSION}.linux-amd64.tar.gz
VOLUME /go

USER ${USER}
WORKDIR /go
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="/go"

ENTRYPOINT ["zsh"]
