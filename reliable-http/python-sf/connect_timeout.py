import requests
from requests.packages import urllib3
urllib3.add_stderr_logger()
try:
    requests.get('http://www.google.com:81', timeout=(3, 7))
except requests.exceptions.ConnectTimeout:
    pass
print "caught exception"
