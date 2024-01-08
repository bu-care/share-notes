import logging
from datetime import datetime
from flask import Flask
# from flask_session import Session
from xbu_api import xbu_api


app = Flask(__name__)

app.register_blueprint(xbu_api, url_prefix="/")

# # app.config['APPLICATION_ROOT'] = ''
# app.config['SECRET_KEY']='s()nIcWa11_rsapi'
# app.config['SESSION_PERMANENT'] = False
# app.config['PERMANENT_SESSION_LIFETIME'] = 120
# app.config['SESSION_USE_SIGNER']=True
# Session(app)
#
# gunicorn_logger = logging.getLogger('gunicorn.error')
# app.logger.handlers = gunicorn_logger.handlers
# app.logger.setLevel(gunicorn_logger.level)


class Logger(object):
    def __init__(self, log_level, logger_name):
        self.__logger = logging.getLogger(logger_name)
        self.__logger.setLevel(log_level)
        # 创建一个FileHandler，并对输出消息的格式进行设置，将其添加到logger，然后将日志写入到指定的文件中
        cur_date = datetime.now().strftime('%Y-%m-%d')
        log_file_name = f"./log/{cur_date}log.txt"
        file_handler = logging.FileHandler(log_file_name)

        # logger中添加StreamHandler，可以将日志输出到屏幕上
        console_handler = logging.StreamHandler()

        # formatter = logging.Formatter(
        #     '[%(asctime)s] - [logger name :%(name)s] - [%(filename)s file line:%(lineno)d] - %(levelname)s: %(message)s')

        formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')

        # formatter = logging.Formatter('%(message)s')

        file_handler.setFormatter(formatter)
        console_handler.setFormatter(formatter)

        self.__logger.addHandler(file_handler)
        self.__logger.addHandler(console_handler)

    def get_log(self):
        return self.__logger


def start_app(is_ssl, api_port):
    from gevent import pywsgi
    from geventwebsocket.handler import WebSocketHandler
    # app.debug = True

    server_host = "0.0.0.0"

    if is_ssl:
        ssl_key_file = 'server.key'
        ssl_cert_file = 'server.crt'
        api_server = pywsgi.WSGIServer((server_host, api_port), app, handler_class=WebSocketHandler,
                                       keyfile=ssl_key_file, certfile=ssl_cert_file)
    else:
        api_server = pywsgi.WSGIServer((server_host, api_port), app, handler_class=WebSocketHandler)

    api_server.prevent_wsgi_call = 1
    api_server.logger = Logger("INFO", "app_log").get_log()
    print("\n\n xbu_api simulator server running!!!")
    api_server.serve_forever()


