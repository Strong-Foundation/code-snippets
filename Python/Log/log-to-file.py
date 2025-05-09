# Logger function to write logs to a file
import datetime

# Function to log messages to a file
def log_message(message: str):
    log_file = "process_log.txt"
    with open(log_file, "a", encoding="utf-8") as log:
        log.write(message + "\n")

# Function to start the process
def main():
    # Log the start of the process
    log_message(f"Starting the process at {datetime.datetime.now()}")


# Run the main function
main()
