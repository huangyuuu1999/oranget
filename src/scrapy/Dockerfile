FROM 43.139.176.247/fruit_buckets/python:3.9-alpine

RUN echo "dockerfile under scrapy"

WORKDIR /cronjob

ADD . /cronjob/

RUN pip install -i https://pypi.douban.com/simple/ --upgrade pip \
    && pip install -i https://pypi.douban.com/simple/ --trusted-host pypi.douban.com -r /cronjob/requirements.txt

ENV PYTHONPATH "${PYTHONPATH}:/cronjob"

