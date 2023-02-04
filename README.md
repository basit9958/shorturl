# ShortUrl 

Golang based url shortening service

## Library Used :
1. Golang fiber
2. Google UUID
3. Redis Client


This binary will be able to support custom configuration of redis database from .env. For now users can define their own short URL or can allow the program to define . 

You would get the response as :-
````
 Long URL
 Short URL
 Expiration Time of the key
````
Default expiration time of the key is set to be 20 minute

## How to use it
- Connect it to the redis server use .env 
- Use postMan or Curl command to send POST request

```bash
CURL - curl -H "Content-Type: application/json" -X POST -d '{"url":"www.youtube.com"}' http://localhost:3000/api/
```

How to Deploy it :
- You can use docker-compose up -d to deploy it using docker 
- You can use the local host but then you need to connect your redis database to the .env

Sample Request :
- API requests can be sent using POSTMAN or CURL 
- Once the path is set use 
 ```bash 
  localhost:3000/{shortURL}
  ```
- To get the long URL again use 
```bash
localhost:3000/get/{shortURL}
```
Example :
```bash
Request : curl -H "Content-Type: application/json" -X POST -d '{"url":"www.youtube.com"}' http://localhost:3000/api/
Response : {"url":"www.youtube.com","short":"localhost:3000/0eec2g","time":20}
```

```bash
Request : curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:3000/get/0eec2g
Response :
HTTP/1.1 200 OK
Date:  Wed, 1 feb 2023 20:31:56 GMT
Content-Type: application/json
Content-Length: 50

{"url":"www.youtube.com","short":"0eec2g","time":0}
```
```bash
Request : curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:3000/0eec2g
Response :
HTTP/1.1 308 Permanent Redirect
Date:  Wed, 1 feb 2023 20:32:29 GMT
Content-Length: 0
```
