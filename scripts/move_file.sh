#! /bin/bash

# Check if correct number of parameters are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <pattern> <destination>"
    echo "Example: $0 '*.txt' /path/to/destination"
    exit 1
fi

PATTERN=$1
DESTINATION=$2

# Check if destination exists
if [ ! -d "$DESTINATION" ]; then
    echo "Error: Destination directory '$DESTINATION' does not exist."
    exit 1
fi

# move files
for file in $(ls $PATTERN 2>/dev/null); do
    mv "$file" "$DESTINATION"
    echo "Moved: $file -> $DESTINATION"
done
