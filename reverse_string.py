def reverse_string(s):
    return s[::-1] #changed step to -1 for reverse iteration on string slicing

text = "Ubuntu Git Training"
print(f"Original: {text}")
print(f"Reversed: {reverse_string(text)}")
