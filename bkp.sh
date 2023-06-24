#!/bin/bash

# MySQL Connection Details
DB_HOST="168.119.50.201"
DB_USER="root"
DB_PASS=""
DB_PORT="3307"
DB_NAME="kline"

# Output file
OUTPUT_FILE="bkp.sql"

# mysqldump -h 168.119.50.201 -u root -P 3305 -p kline > Documentos/CodeRepositories/klineService/bkp.sql
# Function to perform the backup
perform_backup() {
    mysqldump -h $DB_HOST -u $DB_USER -P $DB_PORT -p $DB_NAME > $OUTPUT_FILE
}

# Function to import the backup
perform_import() {
    mysql -u igor -p123456 kline < $OUTPUT_FILE
}

# Function to handle reconnection
handle_reconnection() {
    while true; do
        echo "Attempting to reconnect..."
        perform_backup
        if [ "$?" -eq 0 ]; then
            echo "Backup completed successfully."
            perform_import
            if [ "$?" -eq 0 ]; then
                echo "Import completed successfully."
                break
            else
                echo "Import failed. Retrying in 5 seconds..."
            fi
        else
            echo "Backup failed. Retrying in 5 seconds..."
        fi
        sleep 5
    done
}

# Perform backup with reconnection handling
handle_reconnection
