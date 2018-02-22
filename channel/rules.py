import json

import tornado.ioloop
from nats.io import Client as NATSClient

from settings import *

rules = \
    {
        "rules": [
            {
                "conditions": [
                    {
                        "operator": "BETWEEN",
                        "property": "temperature",
                        "deviceId": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
                        "value": {
                            "to": 20,
                            "from": 10
                        }
                    }
                ],
                "userId": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
                "name": "ruleTemp01",
                "actions": [
                    {
                        "content": "High temperature in living room!",
                        "recipient": "blagojevicb94@gmail.com",
                        "name": "SEND EMAIL"
                    }
                ]
            },
            {
                "conditions": [
                    {
                        "operator": "=",
                        "property": "name",
                        "deviceId": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
                        "value": "t-3"
                    }
                ],
                "userId": "08b0ccf2-09ca-11e8-81a9-0242ac110005",
                "name": "ruleTemp02",
                "actions": [
                    {
                        "content": "Event with name t-3 appeared",
                        "recipient": "blagojevicb94@gmail.com",
                        "name": "SEND EMAIL"
                    }
                ]
            }
        ]
    }


@tornado.gen.coroutine
def main():
    try:
        nats = NATSClient()
        yield nats.connect(**nats_options)

        yield nats.publish(NATS["rl_topic"], json.dumps(rules))
        yield nats.flush()

        print("Published to '{0}'".format(rules))
    except Exception as ex:
        print(ex)


if __name__ == "__main__":
    tornado.ioloop.IOLoop.instance().run_sync(main)
