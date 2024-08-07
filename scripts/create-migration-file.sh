#!/bin/bash

# 첫 번째 인자를 변수에 저장
MIGRATION_NAME=$1

# 마이그레이션 파일 생성
migrate create -ext sql -dir ./internal/migrations -seq $MIGRATION_NAME