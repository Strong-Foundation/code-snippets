# Import the OS module
import os
# Import the CSV module
import csv

# This function will take in a input and than write it to a CSV file
def write_to_csv(file_path, data):
    # Open the file in write mode
    with open(file_path, mode='w', newline='', encoding='utf-8') as csvfile:
        # Create a CSV writer object
        writer = csv.writer(csvfile)
        # Write the data to the CSV file
        writer.writerows(data)

# Define the main function
def main():
    # Define the file path.
    file_path = os.path.join(os.getcwd(), 'output.csv')
    
    # Data to write to the CSV file
    data = [
        ["Name", "Age", "City"],
        ["Alice", 30, "New York"],
        ["Bob", 25, "Los Angeles"],
        ["Charlie", 35, "Chicago"]
    ]

    # Call the function to write data to the CSV file
    write_to_csv(file_path, data)

# Call the main function
main()
