FROM golang:latest
ENV GOPATH /go

# Copy files
ADD Makefile /go/src/github.com/ltacker/poapoc/Makefile
ADD go.mod /go/src/github.com/ltacker/poapoc/go.mod
ADD go.sum /go/src/github.com/ltacker/poapoc/go.sum
ADD x /go/src/github.com/ltacker/poapoc/x
ADD cmd /go/src/github.com/ltacker/poapoc/cmd
ADD app /go/src/github.com/ltacker/poapoc/app

WORKDIR /go/src/github.com/ltacker/poapoc

# Installation
RUN make install

# Listening port
EXPOSE 26656
EXPOSE 26657

# Initialize
RUN scxd init moniker --chain-id testnet

# Import custom config
ADD test_env/config.toml /root/.scxd/config/
ADD test_env/genesis.json /root/.scxd/config/

# Import custom validator key
ADD test_env/testkey.json /root/.scxd/config/priv_validator_key.json

# To have a deterministic node-id (8545e9ece6bd63876d06ae51c9a28238b9ce729b)
ADD test_env/node_key.json /root/.scxd/config/node_key.json

ENTRYPOINT [ "scxd", "start" ]