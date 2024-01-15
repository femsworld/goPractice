# count=$(ls -1 | wc -l)
# count=$(ls . | wc -l)
count=$(find . -type f -or -type d -exec printf x \; | wc -c)
result=$((count * 5))
# printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"
# printf "\\t\\vTotal files * %d: %d\\v\\n" "$count" "$result"
printf "\\t\\v^KTotal files * %d: %d^K\\v\\n" "$count" "$result"