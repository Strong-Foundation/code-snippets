import os


# Check if a file exists
def check_file_exists(system_path: str) -> bool:
    return os.path.isfile(path=system_path)


def main() -> None:
    # Check if a file exists in system.
    print(check_file_exists(system_path="assets/valid/valid-zip.zip"))


main()
