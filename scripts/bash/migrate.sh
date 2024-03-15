#!/bin/bash

DB_NAME="postgres"
DB_USER="postgres"
DB_HOST="localhost"
DB_PORT="5433"
DB_PASSWORD="postgress"
TARGET_DIRECTORY="./scripts/sql"

show_help() {
    echo "Usage: $0 -d <database_name> -u <user> -h <host> -p <port> -w <password> -t <target_directory>"
    echo ""
    echo "  -d  Database name"
    echo "  -u  Database user"
    echo "  -h  Database host (default: localhost)"
    echo "  -p  Database port (default: 5432)"
    echo "  -w  Database password (consider using environment variables for security)"
    echo "  -t  Target directory containing .psql or .sql files"
    echo ""
    echo "Example: $0 -d mydatabase -u myuser -h localhost -p 5432 -w mypassword -t /path/to/sql/files"
}

while getopts ":d:u:h:p:w:t:" opt; do
    case ${opt} in
    d) DB_NAME=$OPTARG ;;
    u) DB_USER=$OPTARG ;;
    h) DB_HOST=$OPTARG ;;
    p) DB_PORT=$OPTARG ;;
    w) DB_PASSWORD=$OPTARG ;;
    t) TARGET_DIRECTORY=$OPTARG ;;
    \?)
        show_help
        exit 1
        ;;
    esac
done

if [ -z "$DB_NAME" ] || [ -z "$DB_USER" ] || [ -z "$DB_PASSWORD" ] || [ -z "$TARGET_DIRECTORY" ]; then
    echo "Error: Missing required arguments."
    show_help
    exit 1
fi

if [ ! -d "$TARGET_DIRECTORY" ]; then
    echo "Error: Target directory does not exist."
    exit 1
fi

cd "$TARGET_DIRECTORY"

find . -type f \( -iname \*.psql -o -iname \*.sql \) | while read sqlfile; do
    echo "Executing $sqlfile..."
    PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f "$sqlfile"
    if [ $? -eq 0 ]; then
        echo "$sqlfile executed successfully."
    else
        echo "Error occurred during execution of $sqlfile"
    fi
done

echo "All scripts have been executed."
