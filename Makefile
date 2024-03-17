#!make

# VARIABLES
PACKAGE="github.com/IsaqueRocha/greenlight"
# ENVFILE=development.env

# IMPORT ENV VARIABLES
# include $(ENVFILE)
# export $(shell sed 's/=.*//' $(ENVFILE))

# TARGETS
default: usage

clean: ## clean: Trash binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@echo "Clean OK"

test: ## test: Run all tests
	@echo "--> testing..."
	@go test -v $(PACKAGE)/...

run:  ## run: Run your application
	@echo "--> running application..."
	@go run ./cmd/api

air: ## air: Run your application with air plugin
	@echo "--> running application..."
	@air

# env: ## env: Simple test to verify ENV load
# 	@echo "--> showing var in $(ENVFILE) ..."
# 	@cat $(ENVFILE)
# 	@echo "--> loading env vars..."
# 	@echo $(APPLICATION)

usage: ## usage: List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make\033[0m %s\n", $$2}'