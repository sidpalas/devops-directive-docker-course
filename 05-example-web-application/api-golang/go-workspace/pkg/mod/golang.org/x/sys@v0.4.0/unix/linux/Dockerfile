FROM ubuntu:20.04

# Disable interactive prompts on package installation
ENV DEBIAN_FRONTEND noninteractive

# Dependencies to get the git sources and go binaries
RUN apt-get update && apt-get install -y  --no-install-recommends \
        ca-certificates \
        curl \
        git \
        rsync \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Get the git sources. If not cached, this takes O(5 minutes).
WORKDIR /git
RUN git config --global advice.detachedHead false
# Linux Kernel: Released 12 Jun 2022
RUN git clone --branch v5.19-rc2 --depth 1 https://kernel.googlesource.com/pub/scm/linux/kernel/git/torvalds/linux
# GNU C library: Released 03 Feb 2022
RUN git clone --branch release/2.35/master --depth 1 https://sourceware.org/git/glibc.git

# Only for loong64, add kernel and glibc patch
RUN git clone https://github.com/loongson/golang-infra.git /git/loong64-patches \
    && git config --global user.name "golang" && git config --global user.email "golang@localhost" \
    && cd /git/loong64-patches && git checkout glibc-v2.35 && cd /git/glibc && git am /git/loong64-patches/*.patch \
    && curl -fsSL https://git.savannah.gnu.org/cgit/config.git/plain/config.sub -o /git/glibc/scripts/config.sub

# Get Go
ENV GOLANG_VERSION 1.20rc1
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 4757fb32d7514145e43d4f37713f98d8cc0ecbbb5b1737accfc84be50e1e2e32

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

ENV PATH /usr/local/go/bin:$PATH

# Linux and Glibc build dependencies and emulator
RUN apt-get update && apt-get install -y  --no-install-recommends \
        bison gawk make python3 \
        gcc gcc-multilib \
        gettext texinfo \
        qemu-user \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
# Cross compilers (install recommended packages to get cross libc-dev)
RUN apt-get update && apt-get install -y \
        gcc-aarch64-linux-gnu       gcc-arm-linux-gnueabi     \
        gcc-mips-linux-gnu          gcc-mips64-linux-gnuabi64 \
        gcc-mips64el-linux-gnuabi64 gcc-mipsel-linux-gnu      \
        gcc-powerpc-linux-gnu       gcc-powerpc64-linux-gnu   \
        gcc-powerpc64le-linux-gnu   gcc-riscv64-linux-gnu     \
        gcc-s390x-linux-gnu         gcc-sparc64-linux-gnu     \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Only for loong64, getting tools of qemu-user and gcc-cross-compiler
RUN apt-get update && apt-get install wget xz-utils -y && mkdir /loong64 && cd /loong64 \
    && wget -q https://github.com/loongson/build-tools/releases/download/2021.12.21/qemu-loongarch-2022-4-01.tar.gz \
    && tar xf qemu-loongarch-2022-4-01.tar.gz && cp ./4-1/new-world/qemu-loongarch64 /usr/bin/ \
    && rm -rf qemu-loongarch-2022-4-01.tar.gz 4-1 \
    && wget -q https://github.com/loongson/build-tools/releases/download/2021.12.21/loongarch64-clfs-2022-03-03-cross-tools-gcc-glibc.tar.xz \
    && tar xf loongarch64-clfs-2022-03-03-cross-tools-gcc-glibc.tar.xz && mv cross-tools.gcc_glibc /usr/local/cross-tools-loong64 \
    && rm -rf loongarch64-clfs-2022-03-03-cross-tools-gcc-glibc.tar.xz \
    && ln -s /usr/local/cross-tools-loong64/bin/loongarch64-unknown-linux-gnu-gcc /usr/bin/loongarch64-linux-gnu-gcc \
    && rm -rf /loong64

# Let the scripts know they are in the docker environment
ENV GOLANG_SYS_BUILD docker
WORKDIR /build/unix
ENTRYPOINT ["go", "run", "linux/mkall.go", "/git/linux", "/git/glibc"]
