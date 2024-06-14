#!/bin/sh
# set -e
echo "Run database migrations"
/app/migrate -path ./migration -database "$dbsource" -verbose up
echo "Start the application"
/app/main


