up: ## Run market
	docker-compose up
logs: ## Get logs container
	rm market.txt
	docker logs market >> market.txt