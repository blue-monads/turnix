#!/bin/bash
set -e

pwd
ls -la
echo "Starting cloudy"
# Start cloudy with the provided arguments
exec ./cloudy "$@"
echo "Cloudy exited with code $?"
sleep Infinity
