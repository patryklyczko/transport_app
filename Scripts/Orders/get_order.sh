#!/bin/bash

URL="http://localhost:8000/orders/od_1685201515060197578-6"

QUERY_PARAMS=${1:-}

# echo "Can specify for example:  id=?"
if [ -z "$QUERY_PARAMS" ]; then
  RESPONSE=$(curl -X GET -G "$URL")
else
  RESPONSE=$(curl -X GET -G "$URL" --data-urlencode "$QUERY_PARAMS")
fi

echo "$RESPONSE" | jq