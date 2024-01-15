#!/bin/bash

# count=$(find . -type f -or -type d -exec printf x \; | wc -c)
# result=$((count * 5))

# printf "\tTotal files * 5: %d\n" "$result"

count=$(find . -type f | wc -l)
result=$((count * 5))
# printf "$result"
# printf "\tTotal file * %d: %d\n" "$count" "$result"
printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"
# echo -e "\t\x0B^KTotal files * $count: $result^K$\n"
# echo -e "\t\x0B^KTotal files * $count: $result^K$\c"
# printf "\t\x0B^KTotal files * %d: %d^K$\n" "$count" "$result"
# printf "\t^KTotal files * %d: %d^K$\n" "$count" "$result" | sed 's/\^K/\x0B/g'
# printf "\t\x0B\033[KTotal files * %d: %d\x0B$\n" "$count" "$result"