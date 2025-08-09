#!/bin/bash
set -e

# Resolve script dir to load .goose.env reliably
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

set -a
source "$SCRIPT_DIR/.goose.env"
set +a

# Normalize migration dir if it's relative
case "$GOOSE_MIGRATION_DIR" in
  (/*) ;;  # absolute, keep
  (*) GOOSE_MIGRATION_DIR="$(cd "$SCRIPT_DIR/.." && pwd)/${GOOSE_MIGRATION_DIR#./}";;
esac

run_goose() {
  # New goose CLI: goose -dir DIR <command> (driver/dbstring từ env)
  GOOSE_DRIVER="$GOOSE_DRIVER" GOOSE_DBSTRING="$GOOSE_DBSTRING" \
    goose -dir "$GOOSE_MIGRATION_DIR" "$@"
}

case "$1" in
  (up|down|status|version|reset)
    echo "Running migrations $1..."
    run_goose "$1"
    ;;
  (create)
    [ -z "$2" ] && { echo "Usage: $0 create <name>"; exit 1; }
    echo "Creating new migration: $2"
    run_goose create "$2" sql
    ;;
  (*)
    echo "Usage: $0 {up|down|status|create <name>|version|reset}"
    exit 1
    ;;
esac