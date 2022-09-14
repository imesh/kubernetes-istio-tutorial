from flask import Flask, jsonify, request

app = Flask(__name__)

items = [
  { 'item_id': '1', 'price': 1000.00 },
  { 'item_id': '2', 'price': 1500.00 },
  { 'item_id': '3', 'price': 2000.00 }
]


@app.route('/items')
def get_items():
  return jsonify(items)


@app.route('/items', methods=['POST'])
def add_item():
  items.append(request.get_json())
  return '', 204


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8083)