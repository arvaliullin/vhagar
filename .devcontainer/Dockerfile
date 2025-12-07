ARG BASE_IMAGE=registry.astralinux.ru/library/astra/ubi18-golang121:1.8.4
FROM ${BASE_IMAGE}

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

ARG DOCKER_VERSION=26.1.0
ARG DOCKER_COMPOSE_VERSION=v2.29.2

RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    curl \
    wget \
    make \
    zsh \
    golang \
    ca-certificates \
    sudo \
    openssh-client \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-gopkg-yaml.v3-dev \
    golang-github-streadway-amqp-dev \
    golang-github-kelseyhightower-envconfig-dev \
    golang-github-lib-pq-dev \
    && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz -o /tmp/docker.tgz \
    && tar -xz -C /tmp -f /tmp/docker.tgz \
    && mv /tmp/docker/docker /usr/local/bin/docker \
    && rm -rf /tmp/docker /tmp/docker.tgz \
    && chmod +x /usr/local/bin/docker

RUN mkdir -p /usr/local/lib/docker/cli-plugins \
    && curl -fsSL https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose \
    && chmod +x /usr/local/lib/docker/cli-plugins/docker-compose

RUN if getent group docker > /dev/null 2>&1; then \
        groupmod -g 1001 docker || true; \
    else \
        groupadd -g 1001 docker; \
    fi

RUN if ! getent group $USERNAME > /dev/null 2>&1; then \
        if getent group $USER_GID > /dev/null 2>&1; then \
            groupmod -n $USERNAME $(getent group $USER_GID | cut -d: -f1); \
        else \
            groupadd --gid $USER_GID $USERNAME; \
        fi; \
    fi \
    && if ! id -u $USERNAME > /dev/null 2>&1; then \
        useradd --uid $USER_UID --gid $USER_GID -m $USERNAME; \
    fi \
    && mkdir -p /etc/sudoers.d \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && (getent group docker > /dev/null 2>&1 && usermod -aG docker $USERNAME || true)

RUN git clone https://github.com/ohmyzsh/ohmyzsh.git /home/$USERNAME/.oh-my-zsh \
    && cp /home/$USERNAME/.oh-my-zsh/templates/zshrc.zsh-template /home/$USERNAME/.zshrc \
    && chown -R $USERNAME:$USERNAME /home/$USERNAME/.oh-my-zsh /home/$USERNAME/.zshrc \
    && chsh -s $(which zsh) $USERNAME \
    && echo "0123456789abcdef0123456789abcdef" > /etc/machine-id \
    && chmod 644 /etc/machine-id

ENV GOPATH=/tmp/go-tools
ENV PATH=$PATH:/tmp/go-tools/bin

RUN mkdir -p $GOPATH/bin \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@v1.21.2 \
    && mv $GOPATH/bin/* /usr/local/bin/ \
    && rm -rf $GOPATH \
    && chmod +x /usr/local/bin/gopls /usr/local/bin/dlv

ENV GOPATH=/usr/share/gocode
ENV PATH=$PATH:/usr/local/bin

USER $USERNAME

RUN rm -rf /home/$USERNAME/go/pkg 2>/dev/null || true

WORKDIR /home/$USERNAME
