# Bytely

It is an API that receives an URL from client and returns a "shortened" url under `localhost` domain.

### Usage
You would need Docker Compose to run this service. Following command will build local image of the service and pull postgres image and start a docker-compose application. Database schema will automatically be created. 

```
$ make up
```

### Configurations
Configurations such as ports can be changed in `docker-compose.yml` file. Currently client server is started on port `80` and GRPC API server is started on port `8080`. 

### Documentation
Documentation uses Swagger UI to display endpoint descriptions. The documentation can be found at [http://localhost/swagger/index.html](http://localhost/swagger/index.html).


### Features / User stories
  - [x] Any user can post a link and get the shortened url. (Same URL can be posted multiple times
	  and receive different shortened URLs)
  - [x] Any user can delete a shortened url.
  - [x] Any user can enter a shortened url and redirected to the original url. 
  - [x] Any user can get stats for the shortened URL. 
