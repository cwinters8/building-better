FROM gitpod/openvscode-server:latest

ARG GOVERSION=1.19.1
ARG GOEXTVERSION=0.35.2
ARG PRETTIEREXTVERSION=9.9.0

# to get permissions to install packages and such
USER root

# the installation process for software needed
RUN apt-get update && apt-get -y install zip

USER openvscode-server

ENV GOVERSION=${GOVERSION}
ENV GOEXTVERSION=${GOEXTVERSION}
ENV GOPATH=$HOME/go-packages
ENV GOROOT=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN curl -fsSL https://dl.google.com/go/go$GOVERSION.linux-amd64.tar.gz | tar xzs
# install VS Code Go tools for use with gopls as per https://github.com/golang/vscode-go/blob/master/docs/tools.md
# also https://github.com/golang/vscode-go/blob/27bbf42a1523cadb19fad21e0f9d7c316b625684/src/goTools.ts#L139
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN sudo rm -rf $GOPATH/src $GOPATH/pkg $HOME/.cache/go $HOME/.cache/go-build
# user Go packages
ENV GOPATH=/workspace/go \
    PATH=/workspace/go/bin:$PATH

# extensions
ENV EXTDIR=$HOME/.openvscode-server/extensions
ENV GODIR=${EXTDIR}/golang.go
RUN wget -O go.vsix https://github.com/golang/vscode-go/releases/download/v${GOEXTVERSION}/go-${GOEXTVERSION}.vsix
RUN mkdir -p ${GODIR} && unzip go.vsix "extension/*" -d ./ && mv extension/* ${GODIR}/
ENV PRETTIERDIR=${EXTDIR}/prettier
RUN wget -O prettier.vsix https://github.com/prettier/prettier-vscode/releases/download/v${PRETTIEREXTVERSION}/prettier-vscode-${PRETTIEREXTVERSION}.vsix
RUN mkdir -p ${PRETTIERDIR} && unzip prettier.vsix "extension/*" -d ./ && mv extension/* ${PRETTIERDIR}/

WORKDIR $HOME

# load lesson content from lessons repo (TODO: move lesson content into a gh repo)

ENV CONNECTION_TOKEN=default
ENV OPENVSCODE_SERVER_ROOT="/home/.openvscode-server"
ENTRYPOINT [ "/bin/sh", "-c", "exec ${OPENVSCODE_SERVER_ROOT}/bin/openvscode-server --host 0.0.0.0 --connection-token ${CONNECTION_TOKEN} \"${@}\"", "--" ]
