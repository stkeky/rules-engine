import logging

import os

logger = logging.getLogger(__name__)


NATS = {
    "url": os.environ.get("NATS_URL", "nats://127.0.0.1:4222"),
    "rl_topic": os.environ.get("NATS_TOPIC", "rules"),
    "ev_topic": os.environ.get("NATS_TOPIC", "msg.http")
}

nats_options = {
    "servers": [NATS["url"]],
    "max_reconnect_attempts": -1,
    "reconnected_cb": lambda: logger.info("Reconnected to NATS server."),
    "disconnected_cb": lambda: logger.info("Disconnected from NATS server."),
    "error_cb": lambda e: logger.error("Error in establishing connection to NATS server."),
    "close_cb": lambda: logger.warn("Connection to NATS server closed.")
}
