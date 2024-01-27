#!/bin/sh

# Create `docker-compose.yml` file from `docker-compose-template-clone-only.yml` file.

BASENAME=$(basename $0)

CURRENT_DIR=$(
  cd $(dirname $0)
  pwd
)

cp -f "$CURRENT_DIR/../docker-compose-template-clone-only.yml" "$CURRENT_DIR/../docker-compose.yml"
echo "${BASENAME}: Created docker-compose.yml file."
