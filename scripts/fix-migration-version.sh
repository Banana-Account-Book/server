#!/bin/bash

# 첫 번째 인자를 변수에 저장
MIGRATION_VERSION=$1

# 마이그레이션 
migrate -database "postgres://postgres:1234@localhost:5432/banana?sslmode=disable" -path ./internal/migrations force $MIGRATION_VERSION