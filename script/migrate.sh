#!/bin/bash

# Load environment variables from .env file
set -a
source .goose.env
set +a

# Function to run migration commands
run_goose() {
    echo $GOOSE_MIGRATION_DIR
    echo $GOOSE_DRIVER
    echo $GOOSE_DBSTRING


    goose -dir "$GOOSE_MIGRATION_DIR" "$GOOSE_DRIVER" "$GOOSE_DBSTRING" "$@"
}

# Parse command line arguments
case "$1" in
    ("up")
        echo "Running migrations up..."
        run_goose up
        ;;
    ("down")
        echo "Running migrations down..."
        run_goose down
        ;;
    ("status")
        echo "Checking migration status..."
        run_goose status
        ;;
    ("create")
        if [ -z "$2" ]; then
            echo "Usage: $0 create <migration_name>"
            exit 1
        fi
        echo "Creating new migration: $2"
        run_goose create "$2" sql
        ;;
    ("version")
        echo "Checking current version..."
        run_goose version
        ;;
    ("reset")
        echo "Resetting database..."
        run_goose reset
        ;;
    (*)
        echo "Usage: $0 {up|down|status|create <name>|version|reset}"
        echo "Examples:"
        echo "  $0 up                    - Run all pending migrations"
        echo "  $0 down                  - Rollback one migration"
        echo "  $0 status                - Show migration status"
        echo "  $0 create add_users      - Create new migration"
        echo "  $0 version               - Show current version"
        echo "  $0 reset                 - Reset all migrations"
        exit 1
        ;;
esac