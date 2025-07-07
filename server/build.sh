go build
docker build --tag server .
docker run --env-file=./.env server