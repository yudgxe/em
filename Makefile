
start: swag migrate serve

swag:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init

migrate:
	go run ./main.go migrate up

serve:
	go run ./main.go serve
