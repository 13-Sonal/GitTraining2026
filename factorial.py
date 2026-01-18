def factorial(n: int) -> int:
    """
    Calculates factorial of a given number.

    Args:
        n (int): Non-negative integer

    Returns:
        int: Factorial of n
    """
    result = 1
    for i in range(1, n + 1):
        result *= i
    return result

if __name__ == "__main__":
    number = int(input("Enter a non-negative integer: "))

    if number < 0:
        print("Error: Factorial is not defined for negative numbers.")
    else:
        print(factorial(number))

