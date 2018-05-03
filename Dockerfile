FROM postgres:10.1-alpine

EXPOSE 5432

FROM golang:1.9.2-alpine3.7

# Install and run
COPY . /go/src/facegrinder/
RUN cd /go/src/facegrinder/cmd/runserver/ \
&& go build

EXPOSE 8080

CMD cd /go/src/facegrinder/cmd/runserver/ \
&& ./runserver
