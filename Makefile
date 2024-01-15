include Makefile.variables

.PHONY: format todo test check prepare
## prefix before other make targets to run in your local dev environment
local: | prepare
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

format: prepare
	${DOCKRUN} bash ./scripts/format.sh

check: prepare format
	${DOCKRUN} bash ./scripts/check.sh

todo:
	${DOCKRUN} bash ./scripts/todo.sh

test: check db_prepare
	${DOCKTEST} bash ./scripts/test.sh

db_start: db_stop
	@docker run --name apniuniversity-mongodb-1 -p 27015-27017:27015-27017

db_stop:
	bash ./scripts/db_stop.sh

codegen: prepare
	${DOCKRUN} bash ./scripts/swagger.sh