#!/bin/sh

# Example ./POST_comment.sh 
URL="http://localhost:8000/process_map"

data='{"path":"./maps/opolskie-latest.osm.pbf"}'

curl -X POST -d "$data" "$URL"  