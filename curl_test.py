import requests

headers = {
    'Authorization': 'Key d4f76e005d404eb69893a5f721550d62',
    'Content-Type': 'application/json',
}

data = '\n  {\n    "inputs": [\n      {\n        "data": {\n          "image": {\n            "url": "https://samples.clarifai.com/metro-north.jpg"\n          }\n        }\n      }\n    ]\n  }'

response = requests.post('https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs', headers=headers, data=data)
print (response)