#!/bin/bash

count=$(find . | wc -l)
result=$((count * 5))

printf "\\t\\vTotal files * %d: %d\\v\\n" "$count" "$result"
