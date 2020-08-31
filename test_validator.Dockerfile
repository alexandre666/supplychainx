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
EXPOSE 26657

# Import custom config
ADD test_env/start_node.sh .
ADD test_env/config.toml /root/.scxd/config/
ADD test_env/genesis.json /root/.scxd/config/

# Import custom validator key
ADD test_env/testkey.json /root/.scxd/config/priv_validator_key.json

# To have a deterministic node-id (8545e9ece6bd63876d06ae51c9a28238b9ce729b)
ADD test_env/node_key.json /root/.scxd/config/node_key.json

RUN chmod +x ./start_node.sh

ENTRYPOINT [ "./start_node.sh" ]