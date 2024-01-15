#!/bin/bash

touch a '!' '\' '"'

if [ ! -d '\`' ]; then
    mkdir '\`'
fi

cp '!' '\`/'

# if [ "$MOVE_A" == "yes" ]; then
#     mv a '\`/'
# fi
mv a '\`/'

# Remove file 'a' if MOVE_A is set to 'no'
if [ "$MOVE_A" == "no" ]; then
    rm a
fi
