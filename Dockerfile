# starting from base image 1.13.8
FROM golang:1.13

# configuring env variables
ENV REPO_URL=github.com/ArminGodiz/Gook-oauth-API
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL

# copy the entire project to the workpath of container(image)
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

# compiling project to binary named outh-api and placing it in the current directory
RUN go build -o outh-api .

# expose the port we are listenning on
EXPOSE 2222

# running the binary file

CMD ["./outh-api"]