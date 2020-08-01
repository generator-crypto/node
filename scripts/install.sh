#!/usr/bin/env bash
# Устанавливаем все необходимое
GO111MODULE=on go get
GO111MODULE=on go install ./cmd/generatord
GO111MODULE=on go install ./cmd/generatorcli

if [ ! -f "~/.generatord/config/node_key.json" ]
then
  mkdir -p ~/.generatord/config
  mkdir -p ~/.generatord/data

  cp -r ./installation/genesis.json ~/.generatord/config/
  cp -r ./installation/config.toml ~/.generatord/config/

  generatorcli config chain-id generator
  generatorcli config output json
  generatorcli config indent true
  generatorcli config trust-node true
  generatorcli config node tcp://127.0.0.1:26657
fi

if [ ! -f "~/.generatord/config/genesis.json" ]
then
  cp -r ./installation/genesis.json ~/.generatord/config/
fi

if [ ! -f "~/.generatord/config/config.toml" ]
then
  cp -r ./installation/config.toml ~/.generatord/config/
fi
