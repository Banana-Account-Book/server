#!/bin/bash

# Run the migrate command in local environment
migrate -database "postgres://postgres:1234@localhost:5432/banana?sslmode=disable" -path ./internal/migrations down