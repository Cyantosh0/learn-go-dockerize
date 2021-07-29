# Learn Docker with me

## Some most used docker commands

#### docker images
To list all the docker images

#### docker build -t image-name .
To create image from Dockerfile by providing name to image
- If you get permission error then add [sudo] in the front

#### docker rmi image-name
To delete image
If you have to force delete the image
> docker rmi -f image-name

#### docker rmi -f $(docker images --filter "dangling=true" -q --no-trunc)
When you want to build over an existing image with existing image name then previous image name will become <none>.
Use this command to remove all the images with <none>

#### docker run image-name
This command will run the container from the image. But the container will run in isolation; hence connection to the server will gets refused whenever API is called. This requires networking to be able to connect with server.

#### docker run -p 8080:8080 image-name
To run a the container at local network using port 8080 from the image. Now we can connect to the server.

#### docker ps
To list all the docker containers that are being run.

#### docker stop container_id
To stop docker container that is being run