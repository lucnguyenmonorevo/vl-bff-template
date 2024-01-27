#!/bin/sh

CURRENT_DIR=$(
  cd $(dirname $0)
  pwd
)

cd "$CURRENT_DIR"

./copy_ssh.sh

./create_docker_network.sh

./create_docker-compose-clone-only.sh
docker compose up --build all

# install husky
#go install github.com/automation-co/husky@latest

# install golangci-lint
# mac OS only
#brew install golangci-lint
#brew upgrade golangci-lint
