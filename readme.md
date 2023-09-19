#### Auditing directions: https://github.com/01-edu/public/tree/master/subjects/ascii-art-web/dockerize/audit #### 

---  

# Dockerize

**Description**  
Ascii-art-web is a webpage where you can turn normal text into ascii art. In this project Docker is also utilized.

## Running the Dockerized Go Application ##
This guide will walk you through running a Go application in a Docker container. The provided Dockerfile uses the official Go image as a base and demonstrates a typical workflow for containerizing a Go application.

> ### Prerequisites ###
> Before you begin, make sure you have the following installed on your system:  
> - [Docker](https://www.docker.com/get-started/)

### Building the Docker Image ###
1. Clone or download the repository containing the Go application and the Dockerfile.
2. Open a terminal and navigate to the directory containing the Dockerfile and the Go application code.
3. Run the following command to build a Docker image tagged as docker-asciiweb. This command instructs Docker to build an image using the Dockerfile and tag it with the name docker-asciiweb. 
- **docker build --tag docker-asciiweb .**

### Running the Docker Container ###
Once the Docker image is built, you can run a Docker container from it. This command starts a Docker container from the docker-asciiweb image and maps port 8080 from the container to the same port on your host machine. Your Go application will be accessible at http://localhost:8080 in your web browser or via a web API request.
- **docker run -p 8080:8080 docker-asciiweb**

### Stopping the Container ###
To stop the running container, open a new terminal window and use the following command. Replace <container_id> with the actual container ID or the name of the container you want to stop. You can find the container ID by running docker ps or docker ps -a to list all running and stopped containers.
- **docker stop <container_id>**

### Cleaning Up ###
After stopping the container, you can remove it by running:
- **docker rm <container_id>**

Additionally, you can remove the Docker image to save disk space:
- **docker rmi docker-asciiweb**

**Authors**  
Katrina Mäeorg

**Implementation details: algorithm**  
Implemented with Go server HTTP requests. CSS, HTML for styling. Docker.