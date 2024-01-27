FROM alpine:3

# hadolint ignore=DL3018
RUN apk update && \
    apk --no-cache add \
    sudo \
    git \
    grep \
    curl \
    openssh-client \
    docker \
    docker-cli-buildx \
    docker-cli-compose

# Create and change user.
ENV USER setup
ENV HOME /home/$USER

RUN echo "$USER ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers && \
    addgroup -S $USER && \
    adduser -S -u 1000 -G $USER $USER && \
    chown -R $USER:$USER $HOME
USER $USER

WORKDIR $HOME

# Copy .ssh directory.
RUN mkdir -m 0700 $HOME/.ssh
COPY .ssh/* $HOME/.ssh/

# Force ssh during git clone.
RUN git config --global url.git@github.com:.insteadOf https://github.com/

ENTRYPOINT [ "sh", "-c", \
    "sh \"${BIND_MOUNT_CONTAINER_DIR}/infrastructures/setup/docker-entrypoint.sh\" \
    \"${SETUP_ONLY}\" \
    \"${BIND_MOUNT_HOST_DIR}\" \
    \"${BIND_MOUNT_CONTAINER_DIR}\" \
    \"${REPOSITORY_NAME}\" \
    \"${BRANCH_NAME}\" \
    \"${COMPOSE_FILE_NAME}\" \
    \"${SERVICE_NAME}\" \
    \"${BUILD}\" \
    \"${DEPENDENCY_SERVICE_NAMES}\"" ]
