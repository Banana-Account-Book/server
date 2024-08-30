#!/bin/bash

NUM_DOWN=${1:-1}

# Run the migrate command in local environment
migrate -database "postgres://postgres:1234@localhost:5432/banana?sslmode=disable" -path ./internal/migrations down $NUM_DOWN