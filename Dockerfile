FROM debian:sid-slim

# OS packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    wget \
    ca-certificates \
    git \
    g++ \
    gcc \
    libc6-dev \
    cmake \
    make \
    pkg-config \
    openssh-client \
    libgtk2.0-dev \
    libavcodec-dev \
    libavformat-dev \
    libswscale-dev \
    libjpeg-dev \
    libpng-dev \
    libblas-dev \
    liblapack-dev \
    libdlib19 \
    libdlib-dev \
    bzip2 \
    unzip \
    sudo \
&& rm -rf /var/lib/apt/lists/*

# Install Go
ENV GOVERSION="1.10.2" \
    GOPATH="/usr/src/go" \
    PATH="/usr/src/go/bin:/usr/local/go/bin:${PATH}"
RUN cd \
    && wget https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz

# Compile Dlib
ENV DLIBVERSION="19.13"
RUN cd \
    && wget http://dlib.net/files/dlib-${DLIBVERSION}.tar.bz2 \
    && tar jxvf dlib-${DLIBVERSION}.tar.bz2 \
    && mkdir dlib-${DLIBVERSION}/build \
    && cd dlib-${DLIBVERSION}/build \
    && cmake .. \
    && cmake --build . --config Release \
    && make install

# Workdir
WORKDIR $GOPATH/src/github.com/jurevic/facegrinder/

# Copy source
COPY . .

# Get dependencies
RUN go get -d

# Install openCV
RUN cd $GOPATH/src/gocv.io/x/gocv \
    && make install

# Install app
RUN cd $GOPATH/src/github.com/jurevic/facegrinder/ \
    && go install

# Generate JWT key pair
RUN mkdir /usr/src/keys \
    && ssh-keygen -t rsa -b 4096 -f /usr/src/keys/jwtRS256.key \
    && openssl rsa -in /usr/src/keys/jwtRS256.key -pubout -outform PEM -out /usr/src/keys/jwtRS256.key.pub

# Run app
CMD ./docker-entrypoint.sh

EXPOSE 80
