FROM golang:latest

LABEL maintainer="Carlos Navas <dante.maniacs@gmail.com>"

# Set the current directory inside the container
WORKDIR /app

COPY . .

#
# If you want create a main.go file and testing in dev environment
# you can build the app and execute it
#
# RUN go build -o .
# CMD ["./app"]

CMD go test -v ./flatten