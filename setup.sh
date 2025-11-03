#!/bin/bash

# Stop on any error
set -e

echo "Installing project dependencies..."
go mod tidy

echo "Generate encryption key..."
go run . artisan key:generate

echo "Setting up database, migrations and seeding..."
go run . artisan migrate:fresh --seed