import shutil

# Remove a directory.


def remove_directory(system_path: str) -> None:
    shutil.rmtree(system_path)


def main() -> None:
    # Remove a directory
    remove_directory(system_path="assets/remove/89mmtNQY7hM7389f48Sw46ZhbRDNQ2h9")


main()
