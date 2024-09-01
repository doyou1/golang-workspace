# docker image build
docker build \
    --no-cache \
    --tag rest-api-test:latest .

# 요거 사용하자
# docker build -t rest-api-test:lastest .

# check docker image
docker images | grep rest-api-test

# docker image push
docker push {dockerhub-link}/rest-api-test:latest

# docker image excute
docker run --rm \
    --publish 8080:8080 \
    --name rest-api-test \
    {docker-image-url-with-tag}

# 요거 사용하자
docker run -d --restart=always \
    --publish 8080:8080 \
    --name rest-api-test \
    {docker-image-url-with-tag}