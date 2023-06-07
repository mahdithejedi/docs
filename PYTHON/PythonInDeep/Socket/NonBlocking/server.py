import socket
import sys


socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

socket.bind(('localhost', 1243))
socket.listen(5)

try:
    while True:
        conn, info = socket.accept()
        data = conn.recv(1024)
        while data:
            print(data)
            data = conn.recv(1024)
except KeyboardInterrupt:
    socket.close()

