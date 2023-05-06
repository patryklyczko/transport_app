#!/bin/sh

# Example ./POST_comment.sh 
URL="http://localhost:8000/driver"

for ((i=0; i<100; i++))
do
    data='{
    "position": {
        "latitude": 0.0,
        "longitude": 0.0
    },
    "time_add": "2006-01-02T15:04:05Z",
    "time_end": "2006-01-02T15:04:05Z",
    "gain": 0
    }'

    curl -X POST -d "$data" "$URL" 
done
