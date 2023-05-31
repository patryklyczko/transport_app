#!/bin/bash

lat_list=(15 9 4 52 12 2 32 4 4 21 65)
lon_list=(1 29 4 2 3 44 3 21 4 9 12)

length=${#lon_list[@]}

URL="http://localhost:8000/order"
for ((i=0; i<$length; i++)); do
    data='{
    "position_send": {
        "lat": '${lat_list[$i]}',
        "lon": '${lon_list[$i]}'
    },
    "time_add": "2023-01-02T15:04:05Z",
    "time_end": "2023-01-02T15:04:05Z",
    "gain": 10000
    }'

    curl -X POST -d "$data" "$URL" 
done