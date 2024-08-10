#!/bin/bash

# Define variables
currentDir=$(pwd)
exeName="chuckle"
installDir="$HOME/chuckle-cli"

# ANSI color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Build the Go application
echo -e "${GREEN}Building the Go application...${NC}"
go build -o "$exeName" main.go

if [ $? -ne 0 ]; then
    echo -e "${RED}ERROR: Failed to build the application.${NC}"
    exit 1
fi

# Create the install directory if it doesn't exist
if [ ! -d "$installDir" ]; then
    mkdir -p "$installDir"
    echo -e "${GREEN}SUCCESS: Created directory $installDir${NC}"
fi

# Move the executable to the install directory
echo -e "${GREEN}Moving the executable to $installDir...${NC}"
mv "$currentDir/$exeName" "$installDir/$exeName"

# Add the install directory to the user's PATH if it's not already there
if [[ ":$PATH:" != *":$installDir:"* ]]; then
    echo "export PATH=\$PATH:$installDir" >> "$HOME/.zshrc"
    echo -e "${GREEN}Added $installDir to PATH in .zshrc${NC}"
else
    echo -e "${YELLOW}WARNING: $installDir is already in PATH${NC}"
fi

echo -e "${GREEN}Setup completed. You must restart terminal before you can use command 'chuckle' from the command line."