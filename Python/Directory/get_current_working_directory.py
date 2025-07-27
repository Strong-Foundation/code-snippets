import os


# Get the current working directory
def get_current_working_directory() -> str:
    return os.getcwd()


def main() -> None:
    # Get the current working directory
    print(get_current_working_directory())


main()
