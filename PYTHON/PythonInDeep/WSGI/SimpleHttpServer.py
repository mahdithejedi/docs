from io import StringIO

def http_parser(http):
    request, *headers, _, body = http.split('\r\n')
    method, path, protocol = request.split(' ')
    headers = dict(
        line.split(':', maxsplit=1) 
        for line in headers
    )
    return method, path, protocol, headers, body

def process_response(response):
    return str((
        'HTTP/1.1 200 OK\r\n',
        f'Content-Lenght: {len(response)}\r\n', 
        'Content-Type: text/html\r\n',
        '\r\n',
        response,
        '\r\n'
    ))

def view(request):
    return """Every One this is a simple response 
    method:{},
    path:{},
    protocol:{},
    headers:{},
    body:{}
    """.format(
        *request
    )

def to_environ(method, path, protocol, headers, body):
    return {
        'REQUEST_METHOD': method,
        'PATH_INFO': path,
        'SERVER_PROTOCOL': protocol,
        'wsgi.input': StringIO(body),
        # **format_headers(headers)
    }