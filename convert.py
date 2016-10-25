import pycurl
import StringIO
import json
import sys

def convert(model,query):
	url=pycurl.Curl()
	out=StringIO.StringIO()
	url.setopt(pycurl.URL,'http://api.fixer.io/latest?base=CNY&symbols=%s'%model)
	url.setopt(pycurl.WRITEFUNCTION,out.write)
	url.setopt(pycurl.FOLLOWLOCATION,1)
	url.perform()
	Convert_data=json.loads(out.getvalue())
	url.close()
	out.close()
	result=query/Convert_data["rates"][model]
	result='%.2f'%result
	return ("{\"items\": [{\"uid\": \"%s\",\"arg\":\"%s\",\"title\": \"%s\",\"icon\": {\"path\":\"%s.png\"}}]}"%(model,result,result,model))