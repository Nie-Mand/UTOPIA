from secret import KEY
from utils import get_g, get_p, get_public_key, get_secret_key
import _thread as thread
import socket

HOST = "0.0.0.0"  
PORT = 8015

def on_new_connection(conn):
    try:
        # DIFFIE HELLMANN
        p = get_p()
        conn.sendall(f"Here's P -> {p}\n".encode())

        g = get_g()
        conn.sendall(f"Here's G -> {g}\n".encode())

        public_key = get_public_key(g)
        conn.sendall(f"Here's my public key -> {public_key}\n".encode())

        conn.sendall(b'Enter your key: ')
        data = conn.recv(1024)
        if not data:
            return
        data = data.decode()
        users_key = int(data)

        if users_key > p:
            raise("too much")
            
        my_secret_key = get_secret_key(users_key, public_key)

        conn.sendall(b'Fist Bump ? -> ')
        data = conn.recv(1024)
        if not data:
            return
        data = data.decode()
        your_secret_key = int(data)

        if your_secret_key == my_secret_key:
            conn.sendall(f"Cya in UTOPIA, -> {KEY}\n".encode())
        else:
            conn.sendall("You wasted my time\n".encode())
    except:
        pass
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