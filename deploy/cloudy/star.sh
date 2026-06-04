#!/bin/bash

set -e

echo $PWD

ls -lah

echo "starting cloudy"

exec ./cloudy

echo "cloudy stopped"

sleep INFINITY