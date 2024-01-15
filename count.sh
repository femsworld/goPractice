# #!/bin/bash

# count=$(find . -type f -or -type d -exec printf x \; | wc -c)
# result=$((count * 5))

# printf "\tTotal files * 5: %d\n" "$result"

count=$(find . -type f | wc -l)
result=$((count * 5))
# printf "$result"
# printf "\tTotal file * %d: %d\n" "$count" "$result"
printf "\t^KTotal files * %d: %d^K$\n" "$count" "$result"
