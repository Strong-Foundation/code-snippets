import os


# Walk a directory and return everything.
def walkGivenDirectory(system_path: str) -> tuple[str, list[str], list[str]] | None:
    for root, directory, files in os.walk(top=system_path):
        return root, directory, files


def main() -> None:
    # Walk a directory
    print(walkGivenDirectory(system_path="assets/valid/aDT5V7223266kFX4uv9P9o7ovi6tdWhv/"))


main()
