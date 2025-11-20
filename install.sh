#!/bin/bash

# Exit on error
set -e

go run . artisan key:generate

go run . artisan jwt:secret

echo "Running database migrations and seeders..."

go run . artisan migrate:fresh --seed

echo "Done!"