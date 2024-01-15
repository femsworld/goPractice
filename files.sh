#!/bin/bash

# Create files a, !, \, and "
touch a '!' '\' '"'

# Create a directory called `
mkdir '\`'

# Copy the file '!' into the ` directory
cp '!' '\`/'

# Check the environment variable
if [ "$MY_VARIABLE" == "yes" ]; then
    mv a '\`/'
elif [ "$MY_VARIABLE" == "no" ]; then
    # If MY_VARIABLE is set to 'no', remove file 'a'
    rm a
else
    # Otherwise, leave the file 'a' as is
    :
fi
