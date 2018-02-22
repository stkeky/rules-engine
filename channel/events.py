import json

import tornado.ioloop
from nats.io import Client as NATSClient

from settings import *

events = [
    {
        "channel": "148fba58-09ca-11e8-81ab-0242ac110005",
        "publisher": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
        "protocol": "http",
        "content_type": "application/senml+json",
        "payload": "W3siYm4iOiJzb21lLWJhc2UtbmFtZToiLCJidCI6MS4yNzYwMjAwNzYwMDFlKzA5LCAiYnUiOiJBIiwiYnZlciI6NSwgIm4iOiJ2b2x0YWdlIiwidSI6IlYiLCJ2IjoxMjAuMX0sIHsibiI6ImN1cnJlbnQiLCJ0IjotNSwidiI6MS4yfSwgeyJuIjoiY3VycmVudCIsInQiOi00LCJ2IjoxLjN9XQ=="
    },
    {
        "channel": "148fba58-09ca-11e8-81ab-0242ac110005",
        "publisher": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
        "protocol": "http",
        "content_type": "application/senml+json",
        "payload": "WyB7Im4iOiJ0ZW1wZXJhdHVyZSIsInQiOi01LCJ2IjoxMC4yfSwgeyJuIjoiY3VycmVudCIsInQiOi00LCJ2IjoxLjN9XQo="
    },
    {
        "channel": "148fba58-09ca-11e8-81ab-0242ac110005",
        "publisher": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
        "protocol": "http",
        "content_type": "application/senml+json",
        "payload": "WyB7Im4iOiJuYW1lIiwidCI6LTUsInZzIjogInQtMyJ9XQ=="
    }
]


@tornado.gen.coroutine
def main():
    try:
        nats = NATSClient()
        yield nats.connect(**nats_options)

        for ev in events:
            yield nats.publish(NATS["ev_topic"], json.dumps(ev))
        yield nats.flush()

        print("Published to '{0}'".format(events))
    except Exception as ex:
        print(ex)


if __name__ == "__main__":
    tornado.ioloop.IOLoop.instance().run_sync(main)
