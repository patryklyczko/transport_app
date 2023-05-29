#!/bin/bash


lat_list=(12 34 43 22 1 2 3 4 45 23 65)
lon_list=(12 23 24 12 3 4 53 2 34 99 1)

length=${#lon_list[@]}

for ((i=0; i<$length; i++)); do
    URL="http://localhost:8000/driver"
    data='{
        "name": "John Doe",
        "position": {
            "lat": '${lat_list[$i]}',
            "lon": '${lon_list[$i]}'
        }
    }'

    curl -X POST -d "$data" "$URL"
done
