# ==============================================================================
# Makefile helper functions for swagger
#

.PHONY: swagger.run
swagger.run: tools.verify.swagger
	@echo "===========> Generating swagger API docs"
	#@swagger generate spec --scan-models -w $(PROJ_ROOT_DIR)/cmd/gen-swagger-type-docs -o $(PROJ_ROOT_DIR)/api/swagger/kubernetes.yaml
	@swagger mixin `find $(PROJ_ROOT_DIR)/api/openapi -name "*.swagger.json"` \
		-q                                                    \
		--keep-spec-order                                     \
		--format=yaml                                         \
		--ignore-conflicts                                    \
		-o $(PROJ_ROOT_DIR)/api/swagger/swagger.yaml
	@echo "Generated at: $(PROJ_ROOT_DIR)/api/swagger/swagger.yaml"

.PHONY: swagger.serve
swagger.serve: tools.verify.swagger
	@swagger serve -F=redoc --no-open --port 65534 $(PROJ_ROOT_DIR)/api/swagger/swagger.yaml
