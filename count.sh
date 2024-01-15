# count=$(ls -1 | wc -l)
count=$(ls . | wc -l)
result=$((count * 5))
# printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"
# printf "\\t\\vTotal files * %d: %d\\v\\n" "$count" "$result"
printf "\\t\\v^KTotal files * %d: %d^K\\v\\n" "$count" "$result"