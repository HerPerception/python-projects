import socket, sys

user_input = input("Input website to retrieve IP addresses: ")

try:
        addr = socket.gethostbyname(user_input)
except socket.gaierror:
        print("Error finding IP address.")
        sys.exit()
print(f"IPV4 address: {addr}")
