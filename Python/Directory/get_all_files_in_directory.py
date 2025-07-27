import os

# Get all the files in a given directory


def get_all_files_in_directory(system_path: str) -> list[str]:
    return [f for f in os.listdir(path=system_path) if os.path.isfile(path=os.path.join(system_path, f))]


def main() -> None:
    # Get all the files in a given directory
    print(get_all_files_in_directory(
        system_path="assets/valid/aDT5V7223266kFX4uv9P9o7ovi6tdWhv"))


main()
