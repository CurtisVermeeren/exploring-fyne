FROM golang:1.18.1
WORKDIR /app
EXPOSE 8080

# Install the Go tools for the Go extension
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \ 
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \ 
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \ 
    && go install golang.org/x/tools/gopls@latest \
    && GO111MODULE=on go install golang.org/x/tools/gopls@latest

# Install gcc and graphics library header files needed for Fyne
RUN apt-get update
RUN apt-get install gcc libgl1-mesa-dev xorg-dev mesa-utils and libgl1-mesa-glx -y
RUN export LIBGL_ALWAYS_INDIRECT=0 

# Uses "robbyrussell" theme (original Oh My Zsh theme), with no plugins
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.2/zsh-in-docker.sh)" -- \
    -t robbyrussell