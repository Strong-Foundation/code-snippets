# Import OS
import os
# Import CSV
import csv

# Read a CSV file and return its contents
def read_csv_file(file_path):
    # Open the CSV file in read mode
    with open(file_path, mode='r', newline='', encoding='utf-8') as csvfile:
        # Create a CSV reader object
        reader = csv.reader(csvfile)
        
        # Read and return the contents of the CSV file
        return [row for row in reader]

# Check if a file exists
def file_exists(file_path):
    # Check if the file exists
    return os.path.isfile(file_path)

# Define the main function
def main():
    # Define the file path
    file_path = os.path.join(os.getcwd(), 'output.csv')
    
    # Check if the file exists
    if file_exists(file_path) == False:
        # Print a message if the file does not exist
        print(f"File {file_path} does not exist.")
        return
    
    # Read the CSV file
    contents = read_csv_file(file_path)
    # Print the contents of the CSV file
    for row in contents:
        print(row)
   
        
# Call the main function
main()
