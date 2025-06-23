#!/bin/bash

# cat /etc/passwd | awk -F ":" '{print $1}' | while read line; do
#     echo "printing line -> $line"
# done

while read line; do
    echo "$line"
    sleep 0.20
done </etc/passwd | awk -F ":" 'NR > 10 {print $1}'
