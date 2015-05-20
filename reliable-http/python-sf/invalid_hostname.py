import requests

import time
start = time.time()
try:
    requests.get('http://foobarbizbangesateeueutseh.com', timeout=3)
    raise ValueError()
except requests.exceptions.ConnectionError:
    pass
end = time.time()
print end - start
