GENERATOR_VERSION=5.0.0
USERID = $(shell id -u)
GROUPID = $(shell id -g)
REPOROOT = $(shell git rev-parse --show-toplevel)

generate:
    # Remove all generated files to re-generate, except for bazel build file,
    # which is needed by bazel to create Docker images.
	rm -rf ${REPOROOT}/routes/zerotier
	docker run --rm -u "$(USERID):$(GROUPID)" \
    	-v "${REPOROOT}:/local" openapitools/openapi-generator-cli:v${GENERATOR_VERSION} generate \
		-i https://apidocs.zerotier.com/v1/api-spec.json \
		-g go \
		--generate-alias-as-model \
		-o /local/routes/zerotier
	rm -f routes/zerotier/go.sum routes/zerotier/go.mod routes/zerotier/git_push.sh
