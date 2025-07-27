import os


# Get all the directory in a given directory.
def get_all_directory_in_directory(system_path: str) -> list[str]:
    return [f for f in os.listdir(path=system_path) if os.path.isdir(s=os.path.join(system_path, f))]


def main() -> None:
    # Get all the directory in a given directory
    print(get_all_directory_in_directory(
        system_path="assets/valid/aDT5V7223266kFX4uv9P9o7ovi6tdWhv"))


main()
