from multiprocessing import Process
import socket
import signal
import os


def _socket(port):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    sock.bind(('localhost', port))
    print('socket is on for ', port)
    sock.listen(5)

    try:
        while True:
            conn, info = sock.accept()
            data = conn.recv(1024)
            while data:
                print(data)
                data = conn.recv(1024)
    except KeyboardInterrupt:
        print('going to close socket', port)
        sock.close()


ports = [1243, 1244, 1245, 1246, 11247]

_process_list = []
for port in ports:
    _process = Process(target=_socket, args=(port,))
    _process.start()
    _process_list.append(
        _process
    )


def signal_handler(_, __):
    [os.kill(p.pid, signal.SIGTERM) for p in _process_list]


signal.signal(signal.SIGINT, signal_handler)
