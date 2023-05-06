#!/bin/sh

for i in $(seq 1 5)
do
    URL="http://localhost:8000/driver"
    data='{
        "name": "John Doe",
        "position": {
            "latitude": 0.0,
            "longitude": 0.0
        }
    }'
    
    curl -X POST -d "$data" "$URL" 
done
