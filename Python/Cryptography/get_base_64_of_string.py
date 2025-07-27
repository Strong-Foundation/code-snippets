import base64


# Get the base-64 of the string.
def get_base_64_of_string(provided_string: str) -> str:
    return base64.b64encode(s=provided_string.encode()).decode()


def main() -> None:
    # Get the base-64 of the string.
    print(get_base_64_of_string(provided_string="Hello World!"))


main()
