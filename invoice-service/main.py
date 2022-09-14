from flask import Flask, jsonify, request

app = Flask(__name__)

invoices = [
  { 'order_id': '1', 'invoice_id': '1', 'amount': 1000.00 },
  { 'order_id': '2', 'invoice_id': '2', 'amount': 1500.00 },
  { 'order_id': '3', 'invoice_id': '3', 'amount': 2000.00 }
]


@app.route('/invoices')
def get_invoices():
  return jsonify(invoices)


@app.route('/orders/<order_id>/invoice')
def get_invoice(order_id):
    for invoice in invoices:
        if invoice.get('order_id') == order_id:
            return invoice
    return '', 404


@app.route('/orders/<order_id>/invoice', methods=['POST'])
def add_invoice():
  invoices.append(request.get_json())
  return '', 204


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8082)