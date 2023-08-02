BINARY_NAME=procurementsAPP

build:
	@echo "Building Procurements APP..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Procurements APP built!"

run: build
	@echo "Starting Procurements APP..."
	@./tmp/${BINARY_NAME}
	@echo "Procurements APP started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Procurements APP..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Procurements APP!"

restart: stop start