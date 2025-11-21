#!/bin/bash
set -e

# Run migrations
./install.sh

GO_PATH=$(go env GOPATH)
GO_BIN=$GO_PATH/bin/air

if [ ! -f $GO_BIN ]; then
    echo "❌ Air not found. Please install it first:"
    echo "   go install github.com/air-verse/air@latest"
    exit 1
fi

echo "🚀 Starting Air..."
# $GO_BIN
# air will read .air.toml server.addr = 0.0.0.0:8080
air -c .air.toml