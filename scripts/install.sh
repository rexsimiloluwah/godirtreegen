#!/bin/bash

# Remove any existing installation of godirtreegen 
which godirtreegen > /dev/null 2>&1
if [ $? -eq 0 ]; then 
    rm -rf $(which godirtreegen) > /dev/null
fi

TARGET_DIR=""

# Check if the $GOBIN variable is set 
if [[ ! -z "$GOBIN" ]]; then 
    TARGET_DIR=$GOBIN
fi 

# Check if the $GOPATH variable is set
if [[ ! -z "$GOPATH" ]]; then
    TARGET_DIR=$GOPATH/bin 
fi 

# If either $GOPATH and $GOBIN are not set 
if [ -z "$TARGET_DIR" ]; then 
    TARGET_DIR=/usr/local/bin  
fi

# Check if the directory exists 
if [ ! -d "$TARGET_DIR" ]; then 
    echo "Seems $TARGET_DIR does not exist."
    echo "Don't bother....Creating new directory: $TARGET_DIR"
    mkdir -p "$TARGET_DIR"
fi

# Install the executable binary in the $TARGET_DIR/godirtreegen folder
echo "Installing in $TARGET_DIR/godirtreegen, wait a second..."

go build -o $TARGET_DIR/godirtreegen main.go 

# Check if installation failed
if [ $? -ne 0 ]; then 
    echo "Failed to compile and install executable binary."
    exit 1 
fi 

# Check if the binary was installed globally 
which godirtreegen > /dev/null 2>&1
if [ $? -ne 0 ]; then 
    echo "Failed to install godirtreegen globally."
    exit 1
fi

# Check if the binary runs correctly 
godirtreegen --help
if [ $? -ne 0 ]; then 
    echo "Godirtreegen is not working properly, unfortunately."
else 
    echo "-----------------------------------------"
    echo "Successfully installed godirtreegen."
fi
