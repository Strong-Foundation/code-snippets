def averageIntInList(nums: list[int]) -> int:
    """
    Calculate the average of a list of integers and return it as an integer.

    Args:
        nums (list[int]): A list of integers.

    Returns:
        int: The average of the list, or 0 if the list is empty.
    """
    if len(nums) == 0:
        # If the list is empty, return 0
        return 0

    total: int = sum(nums)  # Sum of all numbers
    count: int = len(nums)  # Total number of elements

    return total // count  # Integer division to return an int average


# The main function.
def main():
    # Example usage of the averageIntInList function
    example_list = [10, 20, 30, 40]
    average = averageIntInList(example_list)
    print(f"The average of {example_list} is {average}")


# Example usage
if __name__ == "__main__":
    main()
