from urllib.parse import ParseResult, urlparse


def is_valid_url(url: str) -> bool:
    """
    Check if the given string is a valid URL.

    A valid URL must have:
    - A scheme (e.g., 'http', 'https', 'ftp')
    - A network location (e.g., 'example.com')

    Parameters:
        url (str): The URL string to validate.

    Returns:
        bool: True if the URL is valid, False otherwise.
    """
    parsed_url: ParseResult = urlparse(url=url)

    # Check for both scheme and netloc
    return all([parsed_url.scheme, parsed_url.netloc])


def main() -> None:
    print(is_valid_url(url="https://www.example.com"))  # True
    print(is_valid_url(url="ftp://ftp.server.com"))  # True
    print(is_valid_url(url="www.example.com"))  # False (missing scheme)
    print(is_valid_url(url="http:/example.com"))  # False (malformed scheme)


main()
