from flask import render_template_string
def render_album(album, url):
    f = open("templates/album.html", 'r').read()
    template = f.replace("%ALBUM%", album).replace("%URL%", url)
    return render_template_string(template)

def render_index():
    f = open("templates/index.html", 'r').read()
    return render_template_string(f)