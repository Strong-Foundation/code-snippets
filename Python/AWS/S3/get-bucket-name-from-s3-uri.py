# Get the bucket name from a s3 URI
def get_bucket_name_from_s3_uri(s3_uri: str) -> str:
    return s3_uri.split("/")[2]


def main() -> None:
    # Get the bucket name from a s3 URI
    print(get_bucket_name_from_s3_uri(s3_uri="s3://place-where-the-s3-name-goes/example.jpg"))


main()
