import requests

r = requests.post('https://api.github.com/authorizations', headers={
    'large_body': 'a'*10000
})

print r.status_code
print r.headers
print r.content
