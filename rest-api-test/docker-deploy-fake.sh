# docker image build
docker build \
    --no-cache \
    --tag rest-api-test:latest .

# build docker image check
# docker images | grep app-hello-gin

# execute docker image
# docker run --rm \
#     --publish 8080:8080 \
#     --name app-local \
#     rest-api-test:latest

docker run --rm \
    --publish 8080:8080 \
    --name rest-api-test \
    {docker-image-url-with-tag}

# 요거 사용하자
docker run -d --restart=always \
    --publish 8080:8080 \
    --name rest-api-test \
    {docker-image-url-with-tag}