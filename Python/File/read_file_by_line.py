# Read all lines from a file, filter out empty lines, and normalize each line
def read_file_by_line(file_name: str) -> list[str]:
    """
    Read all non-empty lines from a text file, remove surrounding whitespace,
    convert to lowercase, and return them in a list.

    Parameters:
        file_name (str): Path to the input text file.

    Returns:
        list[str]: A list containing cleaned and normalized non-empty lines.
    """
    # Initialize an empty list to store valid lines
    non_empty_lines: list[str] = []

    # Open the file safely using a context manager
    with open(file=file_name, mode="r", encoding="utf-8") as file:
        # Iterate over each line in the file
        for raw_line in file:
            # Strip leading/trailing whitespace and convert to lowercase
            stripped_line: str = raw_line.strip().lower()
            # Only add lines that are not empty after stripping
            if stripped_line:
                non_empty_lines.append(stripped_line)

    # Return the list of cleaned, non-empty lines
    return non_empty_lines

def main():
    # Read the file line by line
    lines = read_file_by_line("assets/valid/valid-json.json")
    for line in lines:
        print(line)


main()
