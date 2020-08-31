FROM golang:latest
ENV GOPATH /go

# Copy files
ADD Makefile /go/src/github.com/ltacker/supplychainx/Makefile
ADD go.mod /go/src/github.com/ltacker/supplychainx/go.mod
ADD go.sum /go/src/github.com/ltacker/supplychainx/go.sum
ADD x /go/src/github.com/ltacker/supplychainx/x
ADD cmd /go/src/github.com/ltacker/supplychainx/cmd
ADD app /go/src/github.com/ltacker/supplychainx/app

WORKDIR /go/src/github.com/ltacker/supplychainx

# Installation
RUN make install

# Listening port
EXPOSE 26656

# Import custom config
ADD test_env/start_node.sh .
ADD test_env/config.toml /root/.scxd/config/
ADD test_env/genesis.json /root/.scxd/config/

RUN chmod +x ./start_node.sh

ENTRYPOINT [ "./start_node.sh" ]