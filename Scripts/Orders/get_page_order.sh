#!/bin/bash

URL="http://localhost:8000/orders/0_20"

RESPONSE=$(curl -X GET -G "$URL")
echo "$RESPONSE" | jq