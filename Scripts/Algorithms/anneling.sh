#!/bin/sh

# Example ./POST_comment.sh 
URL="http://localhost:8000/simulated_anneling"
data='{
  "t_init":1000,
  "cooling":0.90,
  "t_end":1,
  "n_max":100,
  "k":10
}'

curl -X POST -d "$data" "$URL" 