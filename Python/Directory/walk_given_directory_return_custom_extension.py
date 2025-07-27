import os # Import the OS package.

# Function to walk through a directory and extract files with a specific extension
def walkGivenDirectoryAndExtractCustomFileUsingFileExtension(system_path: str, extension: str) -> list[str]:
    matched_files: list[str] = []
    for root, _, files in os.walk(top=system_path):
        for file in files:
            if file.endswith(extension):
                full_path: str = os.path.abspath(os.path.join(root, file))
                matched_files.append(full_path)
    return matched_files

def main() -> None:
    # Walk a directory
    print(walkGivenDirectoryAndExtractCustomFileUsingFileExtension(system_path="./", extension=".pdf"))


main()
