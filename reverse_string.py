def reverse_string(text: str) -> str:
    return text[::-1]


if __name__ == "__main__":
    user_input = input("Enter a string: ")
    print(reverse_string(user_input))
