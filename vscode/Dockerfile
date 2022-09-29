FROM gitpod/openvscode-server:latest

ARG GOVERSION=1.19.1
ARG GOVSCODEEXTVERSION=0.35.2

# to get permissions to install packages and such
USER root

RUN apt-get update && apt-get -y install zip

USER openvscode-server

# the installation process for software needed
ENV GOVERSION=${GOVERSION}
ENV GOVSCODEEXTVERSION=${GOVSCODEEXTVERSION}
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
ENV EXTDIR=$HOME/.openvscode-server/extensions/golang.go
RUN wget -O go.vsix https://github.com/golang/vscode-go/releases/download/v${GOVSCODEEXTVERSION}/go-${GOVSCODEEXTVERSION}.vsix
RUN mkdir -p ${EXTDIR}
RUN unzip go.vsix "extension/*" -d ./ && mv extension/* ${EXTDIR}/
