#!/bin/bash

# while [[ $answer != "yes" ]]
# do 
#	read -p "please enter yes: " answer
#	echo "you enter $answer"
# done



initNumber=1

while [[ initNumber -le 10 ]] 
do

	echo $((initNumber * 2))
	((initNumber++))
done
