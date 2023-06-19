APP=$(shell git remote get-url origin | xargs basename | cut -d '.' -f 1)
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
TARGETARCH=$(shell dpkg --print-architecture)
# LINUX_TARGETOS=linux
# MACOS_TARGETOS=darwin
# WINDOWS_TARGETOS=windows
# DOCKERHUB_REGISTRY=vitalibit
OWNER=vitalibit
REGISTRY=ghcr.io
# PROJECT_ID=k8s-k3s-386218
# GCR_IMAGE_TAG=${REGISTRY}/${PROJECT_ID}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}
IMAGE_TAG=${REGISTRY}/${OWNER}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="github.com/vitalibit/kbot/cmd.appVersion=${VERSION}

# linux:
# 	$(MAKE) build TARGETOS=$(LINUX_TARGETOS)

# macOS:
# 	$(MAKE) build TARGETOS=$(MACOS_TARGETOS)

# windows:
# 	$(MAKE) build TARGETOS=$(WINDOWS_TARGETOS)

# image_dockerhub:
# 	docker build . -t ${DOCKERHUB_REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

# image-gcr:
# 	docker build . -t ${GCR_IMAGE_TAG}

# push_dockerhub:
# 	docker push ${DOCKERHUB_REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

# push_gcr:
# 	gcloud auth login
# 	gcloud config set project ${PROJECT_ID}
# 	gcloud auth configure-docker
# 	docker push ${GCR_IMAGE_TAG}

# clean_gcr:
# 	rm -rf kbot
# 	docker rmi ${GCR_IMAGE_TAG}

image: 
	docker build . -t ${IMAGE_TAG}

push:
	docker push ${IMAGE_TAG}

clean:
	rm -rf kbot
	docker rmi ${IMAGE_TAG}