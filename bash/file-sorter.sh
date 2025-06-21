#!/bin/bash

RED="\e[32m"
ENDCOLOR="\e[0m"
imageDir="Photos"
docsDir="Document"

function readPath() {
    read -p "Enter the directory path: " dir
    echo $dir
}

function validatePath() {
    if [ -d "$1" ]; then
        echo true
    else
        echo false
    fi
}

readDirContent() {
    data=(ls $1)

    echo data
}

function createBucketDirs() {
    if [ ! -d "Document" ]; then
        mkdir Document
    fi

    if [ ! -d "Images" ]; then
        mkdir Images
    fi
}

function readNMoveImageFile() {
    for file in "${dirPath}"/*.{png,jpg,jpeg,webp}; do
        if [ -e "$file" ]; then
            mv $file "Images/"
        else
            continue
        fi
    done
}

function readNDocFile() {
    for file in "${dirPath}"/*.{pdf,html,docx}; do
        if [ -e "$file" ]; then
            # echo "$file"
            mv "$file" Document
        else
            continue
        fi
    done

}

dirPath=$(readPath)
isValidPath=$(validatePath "${dirPath}")
createBucketDirs
readNMoveImageFile
