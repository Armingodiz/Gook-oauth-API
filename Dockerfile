# this is image making approach , in this approach we need to make sure that we have redis container up and running too in the specified ipaddres .
# using docker network or docker compose is a better choice for multi container applications like this .
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
# in this way size of our image will be inefficent we can replace line 14  :
# putting binary file in first layer
# RUN CGO_ENABLED=0 GDOS=linux go build -a -installsuffix cgo -o outh-api .
# FROM alpine:latest
# RUN apk --no-cashe add ca-certificates
# WORKDIR /root/
# using first layer (0) to create image with alpine base layer
# COPY --from=0 /src/app .
# CMD ["./outh-api"]
# this approach causes alpine to be base image insted of golang which hase much smaller size