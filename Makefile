PUBLISH_VERSION=$(shell echo ${NEW_VERSION} | sed 's/inner-999/1/g')


build:
	docker build --tag rookout/tutorial-go:latest --tag rookout/tutorial-go:${PUBLISH_VERSION} . --build-arg ARTIFACTORY_CREDS=${JFROG_ARTIFACTORY_CREDS}

upload-no-latest:
	docker push rookout/tutorial-go:${PUBLISH_VERSION}

upload: upload-no-latest
	@if [ ${CIRCLE_BRANCH} = "master" ]; then \
		docker push rookout/tutorial-go:latest; \
	fi

build-and-upload: build upload