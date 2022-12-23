import sys
# ask for hex input
hex = input("Enter a hex number: ")
# convert hex to bytes
bytes = bytes.fromhex(hex)
# convert bytes to string
string = bytes.decode("utf-8")
# print string
print(string)
    