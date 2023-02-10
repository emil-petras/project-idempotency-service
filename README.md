# project-idempotency-service

## list of related repositories
https://github.com/emil-petras/project-web-service
https://github.com/emil-petras/project-db-service
https://github.com/emil-petras/project-proto

this service depends on redis

## run using docker compose
docker compose up -d

## run using docker file
docker pull redis
docker run --name redisdb -p 6379:6379 -d redis

docker build --rm -t project-idempotency-service . 
docker run -p 9002:9002 -d project-idempotency-service