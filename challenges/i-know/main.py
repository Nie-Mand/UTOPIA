import redis 
import socket
from secret import KEY, VALUE
import _thread as thread

HOST = "0.0.0.0"  
PORT = 8010
kv = redis.Redis(host='KV', port=6379, decode_responses=True)

def save_key():
    kv.set(KEY, VALUE)

def filter(code):
    forbidden = ['get', 'set', 'keys', 'mget', 'echo']
    for word in forbidden:
        if word in code.lower():
            return ""
        
    return code

def getter(conn):
    data = conn.recv(1024)
    if not data or data == b'\n':
        return ""
    data = data.decode()
    data = data.replace('\n', '')
    return filter(data)

save_key()

def on_new_connection(conn):
    try:
        while True:
            ## focus here
                conn.sendall(b'$ (command) ')
                command = getter(conn)
                print(command)
                args = []
                i = 0
                while True:
                    conn.sendall(f'$ (args[{i}]) '.encode())
                    arg = getter(conn)
                    if arg == "":
                        break
                    args.append(arg)
                    i += 1
                out = kv.execute_command(command, *args)
                out = str(out)
                out += "\n"
                print(out)
                conn.sendall(out.encode())
    except Exception as e:
        conn.sendall(str(e).encode())
    finally:
        conn.close()

# just socket stuff
def accept_connections(ServerSocket):
    Client, address = ServerSocket.accept()
    print('Connected to: ' + address[0] + ':' + str(address[1]))
    thread.start_new_thread(on_new_connection, (Client, ))

def start_server(host, port):
    ServerSocket = socket.socket()
    try:
        ServerSocket.bind((host, port))
    except socket.error as e:
        print(str(e))
    print(f'Server is listing on the port {port}...')
    ServerSocket.listen()

    while True:
        accept_connections(ServerSocket)

start_server(HOST, PORT)