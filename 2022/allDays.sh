#!/usr/bin/env bash

for i in $(ls -d day* | grep -v ".sh$" | cut -d"y" -f2); do echo Day $i && ./day.sh $i; done
