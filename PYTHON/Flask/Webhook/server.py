from flask import Flask, request
from json import loads

app = Flask(__name__)


@app.route('/api/webhook', methods=['POST'])
def test():
    data = loads(request.get_data())
    print("data is ", data)
    return "data is %s" % data, 200, None
