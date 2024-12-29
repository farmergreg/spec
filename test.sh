#!/bin/env sh

echo "Running Go tests..."
go test ./... -v

# Exit with the same status as the test command
exit $?