# #!/bin/bash

# # Look for files starting with 'a'
# find . -type f -name 'a*' 

# find . -type f -name '*z'

# find . -type f -name 'z*a'


#!/bin/bash

# Find and display files starting with 'a'
echo "Files starting with 'a':"
find . -type f -name 'a*' 

# Find and display files ending with 'z'
echo -e "\nFiles ending with 'z':"
find . -type f -name '*z'

# Find and display files starting with 'z' and ending with 'a'
echo -e "\nFiles starting with 'z' and ending with 'a':"
find . -type f -name 'z*a'
