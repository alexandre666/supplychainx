#!/bin/sh

scxd unsafe-reset-all
scxd init $MONIKER --chain-id testnet
scxd start