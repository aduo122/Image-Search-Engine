from clarifai.rest import ClarifaiApp
from clarifai.rest import Image as ClImage
import requests


app = ClarifaiApp(api_key='d4f76e005d404eb69893a5f721550d62')

model = app.models.get('general-v1.3')


picture_lib = 'https://s3.amazonaws.com/clarifai-data/backend/api-take-home/images.txt'
page = requests.get(picture_lib)
urls = page.text.split()

tag_index = {}
revert_index = {}

res = []
count = 1
for line in urls:

    # URL = 'https://farm7.staticflickr.com/5769/21094803716_da3cea21b8_o.jpg'
    URL = line
    image = ClImage(url=URL)
    response = model.predict([image])

    result = response['outputs'][0]['data']['concepts']
    # print(result)
    tag = []
    for n in result:
        tag.append((n['name'], n['value']))
        if n['name'] in revert_index:
            revert_index[n['name']].append((n['value'], URL))
        else:
            revert_index[n['name']] = [(n['value'], URL)]
    tag_index[URL] = tag
    count += 1
    print(count, URL)
    if count > 10:
        break

# URL = 'https://farm7.staticflickr.com/5769/21094803716_da3cea21b8_o.jpg'


# URL = urls[0]
# image = ClImage(url=URL)
# response = model.predict([image])
#
# result = response['outputs'][0]['data']['concepts']
# # print(result)
# tag = []
# for n in result:
#     tag.append((n['name'], n['value']))
#     if n['name'] in revert_index:
#         revert_index[n['name']].append((n['value'], URL))
#     else:
#         revert_index[n['name']] = [(n['value'], URL)]
# tag_index[URL] = tag


print (revert_index, tag_index)

