#!/bin/bash

count=$(find . | wc )
result=$((count * 5))

printf "\\t\\vTotal files * $count: $result\\v\\n"
