#!/bin/bash

DOCKERD_CMD="sudo dockerd -H unix:///var/run/docker.sock -H tcp://127.0.0.1"

start_dockerd() {
    echo "Starting dockerd..."
    $DOCKERD_CMD &
    echo "dockerd started."
}

if pgrep -x "dockerd" > /dev/null; then
    echo "dockerd is running. Stopping dockerd..."
    sudo pkill -x "dockerd"
    echo "dockerd stopped."
    sleep 2
else
    echo "dockerd is not running."
fi




start_dockerd

# sudo dockerd -H unix:///var/run/docker.sock -H tcp://127.0.0.1