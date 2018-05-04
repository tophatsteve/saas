# saas

Sausage as a service. 

A simple service used to try out various cloud infrastructure services.

### Run the service

To run the service
```bash
go run service/cmd/main.go
```

To run the client
```bash
go run main.go
```

### Create a container image

Top create a docker container and push it to the Docker Hub registry

```bash
docker build . -t tophatsteve/saas:latest
docker push tophatsteve/saas:latest
```