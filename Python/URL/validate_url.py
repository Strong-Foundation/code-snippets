import urllib.parse


def is_valid_url(url: str) -> bool:
    """
    Check if a given URL string is structurally valid.

    A valid URL must include:
    - A scheme (like 'http', 'https', 'ftp')
    - A network location (domain or IP)

    Parameters:
        url (str): The URL to validate.

    Returns:
        bool: True if the URL is structurally valid, False otherwise.
    """
    parsed: urllib.parse.ParseResult = urllib.parse.urlparse(url=url)
    return bool(parsed.scheme and parsed.netloc)


def main() -> None:
    print(is_valid_url(url="https://www.example.com"))  # True
    print(is_valid_url(url="ftp://ftp.server.com"))  # True
    print(is_valid_url(url="www.example.com"))  # False (missing scheme)
    print(is_valid_url(url="http:/example.com"))  # False (malformed scheme)


main()
