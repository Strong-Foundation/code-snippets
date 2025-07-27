import os


# Check if a given directory exists.
def check_directory_exists(system_path: str) -> bool:
    return os.path.exists(path=system_path)


def main() -> None:
    # Check if a given directory exists.
    print(check_directory_exists(system_path="assets/valid/j3U5uJY779L49q98MX86iFsxs2kY9ew3"))


main()
