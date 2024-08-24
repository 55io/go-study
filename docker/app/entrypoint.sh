#!/bin/sh
# Create config
# if [ ! -f /app/config/config.yaml ];
# then cp /app/config/config.example.yaml /app/config/config.yaml;
# echo "Создали config.yaml";
# fi
# Run service
go run ./cmd/service/main.go ./config/config.yaml