# count=$(find . -type f | wc -l)
# result=$((count * 5))
# printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"

count=$(ls -1 | wc -l)
result=$((count * 5))
#  printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"
# printf "\t^KTotal files * %d: %d^K\n" "$count" "$result" | sed -e 's/\t/^I/g'
printf "\t^KTotal files * %d: %d^K\n" "$count" "$result"