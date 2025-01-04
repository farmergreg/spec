#!/bin/bash

# Create output directory for coverage reports
mkdir -p coverage

# Run tests with coverage
go test ./... -coverprofile=coverage/coverage.out

# Generate HTML coverage report
go tool cover -html=coverage/coverage.out -o coverage/coverage.html
