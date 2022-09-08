PKG_LIST 	 := $(shell go list .market-api/internal/... | grep -v /vendor/)

up: ## Run market
	docker-compose up
logs: ## Get logs container
	rm market.txt
	docker logs market >> market.txt
test: ## Run unittests
	@rm -rf coverage
	@go test -short ${PKG_LIST}
coverage: ## Generate global code coverage report
	rm -rf coverage
	./market-api/tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	rm -rf coverage
	./market-api/tools/coverage.sh html;