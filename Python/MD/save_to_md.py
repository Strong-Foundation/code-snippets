# This function saves a given string to a Markdown file with a specified name.
def save_to_md(content: str, file_path: str) -> None:
    with open(file_path, "w", encoding="utf-8") as md_file:
        md_file.write(content)

def main():
    # Example content to save
    content = "### Hello, World!\nThis is a sample Markdown file."
    
    # Specify the file path (name) where you want to save the content
    file_path = "example_file.md"
    
    # Save the content to the specified Markdown file
    save_to_md(content, file_path)

    # Print a message indicating that the content has been saved
    print(f"Content saved to {file_path}")

main()
