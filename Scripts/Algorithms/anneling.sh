#!/bin/sh

# Example ./POST_comment.sh 
URL="http://localhost:8000/simulated_anneling"
data='{
  "t_init":1000,
  "cooling":0.9,
  "t_end":1,
  "n_max":100,
  "k":5
}'

curl -X POST -d "$data" "$URL" 