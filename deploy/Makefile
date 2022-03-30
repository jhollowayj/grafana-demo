RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

DEPLOY_FOLDER := $(shell pwd)

network:
	docker network create grafana-demo

# Usage:
# `make metrics -- up --detach`
# `make metrics down`
metrics: 
	DEPLOY_FOLDER=${DEPLOY_FOLDER}/configs docker-compose -f metrics.docker-compose.yaml $(RUN_ARGS)

# Usage:
# `make metrics -- up --detach`
# `make metrics down`
gobin:
	docker-compose -f single.docker-compose.yaml $(RUN_ARGS)

