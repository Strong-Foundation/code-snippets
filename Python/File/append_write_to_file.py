# Append and write some content to a file.
def append_write_to_file(system_path: str, content: str):
    file = open(file=system_path, mode="a")
    file.write(content)
    file.close()


def main():
    # Append and write some content to a file.
    append_write_to_file(
        "assets/remove/zVA3m3BEJ5wSe45v93Fi8788z4GoKjnW/tEU29kAcG8m5C4y268tmJFH35", "Hello, World!")


main()
