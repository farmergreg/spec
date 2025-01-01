###############
#  Arguments  #
###############
# Toolchain Software Versions
ARG VERSION_GOLANG=1.23.4-bookworm

# Development Tool Versions
ARG VERSION_DELVE=latest
ARG VERSION_GOPLS=latest
ARG VERSION_GOIMPORTS=latest
ARG VERSION_STATICCHECK=latest
ARG VERSION_GOCRITIC=latest

# User / Group ID
ARG USER=adif
ARG UID=1000
ARG GID=1000

######################
#  Stage: Toolchain  #
######################
FROM golang:${VERSION_GOLANG} AS toolchain
ARG USER
ARG UID
ARG GID

# create ${USER} and configure home directory
ENV HOME=/home/${USER}
ENV XDG_CACHE_HOME=${HOME}/.cache
RUN groupadd -g ${GID} ${USER} && \
    useradd -m -u ${UID} -g ${GID} -d ${HOME} -s /bin/sh ${USER}

# Drop permissions and become ${USER}
USER ${UID}:${GID}
WORKDIR /app

# configure the go path, and create the pkg/mod directory so that we can mount a cache there
ENV GOPATH=${HOME}/go
ENV PATH=${GOPATH}/bin:${PATH}
RUN mkdir -p ${GOPATH}/pkg/mod


#################################
#  Stage: vscode dev container  #
#################################
# vscode dev container
FROM toolchain AS dev-container
ARG USER
ARG UID
ARG GID

ARG VERSION_DELVE
ARG VERSION_GOPLS
ARG VERSION_GOIMPORTS
ARG VERSION_STATICCHECK
ARG VERSION_GOCRITIC

USER root:root
RUN apt-get update && apt-get upgrade -y && apt-get autoremove -y && apt-get autoclean && \
    apt-get install -y                                                          \
    fish                                                                        \
    graphviz                                                                    \
    ripgrep                                                                     \
    sudo                                                                        \
    vim-nox                                                                   &&\
    rm -rf /var/lib/apt/lists/*

RUN \
    chsh -s /usr/bin/fish ${USER}                                           &&  \
    mkdir -p /home/${USER}/.config/fish                                     &&  \
    echo "set -U fish_greeting" > /home/${USER}/.config/fish/config.fish    &&  \
    echo "${USER} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/${USER}         &&  \
    chmod 0440 /etc/sudoers.d/${USER}                                       &&  \
    chown -R ${USER}:${GID} /home/${USER}/.config

USER ${UID}:${GID}
RUN --mount=type=cache,target=${HOME}/go/pkg/mod,uid=${UID},gid=${GID}              \
    --mount=type=cache,target=${XDG_CACHE_HOME}/go-build,uid=${UID},gid=${GID}      \
    go install github.com/go-delve/delve/cmd/dlv@${VERSION_DELVE}               &&  \
    go install golang.org/x/tools/gopls@${VERSION_GOPLS}                        &&  \
    go install golang.org/x/tools/cmd/goimports@${VERSION_GOIMPORTS}            &&  \
    go install honnef.co/go/tools/cmd/staticcheck@${VERSION_STATICCHECK}        &&  \
    go install -v github.com/go-critic/go-critic/cmd/gocritic@${VERSION_GOCRITIC}

WORKDIR /app
CMD ["sleep", "infinity"]