# drone-docker-buildx

[![Build Status](https://drone.owncloud.com/api/badges/owncloud-ci/drone-docker-buildx/status.svg)](https://drone.owncloud.com/owncloud-ci/drone-docker-buildx)
[![Docker Hub](https://img.shields.io/docker/v/owncloudci/drone-docker-buildx?logo=docker&label=dockerhub&sort=semver&logoColor=white)](https://hub.docker.com/r/owncloudci/drone-docker-buildx)
[![GitHub contributors](https://img.shields.io/github/contributors/owncloud-ci/drone-docker-buildx)](https://github.com/owncloud-ci/drone-docker-buildx/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/owncloud-ci/drone-docker-buildx)
[![License: Apache-2.0](https://img.shields.io/github/license/owncloud-ci/drone-docker-buildx)](https://github.com/owncloud-ci/drone-docker-buildx/blob/main/LICENSE)

Drone plugin to build multiarch Docker images with buildx. This plugin is a fork of [drone-plugins/drone-docker](https://github.com/drone-plugins/drone-docker).

## Usage

> **IMPORTANT:** Be aware that the this plugin requires [privileged](https://docs.drone.io/pipeline/docker/syntax/steps/#privileged-mode) capabilities, otherwise the integrated Docker daemon is not able to start.

```yaml
kind: pipeline
type: docker
name: default

steps:
  - name: docker
    image: owncloudci/drone-docker-buildx
    privileged: true
    settings:
      username: octocat
      password: secure
      repo: octocat/example
      tags: latest
```

## Build

Build the binary with the following command:

```console
make build
```

Build the Docker image with the following command:

```console
docker build --file Dockerfile.multiarch --tag owncloudci/drone-docker-buildx .
```

## Test

```console
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=00000000 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  owncloudci/drone-docker-buildx --dry-run
```

## Releases

Create and push the new tag to trigger the CI release process:

```console
git tag v2.10.3
git push origin v2.10.3
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](https://github.com/owncloud-ci/drone-docker-buildx/blob/main/LICENSE) file for details.

## Copyright

```text
Copyright (c) 2022 ownCloud GmbH
```
