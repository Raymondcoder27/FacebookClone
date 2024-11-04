#!/bin/sh
if [ -f /usr/local/bin/.env ]; then
    export $(grep -v '^#' /usr/local/bin/.env | xargs)
fi

#Start the backend application
/usr/local/bin/main &

#Start nginx
serve -s /usr/share/nginx/html -l 80
