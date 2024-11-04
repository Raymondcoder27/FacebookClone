#!/bin/sh
if [ -f /usr/local/bin/ .env ]; then
    export $(grep -v '^#' /usr/local/bin/.env | xargs)

#Start the backend application
/usr/local/bin/main &

#Start nginx
serve -s /usr/share