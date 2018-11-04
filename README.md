# FaceGrinder - facial recognition for drones

![FaceGrinder](https://raw.githubusercontent.com/Jurevic/facegrinder/master/img/screens.png)

Written in Go and Vue.js.

This application is in experimental stage thus things might not always run smoothly, the setup can be tricky.

This application allows You to easily connect Your drones live video stream to Facial recognition toolchain. Setup known
Faces database and run facial recognition on live video streams from drones. The application can be set up and
controlled over JSON REST API or via a web interface.

The application is intended to be modifiable and enable building customized video processing, recognition chains.

## Disclaimer

This app is not intended to be used for spying or any other illegal criminal activity. The app must not be used in any
way that may restrict persons right to privacy.

## Open source project

### Drones

The easiest way to use this application is with DJI drones. DJI GO app allows you to connect drones video stream to RTMP
server. Any drone that can stream to RTMP can be connected. To do that open General Setting menu and choose "Live
Streaming Platform" -> "Custom RTMP" and enter your RTMP server URL (your host IP, if your machine does not have DNS
name).

Using drone with zooming feature gives more flexibility as the distance from the face can be greater. For regular DJI
Mavic drone with digital zoom, max distance is about 6m, this is not very far. Digital zoom does not reduce live video
quality as live stream video (for Mavic) is at most 1080p which is way lower than recorded to SD.

### Prerequisites

1. DLIB package 19.10-2 or later. Notice that it is installable over apt-get only in Debian SID. So if You are not using
it You'll have to install it yourself;
2. Recognition models can be found [here](https://github.com/davisking/dlib-models). Place them in the models directory;
3. PostgreSQL database, info on how to set up and run one can be found [here](https://www.postgresql.org/docs/).

### Getting started

To start the application:
1. Install DLIB and its facial recognition models;
2. Install, run and setup PostgreSQL database;
3. Setup environment variables as shown in env.example;
4. Run go build;
5. Run facegrinder migrate;
6. Run facegrinder serve;
7. Launch web interface by visiting localhost:8000.

To run recognition:
1. Create new channel and set its password;
2. Connect Your drone to RTMP server at: :hostIp/:channel/?key=:password;
3. Go to Processors -> DefaultRTMP and hit Run.

### License

MIT