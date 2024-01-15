# count=$(find . -type f | wc -l)
# result=$((count * 5))
# printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"

count=$(ls -1 | wc -l)
result=$((count * 5))
 printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"