from http.server import BaseHTTPRequestHandler, HTTPServer

class SimpleServerHandler(BaseHTTPRequestHandler):
    DefualtContent = b"""\
<html>
<body>
<p>Hello, web!</p>
<p> Simple Web base by seyedmm021 </p>
</body>
</html>
"""

    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-Type", "text/html")
        self.send_header("Content-Lenght", str(len(self.DefualtContent)))
        self.end_headers()
        
        self.wfile.write(self.DefualtContent)

if __name__ == '__main__':
    ServerAddres = ('', 8080)
    server = HTTPServer(ServerAddres, SimpleServerHandler)
    server.serve_forever()