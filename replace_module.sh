#!/bin/bash

# Check if the number of arguments is correct
if [ $# -ne 3 ]; then
    echo "Usage: $0 <old_module_name> <new_module_name> <project_directory>"
    exit 1
fi

old_module_name="$1"
new_module_name="$2"
project_directory="$3"

# Check if the project directory exists
if [ ! -d "$project_directory" ]; then
    echo "Project directory '$project_directory' does not exist."
    exit 1
fi

# Replace the module name in .go files recursively
grep -rlF "$old_module_name" "$project_directory" | while IFS= read -r file; do
    sed -i.bak "s@$old_module_name@$new_module_name@g" "$file"
done

# Remove backup files created by sed
find "$project_directory" -type f -name "*.bak" -delete

# Inform the user about the replacement
echo "Module name '$old_module_name' replaced with '$new_module_name' in '.go' files within '$project_directory'."