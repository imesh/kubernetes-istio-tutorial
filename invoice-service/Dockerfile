FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY main.py ./

CMD ["gunicorn"  , "-b", "0.0.0.0:8082", "main:app"]
