import os


# Change the current working directory.
def change_current_working_directory(system_path: str) -> None:
    os.chdir(path=system_path)


def main() -> None:
    # Change the current working directory.
    change_current_working_directory(
        system_path="assets/remove/zVA3m3BEJ5wSe45v93Fi8788z4GoKjnW")


main()
