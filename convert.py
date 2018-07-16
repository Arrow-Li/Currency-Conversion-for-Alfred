import pycurl
import StringIO
import json
import sys

api_key = "5x16quUWnKElo2s5Ytg0uHKlXdocn02f"


def convert(model, query):
    url = pycurl.Curl()
    out = StringIO.StringIO()
    url.setopt(
        pycurl.URL, "https://forex.1forge.com/1.0.3/convert?from=%s&to=CNH&quantity=%s&api_key=%s" % (model, str(query), api_key))
    url.setopt(pycurl.WRITEFUNCTION, out.write)
    url.setopt(pycurl.FOLLOWLOCATION, 1)
    url.perform()
    Convert_data = json.loads(out.getvalue())
    url.close()
    out.close()
    result = Convert_data["value"]
    result = '%.2f' % result
    return ("{\"items\": [{\"uid\": \"%s\",\"arg\":\"%s\",\"title\": \"%s\",\"icon\": {\"path\":\"%s.png\"}}]}" % (model, result, result, model))
