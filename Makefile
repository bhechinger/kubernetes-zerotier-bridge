GENERATOR_VERSION=latest
USERID = $(shell id -u)
GROUPID = $(shell id -g)
REPOROOT = $(shell git rev-parse --show-toplevel)
GOFMT = $(shell which gofmt)

generate:
    # Remove all generated files to re-generate, except for bazel build file,
    # which is needed by bazel to create Docker images.
	rm -rf ${REPOROOT}/routes/zerotier
	docker run --rm -u "$(USERID):$(GROUPID)" \
    	-e GO_POST_PROCESS_FILE="${GOFMT} -w" \
    	-v "${REPOROOT}:/local" openapitools/openapi-generator-cli:${GENERATOR_VERSION} generate \
		-i https://apidocs.zerotier.com/v1/api-spec.json \
		-g go \
		--generate-alias-as-model \
		--instantiation-types int64=int64 \
		-o /local/routes/zerotier
	rm -f routes/zerotier/go.sum routes/zerotier/go.mod routes/zerotier/git_push.sh
	find ./ -type f -exec sed -i 's/int64/int64/g' {} \;

