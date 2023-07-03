# Microservices Learning Project

This repository is a testbed for learning and understanding the practical implementation of microservices architecture using Go.

The project is comprised of six different services, each running in its own Docker container, demonstrating the modularity and the independence of components in a microservices architecture.

## Services
1. **Broker-Service**: A broker service that listens to requests and redirects them to the appropriate service based on the action in the request. It's capable of using RPC, gRPC, and RabbitMQ for communication between the services. Although handlers for all three are available in the codebase, currently gRPC is the active communication protocol.
2. **Authentication-Service**: This service handles authentication requests. The credentials provided are compared to the ones stored in a PostgreSQL database, and a response is returned based on whether the authentication is successful or not.
3. **Logger-Service**: This service handles logging requests, and storing logs in a MongoDB database.
4. **Listener-Service**: A service that listens to different actions/events taking place and responds accordingly.
5. **Mail-Service**: This service handles sending emails. For the sake of this project, it sends emails to a Mailhog server, which is a dummy SMTP server for developers.
6. **Front-End**: This is the user interface of the project. It's a web page where you can trigger actions that result in requests being sent to the broker service. The payload of the request and the response from the server are displayed on the page.

For easier management, the project includes a `docker-compose` file to start all the services in one command and a `Makefile` for building the binaries for each service.

The source code for each service can be found in its respective directory. Each service is a Go application with its own Dockerfile, which builds a Docker image for that service.

Please note that this project is purely for learning purposes and is not intended for production use.
