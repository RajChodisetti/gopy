from flask import Flask, request

import logging
from datetime import datetime

app = Flask(__name__)

@app.route('/receive_filenames', methods=['POST'])
def receive_filenames():
    filenames = request.json
    if filenames:
        for filename in filenames:
            log_relative_path(filename)
    return "Filenames received successfully!", 200

def log_relative_path(relative_path):
    current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    log_message = f"{current_time} - Relative Path: {relative_path}"
    logging.info(log_message)

if __name__ == '__main__':
    logging.basicConfig(filename='file_logs.log', level=logging.INFO)
    app.run(port=8000,debug=True)
