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

# Initialize
RUN scxd init moniker --chain-id testnet

# Import custom config
ADD test_env/config.toml /root/.scxd/config/
ADD test_env/genesis.json /root/.scxd/config/

ENTRYPOINT [ "scxd", "start" ]