from flask import Flask
from albums import albums
from utils import render_album, render_index

app = Flask(__name__)

@app.get('/')
def index():
    return render_index()

@app.get('/<album_name>')
def get_album(album_name):
    album = {
        "name": album_name,
        "url": "https://notfound.com"
    }
    for a in albums:
        if a['name'] == album_name:
            album = a
            break

    return render_album(album["name"], album["url"])

app.run(host='0.0.0.0', port=8008)