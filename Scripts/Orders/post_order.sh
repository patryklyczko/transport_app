#!/bin/sh

# Example ./POST_comment.sh 
URL="http://localhost:8000/order"
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