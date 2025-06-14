#!/bin/bash
set -e
# Get the root directory (the parent of this script's parent)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="${SCRIPT_DIR}/../.."
cd "$ROOT_DIR"

# Step 1: Download dependencies
echo "Fetching Go dependencies..."
go mod download

# Step 2: Vendor dependencies
echo "Vendoring Go dependencies..."
go mod vendor

# Step 3: Build the app
echo "Building the app..."
go build -o app cmd/main.go

# Step 4: Run the app
echo "Running the app..."
./app
