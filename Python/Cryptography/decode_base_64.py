import base64


# Decode the base-64 string.
def decode_base_64(provided_string: str) -> str:
    return base64.b64decode(provided_string).decode()


def main() -> None:
    # Decode the base-64 string.
    print(decode_base_64(provided_string="SGVsbG8gV29ybGQh"))


main()
