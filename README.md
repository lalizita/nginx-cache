# NGINX Cache basics

Simple app to study about the basics for cache requests using NGINX.
Common topics that is applied in this repo:
- Reverse Proxy
- Cache configurations
- Bypass requests

### Run project
This projects has a simple Golang app and a simple NGINX config.
```
docker-compose up
```

### Test project
to make requests for NGINX app

```
curl -X GET 'http://localhost:8081'
```