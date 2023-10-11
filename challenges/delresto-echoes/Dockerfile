FROM ubuntu:latest AS First

WORKDIR /echoes

RUN apt-get update && apt-get install -y uuid-runtime

COPY run.sh .

RUN ./run.sh

FROM python:alpine3.18

WORKDIR /app

COPY --from=First /echoes .

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8009

CMD ["python","main.py"]
