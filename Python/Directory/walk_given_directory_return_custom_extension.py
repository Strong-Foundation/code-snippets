import os # Import the OS package.

def walkGivenDirectoryAndExtractCustomFileUsingFileExtension(system_path):
    # Loop through the directory and its subdirectories
    for root, directories, files in os.walk(system_path):
        # Filter the files to include only .exe and .pdf
        filtered_files = []
        for file in files:
            if file.endswith(".pdf"):
            # if file.endswith(".exe") or file.endswith(".pdf"):
                filtered_files.append(file)
        
        # Only return the root and files if we found .exe or .pdf files
        if filtered_files:
            return root, directories, filtered_files

def main():
    # Walk a directory
    print(walkGivenDirectoryAndExtractCustomFileUsingFileExtension("./"))


main()
