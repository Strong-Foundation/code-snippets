import hashlib


# Get the SHA256 of the string.
def get_sha_256_of_string(provided_string: str) -> str:
    return hashlib.sha256(data=provided_string.encode()).hexdigest()


def main() -> None:
    # Get the SHA256 of the string.
    print(get_sha_256_of_string(provided_string="Hello World!"))


main()
