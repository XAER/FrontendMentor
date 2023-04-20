#!/bin/bash

# Script to update the challenges.json file

# Check if the challenges.json file exists
if [ -f ../../../challenges.json ]; then
    echo "challenges.json file exists"
else
    echo "challenges.json file does not exist"
    exit 1
fi

# Get a list of all the directories in the main folder (../../../)
# and store it in a variable
declare -a directories=($(ls ../../../))

directoriesToAdd=()
directoriesToIgnore=(".git", ".github", "scripts")

# Loop through the directories
for directory in "${directories[@]}"
do 
    # Check if the directory is a directory
    if [ -d ../../../$directory ]; then
        # Check if the directory is in the ignore list
        if [[ " ${directoriesToIgnore[@]} " =~ " ${directory} " ]]; then
            echo "Ignoring $directory"
            continue
        else
            echo "Adding $directory"
            directoriesToAdd+=($directory)
        fi
    else
        continue
    fi
done

echo "Directories to add: ${directoriesToAdd[@]}"

# Creating a new challenges.json file
echo "{" > challenges.json
echo "  \"challengesDirectories\": [" >> challenges.json

# Loop through the directories to add
for directory in "${directoriesToAdd[@]}"
do
    echo "    \"${directory}\"," >> challenges.json
    
done

echo "  ]" >> challenges.json
echo "}" >> challenges.json

# Copy the new challenges.json file to the main folder, if it exists already it will be overwritten
cp challenges.json ../../../challenges.json

# Delete the challenges.json file
rm challenges.json

