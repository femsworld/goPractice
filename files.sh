#!/bin/bash

touch a '!' '\' '"'

if [ ! -d '\`' ]; then
    mkdir '\`'
fi

cp '!' '`'

if [ "$MOVE_A" == "yes" ]; then
    # Move the file a into the  directory
    mv a '`'
elif [ "$MOVE_A" == "no" ]; then
    # Remove the file a
    rm a
else
    # Leave the file a as is
    :
fi

