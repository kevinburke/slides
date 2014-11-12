import requests

r = requests.post('https://api.sendgrid.com/api/mail.send.json', data={
    'large_body': 'a'*100000
})

print r.status_code
print r.headers
print r.content
