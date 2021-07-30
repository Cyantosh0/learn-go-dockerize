FROM golang:alpine

# if not included then error occurs in dependency import using go module path
ENV GO111MODULE=on

# define environment variables
ENV SERVER_PORT=5000
ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_NAME=learn_docker_db
ENV DB_USER=root
ENV DB_PASSWORD=password

# create your own working directory
WORKDIR /my-app

COPY go.mod go.sum ./

RUN go mod download

# Copy all files from the directory to the docker image
COPY . .

RUN go build

# If not included:doesnot run the app And must have the name of directory/app where this Dockerfile is in
CMD ["./dockerize-go-app"]