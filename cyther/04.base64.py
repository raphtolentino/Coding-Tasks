import sys
import base64
# ask for hex input
hex = input("Enter a hex number: ")
# convert hex to bytes
bytes = bytes.fromhex(hex)
# encode bytes to base64
base64 = base64.b64encode(bytes)
# convert bytes to string
string = base64.decode("utf-8")
# print string
print(string)
    