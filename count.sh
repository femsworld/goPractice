#!/bin/bash

count=$(find . | wc -l)
result=$((count * 5))

printf "\\t\\vTotal files * $count: $result\\v\\n"
