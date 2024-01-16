#!/bin/bash

# Look for files starting with 'a'
# find . -type f -name 'a*' 

# find . -name 'a*'

# find . -type f -name '*z'

# find . -type f -name 'z*a'

find . \( -name 'a*' -o -type f -name '*z' -o -type f -name 'z*a' \)



