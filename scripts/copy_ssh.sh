#!/bin/sh

# Copy the `~/.ssh` directory to the root directory.

BASENAME=$(basename $0)

CURRENT_DIR=$(
  cd $(dirname $0)
  pwd
)

if [ -d "$CURRENT_DIR/../.ssh" ]; then
  echo "${BASENAME}: WARN: .ssh directory already exists."
fi

cp -r ~/.ssh "$CURRENT_DIR/../"
chmod -R 744 "$CURRENT_DIR/../.ssh/"
echo "${BASENAME}: Copied .ssh directory."
