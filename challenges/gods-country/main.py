from flask import Flask, request
import sqlite3
from key import KEY

con = sqlite3.connect("xors.db", check_same_thread=False)
cur = con.cursor()

def xor(a, b):
    print(a, b)
    xored = []
    for i in range(max(len(a), len(b))):
        xored_value = ord(a[i]) ^ ord(b[i%len(b)])
        xored.append(xored_value)
    return bytes(xored).hex()

def get_database_version():
    return '3.41.2'
    # return cur.execute("select sqlite_version()").fetchone()[0]

def save_key(key):
    cur.execute("create table if not exists keys (key text)")
    cur.execute("insert into keys (key) values (?)", (key,))
    con.commit()


save_key(xor(KEY, str(get_database_version())))

app = Flask(__name__)

@app.get("/")
def index():
    q = request.args.get('q')
    if (q is not None) and (q != ''):
        return cur.execute("select * from keys where key = '" + q + "'").fetchall()
    return "hey"

app.run(host='0.0.0.0', port=8006)