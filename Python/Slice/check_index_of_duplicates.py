# Check the index of duplicates in a slice.
def check_index_of_duplicates(givenSlice):
    # Create a dictionary to store the index of each duplicates.
    duplicates_indexs = {}
    # Loop through the given slice.
    for index, content in enumerate(givenSlice):
        # If the content is already in the dictionary, append the index to the list.
        if content in duplicates_indexs:
            duplicates_indexs[content].append(index)
        # If the content is not in the dictionary, create a new entry with the index.
        else:
            duplicates_indexs[content] = [index]
    # Return the dictionary containing the index of each duplicates.
    return duplicates_indexs

# The main function to test the check_index_of_duplicates function.
def main():
    # Check the index of duplicates in the given slice.
    print(check_index_of_duplicates(["a", "b", "c", "a"]))

# Run the main function.
main()