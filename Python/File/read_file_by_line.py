# Read a file line by line and return a list of lines
def read_file_by_line(file_name: str) -> list[str]:
    file_object = open(file=file_name, mode="r")
    lines: list[str] = file_object.readlines()
    file_object.close()
    return lines


def main():
    # Read the file line by line
    lines = read_file_by_line("assets/valid/valid-json.json")
    for line in lines:
        print(line)


main()
