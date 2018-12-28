TARGET              := elementary-cacher
DOCKER_IMAGE_NAME   := elementary-cacher
GO                  := go
GIT_SUMMARY         := sleep

build:
	@echo ">> building binary..."
	@echo ">> GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GOFLAGS) -o $(TARGET)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GOFLAGS) -o $(TARGET) cmd/cacher/main.go

docker: Dockerfile
	@echo ">> building docker image..."
	@echo "docker build -t $(DOCKER_IMAGE_NAME):$(GIT_SUMMARY) ."
	@docker build -t $(DOCKER_IMAGE_NAME):$(GIT_SUMMARY) .

docker-push:
	@echo ">> docker login..."
	@echo ">> docker login -u vlamug"
	@docker login -u vlamug
	@echo ">> docker tag..."
	@echo ">> docker tag $(DOCKER_IMAGE_NAME):$(GIT_SUMMARY) vlamug/$(DOCKER_IMAGE_NAME):$(GIT_SUMMARY)"
	@docker tag $(DOCKER_IMAGE_NAME):$(GIT_SUMMARY) vlamug/$(DOCKER_IMAGE_NAME):$(GIT_SUMMARY)
	@echo ">> docker push..."
	@echo ">> docker push vlamug/$(DOCKER_IMAGE_NAME):$(GIT_SUMMARY)"
	@docker push vlamug/$(DOCKER_IMAGE_NAME):$(GIT_SUMMARY)
