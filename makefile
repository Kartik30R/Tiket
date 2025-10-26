start:
	@docker compose up --build

stop:
	@docker compose down -v
	@docker rmi ticket-booking
