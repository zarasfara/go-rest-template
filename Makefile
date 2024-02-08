init:
	cp .env.example .env
	git config --local core.hooksPath githooks
start:
	docker compose up -d --build