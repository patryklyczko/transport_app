#!/bin/bash

URL="http://localhost:8000/drivers"

RESPONSE=$(curl -X GET -G "$URL")
echo "$RESPONSE" | jq