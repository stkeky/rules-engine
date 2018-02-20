import logging
import os

import tornado
import tornado.template
from tornado.options import define, options

path = lambda root, *a: os.path.join(root, *a)

logger = logging.getLogger(__name__)

VERSION = "0.3.0-rc.1"
GRAMMAR_FILE = "spec/rules_grammar.tx"

ROOT = os.path.dirname(os.path.abspath(__file__))
TEMPLATE_ROOT = path(ROOT, 'templates')
STATIC_ROOT = path(ROOT, 'static')

APP = {
    "port": os.environ.get("APP_PORT", 8888),
    "debug": os.environ.get("APP_DEBUG_MODE", False)
}

NATS = {
    "url": os.environ.get("NATS_URL", "nats://127.0.0.1:4222"),
    "topic": os.environ.get("NATS_TOPIC", "rules")
}

define("port", default=APP["port"], help="Run on the given port", type=int)
define("config", default=None, help="Tornado config file")
define("debug", default=APP["debug"], help="Debug mode")

options.parse_command_line()

tornado_settings = {"debug": options.debug,
                    'static_path': STATIC_ROOT,
                    'template_loader': tornado.template.Loader(TEMPLATE_ROOT)
                    }

nats_options = {
    "servers": [NATS["url"]],
    "max_reconnect_attempts": -1,
    "reconnected_cb": lambda: logger.info("Reconnected to NATS server."),
    "disconnected_cb": lambda: logger.info("Disconnected from NATS server."),
    "error_cb": lambda e: logger.error("Error in establishing connection to NATS server."),
    "close_cb": lambda: logger.warn("Connection to NATS server closed.")
}

if options.config:
    options.parse_config_file(options.config)
