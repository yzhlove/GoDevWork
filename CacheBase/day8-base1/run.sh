#!/bin/bash

go build -o rpc_server
./rpc_server -port=8001 &
./rpc_server -port=8002 &
./rpc_server -port=8003 -api=true &
wait
