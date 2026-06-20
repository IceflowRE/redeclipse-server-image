#!/bin/bash
set -ex

# download go modules
cd go-image-updater
go mod download

# install golangci-lint and build custom golangci-lint
echo "Installing golangci-lint"
curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin
echo "Building custom golangci-lint"
golangci-lint custom --destination "/home/vscode/.local/bin" --name golangci-lint
rm $(go env GOPATH)/bin/golangci-lint
cd ..

# install OS packages
sudo apt update
sudo apt install -y graphviz
