# inherit from the Go official Image
ARG  CODE_VERSION=latest
FROM golang:${CODE_VERSION}

# set a workdir inside docker
WORKDIR /go/src/server

# copy . (all in the current directory) to . (WORKDIR)
COPY . .

# run a command - this will run when building the image
RUN go build -o server

# the port we wish to expose
EXPOSE 8000

# run a command when running the container
CMD ./server
