#!/bin/bash

count=$(find . -type f -or -type d -exec printf x \; | wc -c)
result=$((count * 5))

printf "\tTotal files * 5: %d\n" "$result"
