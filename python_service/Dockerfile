FROM python:3.8

RUN mkdir /app
WORKDIR /app

EXPOSE 5000/tcp

COPY . .

RUN pip3 install -r requirements.txt

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]