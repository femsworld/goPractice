# #!/bin/bash

# # Look for files starting with 'a'
# find . -type f -name 'a*' 

# find . -type f -name '*z'

# find . -type f -name 'z*a'

#!/bin/bash

# Look for files starting with 'a'
find . -type f -name 'a*' -exec echo {} \;

find . -type f -name '*z' -exec echo {} \;

find . -type f -name 'z*a' -exec echo {} \;
