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
&& rm -rf /var/lib/apt/lists/*

# Install Go
ENV GOVERSION="1.10.2" \
    GOPATH="/usr/src/go" \
    PATH="/usr/src/go/bin:/usr/local/go/bin:${PATH}"
RUN cd \
&& wget https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz \
&& tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz

# Compile OpenCV
ENV CVVERSION="3.4.1"
RUN cd \
&& wget -O opencv-${CVVERSION}.zip https://github.com/opencv/opencv/archive/${CVVERSION}.zip \
&& unzip opencv-${CVVERSION}.zip \
&& mkdir opencv-${CVVERSION}/build \
&& cd opencv-${CVVERSION}/build \
&& cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_PREFIX=/usr/local .. \
&& make -j4 \
&& make install

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

# Generate JWT key pair
RUN mkdir /usr/src/keys \
&& ssh-keygen -t rsa -b 4096 -f /usr/src/keys/jwtRS256.key \
&& openssl rsa -in /usr/src/keys/jwtRS256.key -pubout -outform PEM -out /usr/src/keys/jwtRS256.key.pub

# Workdir
WORKDIR /usr/src/

# Install app
COPY . go/src/github.com/jurevic/facegrinder/
RUN cd go/src/github.com/jurevic/facegrinder/ \
&& ./cv_env.sh \
&& go get ./... \
&& go install

# Run app
CMD facegrinder serve

EXPOSE 8000
