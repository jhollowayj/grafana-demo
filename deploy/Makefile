# This section allow me to pass in --detach flags to docker compose
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

DEPLOY_FOLDER := $(shell pwd)

network:
	docker network create grafana-demo

# Usage:
# `make metrics -- up --detach`
# `make metrics down`
metrics: 
	docker-compose -f metrics.docker-compose.yaml $(RUN_ARGS)

# Usage:
# `make apps -- up --detach`
# `make apps down`
apps:
	docker-compose -f service.docker-compose.yaml $(RUN_ARGS)

