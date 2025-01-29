# Makefile
GOEXEC = go
COVERAGE_REPORT = coverage.out
TEST_REPORT = report.out
TEST_FILES = ./internal/...
COMPOSE_FILE = ./docker-compose.yml
DOCKER_COMPOSE = docker compose -f "${COMPOSE_FILE}"

.PHONY: test
test:
	@echo "Running tests..."
	${GOEXEC} test -v ${TEST_FILES}

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	${GOEXEC} test -bench=. -json -coverprofile="${COVERAGE_REPORT}" ${TEST_FILES} > "${TEST_REPORT}"

.PHONY: test-integration
test-integration:
	@echo "Running integration tests..."
	${GOEXEC} test -v ${TEST_FILES}

# WIP: This target is not working yet
.PHONY: test-integration-pipeline
test-integration-pipeline:
	@echo "Running integration tests..."
	make infra-up
	${GOEXEC} test -v ${TEST_FILES}
	make infra-down

.PHONY: serve-dev
serve:
	@echo "Starting server..."
	${GOEXEC} run cmd/server/main.go serve --env=dev

.PHONY: infra-up
infra-up:
	@echo "Starting infrastructure..."
	${DOCKER_COMPOSE} up -d

.PHONY: infra-down
infra-down:
	@echo "Stopping infrastructure..."
	${DOCKER_COMPOSE} down