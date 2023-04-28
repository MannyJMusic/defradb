#!/bin/sh

/defradb start --url 0.0.0.0:9181 &

PID_DEFRA=$!

echo $PID_DEFRA

echo "Process 1 lasts for 5s" && sleep 5 &

PID=$!


wait $PID

/defradb client schema add --url 0.0.0.0:9181 -f /usr/ragtag_schemas/ragtag_track_schema.graphql
/defradb client schema add --url 0.0.0.0:9181 -f /usr/ragtag_schemas/ragtag_user_auth_schema.graphql

DIP=$!

echo "Process 2 lasts for 5s" && sleep 5 &

wait $DIP


kill $PID_DEFRA



/defradb start --url 0.0.0.0:9181