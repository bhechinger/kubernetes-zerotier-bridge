VERSION=0.0.21
USER=wonko
NAME=kubernetes-zerotier-bridge

docker:
	docker buildx build -t ${USER}/${NAME}:${VERSION} --push .