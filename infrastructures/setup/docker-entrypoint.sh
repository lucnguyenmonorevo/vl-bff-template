#!bin/sh

# set -x # For debug.

BASENAME=$(basename $0)

# Get arguments.
SETUP_ONLY="$1"
BIND_MOUNT_HOST_DIR="$2"
BIND_MOUNT_CONTAINER_DIR="$3"
REPOSITORY_NAME="$4"
BRANCH_NAME="$5"
COMPOSE_FILE_NAME="$6"
SERVICE_NAME="$7"
BUILD="$8"
DEPENDENCY_SERVICE_NAMES="$9"

# Evaluate ssh-agent.
echo "${BASENAME}: Evaluate ssh-agent."
eval "$(ssh-agent -s)"
# ssh -T git@github.com # For debug.

git config --global submodule.recurse true

# git clone or git checkout.
cd "$BIND_MOUNT_CONTAINER_DIR"

if [ ! -z "$REPOSITORY_NAME" ] && [ ! -z "$BRANCH_NAME" ]; then
  if [ ! -d "$REPOSITORY_NAME" ] || [ ! -d "$REPOSITORY_NAME/.git" ]; then
    echo "${BASENAME}: Clone repository from github."
    git clone --recursive -b $BRANCH_NAME git@github.com:mono-revo/$REPOSITORY_NAME.git
  else
    git config --global --add safe.directory "$BIND_MOUNT_CONTAINER_DIR/$REPOSITORY_NAME"

    echo "${BASENAME}: Get current branch name from github."
    CURRENT_BRANCH_NAME=$( (
      cd "$REPOSITORY_NAME"
      git rev-parse --abbrev-ref HEAD
    ))
    if [ "$CURRENT_BRANCH_NAME" != "$BRANCH_NAME" ]; then
      echo "${BASENAME}: Check out branch from github."
      (
        cd "./$REPOSITORY_NAME"
        git checkout $BRANCH_NAME
      )
    fi
    echo "${BASENAME}: Pull branch from github."
    (
      cd "./$REPOSITORY_NAME"
      git pull origin $BRANCH_NAME
    )
    echo "${BASENAME}: Update submodule."
    (
      cd "./$REPOSITORY_NAME"
      git submodule update --init --remote --recursive
    )
  fi
fi

if [ "$SETUP_ONLY" == "true" ]; then
  exit 0
fi

if [ ! -e $BIND_MOUNT_CONTAINER_DIR/$REPOSITORY_NAME/$COMPOSE_FILE_NAME ]; then
  echo "$BIND_MOUNT_CONTAINER_DIR/$REPOSITORY_NAME/$COMPOSE_FILE_NAME: no such file"
  exit 1
fi

# Run command.
echo "${BASENAME}: Run command."
if [ ! -z "$COMMAND" ]; then
  sh -c "$COMMAND"
fi

# Set environment variable for docker compose.
export COMPOSE_FILE="\
./docker-compose.yml:\
./$REPOSITORY_NAME/$COMPOSE_FILE_NAME"
export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1
export BIND_MOUNT_HOST_DIR="$BIND_MOUNT_HOST_DIR/$REPOSITORY_NAME"
export CONTEXT_DIR="$BIND_MOUNT_CONTAINER_DIR/$REPOSITORY_NAME"

# Terminate the container when finished.
stop_service() {
  sudo -E docker compose stop $SERVICE_NAME
  if [ ! -z "$DEPENDENCY_SERVICE_NAMES" ]; then
    for dependency_service_name in ${DEPENDENCY_SERVICE_NAMES//,/ }; do
      sudo -E docker compose stop $dependency_service_name &
    done
  fi
}

trap 'stop_service' HUP
trap 'stop_service' INT
trap 'stop_service' QUIT
trap 'stop_service' TERM
trap 'stop_service' EXIT

# Up service.
echo "${BASENAME}: Up service."

if [ -z "$BUILD" ] || [ "$BUILD" == "true" ]; then
  dcup_cmd="docker compose up --build $SERVICE_NAME"
  sudo -E docker compose up --build $SERVICE_NAME &
else
  dcup_cmd="docker compose up $SERVICE_NAME"
  sudo -E docker compose up $SERVICE_NAME &
fi

# Waiting to terminate.
set +x

RAN=0
CONTAINER_NAME=vl-${SERVICE_NAME}

while :; do

  if [ $RAN -eq 0 ]; then
    sleep 0.1
    sudo -E docker ps --format="{{.Names}}" | grep -x "${CONTAINER_NAME}"
    not_run=`echo $?`
    # echo not_run=$not_run
    if [ ${not_run} -eq 0 ]; then
      RAN=1
      # echo RAN!!!!
      CONTAINER_ID=`sudo -E docker ps --all --filter "name=/${CONTAINER_NAME}$" --format="{{.ID}}"`
      # echo CONTAINER_ID=$CONTAINER_ID
    fi
  else
    sleep 1
    x=`ps | grep -E " root .*${dcup_cmd}$"`
    shutdown=$(echo $?)
    # echo "shutdown: $shutdown"
    if [ ${shutdown} -eq 1 ]; then
      EXIT_CODE=`sudo -E docker inspect ${CONTAINER_ID} --format="{{.State.ExitCode}}"`
      exit $EXIT_CODE
    fi
  fi
done
