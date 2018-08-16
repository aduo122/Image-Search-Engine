from clarifai.rest import ClarifaiApp
from clarifai.rest import Image as ClImage


app = ClarifaiApp(api_key='d4f76e005d404eb69893a5f721550d62')

# app = ClarifaiApp()

model = app.models.get('general-v1.3')

URL = 'https://farm7.staticflickr.com/5769/21094803716_da3cea21b8_o.jpg'
# response = model.predict_by_url(url=URL)

image = ClImage(url=URL)

response = model.predict([image])

print(response)