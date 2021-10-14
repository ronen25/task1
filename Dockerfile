FROM golang:1.17-alpine

# Container listens on port 8001
ENV SERVER_PORT=8001

# Copy all sources and build
COPY . /helloworld/

WORKDIR /helloworld
RUN go build ./...

EXPOSE ${SERVER_PORT}

CMD ["/helloworld/task1"]
