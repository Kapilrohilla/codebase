#!/bin/bash

function takeInput() {
	read -p "Do you want to continue (Y/N)" inputVal
	if [[ $inputVal == "" ]]; then
		takeInput
	elif [[ ("$inputVal" != 'Y' && "$inputVal" != "N") ]]; then
		echo "Invalid value"
		takeInput
	else
		echo "Correct value : $inputVal"
	fi

}

takeInput
