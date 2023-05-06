#!/bin/sh

URL="http://localhost:8000/order"

data='{"id":"od_1683359699797216804-8"}'

curl -X DELETE -d "$data" "$URL"