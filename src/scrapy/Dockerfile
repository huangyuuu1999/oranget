FROM 43.139.176.247/fruit_buckets/python:3.9-alpine

RUN echo "dockerfile under scrapy"

WORKDIR /cronjob

ADD . /cronjob/

RUN pip install -i https://mirrors.aliyun.com/pypi/simple/ --upgrade pip \
    && pip install -i https://mirrors.aliyun.com/pypi/simple/ -r /cronjob/requirements.txt

ENV PYTHONPATH "${PYTHONPATH}:/cronjob"
