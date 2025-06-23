#!/bin/bash

# for variableName in item1 item2 item3; do
#     echo "$variableName"
# done

read -p "Please enter your number: " num

for i in {1..10}; do
    echo $((num * i))
done
