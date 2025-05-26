#!/bin/bash
set -eu

BLUE='\E[1;34m'
YELLOW='\E[1;33m'
RESET='\E[0m'
RESULT=0

if ! command -v golangci-lint &>/dev/null; then
	echo -e "${YELLOW}Installing golangci-lint ...${RESET}"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
fi

echo -e "${YELLOW}golangci-lint ...${RESET}"
golangci-lint --max-same-issues=0 --max-issues-per-linter=0 run ./...
