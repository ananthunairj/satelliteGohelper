     Rocket Visualizer
 ----------------------------

This project includes two components:
1. **Go Service (Dockerized)**: A backend service packaged as a Docker image.
2. **Electron Application**: A cross-platform desktop application built with Electron.



In this release, you'll find the following files:

- **goservice.tar**: A Docker image containing the Go backend service.
- **rocketStimulator.zip**: A packaged version of the Electron desktop application for Windows. (or `.tar.gz` for other platforms)


### 1. **Running the Go Service (Dockerized)**

To run the Go service as a Docker container:

1. **Install Docker**: If Docker is not already installed, download and install it from [here](https://www.docker.com/products/docker-desktop).

2. **Load the Docker Image**:
   After downloading the `go-service-myservice.tar` file, load it into Docker:
   ```bash
   docker load -i go-service-myservice.tar

   docker run -d -p 8080:8080 goservice:1.0.0

3. **Extract the electron app**:
     tar -xvzf rocketStimulator.tar.gz 
     
     and run the .exe file



     Dont forget to install the node modules😒

