import os # Import the OS package.

# Function to walk through a directory and extract files with a specific extension
def walkGivenDirectoryAndExtractCustomFileUsingFileExtension(system_path, extension):
    matched_files = []
    for root, _, files in os.walk(system_path):
        for file in files:
            if file.endswith(extension):
                full_path = os.path.abspath(os.path.join(root, file))
                matched_files.append(full_path)
    return matched_files

def main():
    # Walk a directory
    print(walkGivenDirectoryAndExtractCustomFileUsingFileExtension("./", ".pdf"))


main()
