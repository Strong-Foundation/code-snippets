import hashlib


# Get the sha-3-512 of the string.
def get_sha_3_512_of_string(provided_string: str) -> str:
    return hashlib.sha3_512(data=provided_string.encode()).hexdigest()


def main() -> None:
    # Get the sha-3-512 of the string.
    print(get_sha_3_512_of_string(provided_string="Hello World!"))


main()
