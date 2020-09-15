FROM ubuntu:18.04
LABEL maintainer="Yogesh Kulkarni (yrkulkarni@live.com)"
LABEL updated="15 Sep 2020"
LABEL updated_by="codeyrk"
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8

USER root

RUN apt-get update && apt-get install -y \
  sudo \
  ccache \
  wget \
  build-essential \
  software-properties-common \
  git && \
  rm -rf /var/lib/apt/lists/*

RUN groupadd -r nonroot && \
  useradd -m -d /home/nonroot -g nonroot -s /usr/sbin/nologin -c "Nonroot User" nonroot && \
  mkdir -p /home/nonroot/workdir && \
  chown -R nonroot:nonroot /home/nonroot && \
  usermod -a -G sudo nonroot && echo 'nonroot:nonroot' | chpasswd

RUN git clone https://github.com/radare/radare2.git && \
  cd radare2 && \
  ./sys/install.sh && \
  r2pm init && \
  r2pm update

RUN add-apt-repository ppa:longsleep/golang-backports

RUN apt-get install -y golang-go

RUN r2pm -i r2pipe-go

ARG currDir=/var/src
RUN mkdir ${currDir}

# Set the Current Working Directory inside the container
WORKDIR ${currDir}

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN go mod verify

## Build the Go app
RUN go build -o ./r2-go

USER nonroot

WORKDIR ${currDir}

CMD ./r2-go -i /binaries/input -o /binaries/output -p /binaries/processed

