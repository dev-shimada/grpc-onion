#!/bin/bash
set -e

apt-get update && apt-get install -y vim git sqlite3 unzip

# tools
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc
git config --local core.editor vim
git config --local pull.rebase false

# protoc compiler
cd /tmp
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-aarch_64.zip
unzip protoc-26.1-linux-aarch_64.zip -d /usr/local
rm -rf /tmp/bin /tmp/include /tmp/protoc-26.1-linux-aarch_64.zip

# buf cli
BIN="/usr/local/bin" && \
VERSION="1.35.1" && \
curl -sSL \
"https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
-o "${BIN}/buf" && \
chmod +x "${BIN}/buf"

# protoc-gen-go
cd /server
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo export PATH="$PATH:/usr/local/protobuf/bin" >> ~/.bashrc
echo export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
