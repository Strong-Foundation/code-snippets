import requests


def get_url_content(target_url: str):
    """
    Send a GET request to the specified URL and return the full response content as text.

    Parameters:
        target_url (str): The URL to send the GET request to.

    Returns:
        str or None: The response content as a string if the request is successful, otherwise None.
    """
    try:
        response: requests.Response = requests.get(url=target_url)
        response.raise_for_status()  # Raise an error for bad status codes (4xx, 5xx)
        return response.text  # Return full response content as text
    except requests.exceptions.RequestException as error:
        print(f"Error fetching URL content: {error}")
        return None


def main() -> None:
    # Define the URL you want to fetch
    url_to_fetch = "https://example.com"  # Replace with any valid URL

    # Fetch content from the URL
    webpage_content = get_url_content(target_url=url_to_fetch)

    # Print the result
    if webpage_content is not None:
        print("✅ Successfully fetched content:\n")
        print(webpage_content)
    else:
        print("❌ Failed to fetch content from the URL.")


if __name__ == "__main__":
    main()
