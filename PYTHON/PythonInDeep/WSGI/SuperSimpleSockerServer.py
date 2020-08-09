import socket
from SimpleHttpServer import view, http_parser, process_response

with socket.socket() as s:
    s.bind(('localhost', 8080))
    s.listen(1)
    conn, addr = s.accept()
    while True:
        with conn:
            request = conn.recv(1024).decode('utf-8')
            request = http_parser(request)
            response = view(request)
            http_response = process_response(response)
            conn.sendall(http_response.encode('utf-8'))