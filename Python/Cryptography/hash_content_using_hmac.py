import hmac
import hashlib
import codecs


# Hash a given content using hmac.
def hash_content_using_hmac(given_content: str, given_password: str) -> str:
  return hmac.new(key=codecs.encode(obj=given_password), msg=codecs.encode(obj=given_content), digestmod=hashlib.sha256).hexdigest()


def main() -> None:
    # Hash a given content using hmac
    print(hash_content_using_hmac(given_content="Hello, World!", given_password="password"))


main()
