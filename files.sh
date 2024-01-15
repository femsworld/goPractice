# #!/bin/bash

# # Create files a, !, \, and "
# touch a '!' '\' '"'

# # Create a directory called ` if it doesn't exist
# if [ ! -d '\`' ]; then
#     mkdir '\`'
# fi

# # Copy the file '!' into the ` directory
# cp '!' '\`/'

# # Check the environment variable
# if [ "$MOVE_A" == "yes" ]; then
#     # If MOVE_A is set to 'yes', move file 'a' into the ` directory
#     mv a '\`/'
# elif [ "$MOVE_A" == "no" ]; then
#     # If MOVE_A is set to 'no', remove file 'a'
#     rm a
# fi

#!/bin/bash

# Create files a, !, \, and "
touch a '!' '\' '"'

# Create a directory called ` if it doesn't exist
if [ ! -d '\`' ]; then
    mkdir '\`'
fi

# Copy the file '!' into the ` directory
cp '!' '\`/'

# Check the environment variable
if [ "$MOVE_A" == "yes" ]; then
    # If MOVE_A is set to 'yes', move file 'a' into the ` directory
    mv a '\`/'
fi

# Print the current state
echo "Current state:"
ls -R

# Remove file 'a' if MOVE_A is set to 'no'
if [ "$MOVE_A" == "no" ]; then
    rm a
fi

# Print the final state
echo "Final state:"
ls -R
