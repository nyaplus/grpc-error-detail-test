FROM golang:latest

ARG LANG="ja_JP.utf8"
# ARG LC_ALL='ja_JP.utf8'
ARG LOCALE_ZONE="Japan"
ARG USERNAME="vscode"
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# https://github.com/protocolbuffers/protobuf/tags
ARG PROTOBUF_VERSION=22.3
# https://github.com/bufbuild/buf/tags
ARG BUF_VERSION=1.17.0


ENV \
    # LANG="${LANG}" \
    # LC_ALL="${LC_ALL}"
    LANG="${LANG}"

RUN apt-get update \
    && cp "/usr/share/zoneinfo/${LOCALE_ZONE}" /etc/localtime \
    && echo "${LOCALE_ZONE}" >/etc/timezone \
    && groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd -s /bin/bash --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} \
    && apt-get install -y sudo \
    && echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME} \
    && apt-get install -y locales \
    && locale-gen ${LANG} \
    && apt-get install -y \
    default-mysql-client \
    unzip \
    curl \
    git \
    jq \
    vim \
    less

# protoc
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-aarch_64.zip -o protoc.zip && \
    unzip protoc.zip && \
    mv bin/* /usr/local/bin/ && \
    mv include/* /usr/local/include/

# buf
WORKDIR /tmp/buf
RUN curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "/usr/local/bin/buf" && \
    chmod +x "/usr/local/bin/buf"

USER vscode

# set env varivales for user 'vscode'
RUN echo "export LANG=${LANG}" >> ~/.bashrc \
    && echo "export LESSCHARSET=utf-8" >> ~/.bashrc \
    && echo "alias ll='ls -l'" >> ~/.bashrc

# vim setting
RUN echo '\n\
    set fenc=utf-8\n\
    set encoding=utf-8\n\
    set fileencodings=iso-2022-jp,euc-jp,sjis,utf-8\n\
    set fileformats=unix,dos,mac\n\
    syntax on' >> ~/.vimrc

# go tools
RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v github.com/cweill/gotests/gotests@v1.6.0 \
    && go install -v mvdan.cc/sh/v3/cmd/shfmt@v3.2.1 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/golang/mock/mockgen@v1.6.0

# grpc tools (required both of 'protoc-gen-go' and 'protoc-gen-go-grpc')
RUN go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    # ProtocolBuffer generator
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    # gRPCurl
    && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
