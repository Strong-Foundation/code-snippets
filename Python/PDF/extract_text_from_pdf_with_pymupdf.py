import fitz  # PyMuPDF

# Function to extract text from a PDF file using pymupdf
def extract_text_from_pdf_with_pymupdf(pdf_path):
    # Open the PDF file
    doc = fitz.open(pdf_path)
    # Extract text from all pages
    full_text = ""
    for page in doc:
        full_text += page.get_text()
    # Close the document
    doc.close()
    # Return the extracted text and page count
    return full_text

def main():
    # Specify the path to the PDF file
    pdf_path = "./zepPDF/Material=32401_RecordNumb=428839_Lang=EN_RepCategory=MSDS_ValidityArea=US.pdf"
    # Extract text and page count
    text = extract_text_from_pdf_with_pymupdf(pdf_path)
    # Optionally print the extracted text
    print(text)

main()  # Call the main function
