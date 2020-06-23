import requests
import  json
from datetime import datetime

tstart = datetime.now()

resp = requests.get("https://api.ratesapi.io/api/latest")
print(json.loads(resp.content)['rates']['GBP'])
tend = datetime.now()
print (tend - tstart)