python脚本起一个简单的http的server，默认路径是在当前目录下，以前可以通过偏函数partial来修改SimpleHTTPRequestHandler的默认directory实现效果，但是最近不知道为什么不能用了。

```python
def local_start_server(self, file_dir='./', server_class=HTTPServer):
        handler_class = partial(SimpleHTTPRequestHandler, directory=file_dir)
        server_address = (self.cur_ip, self.transfer_server_port)
        httpd = server_class(server_address, handler_class)
        httpd.serve_forever()
```

通过另一种方法来实现，这是参考代码

```python
from http.server import HTTPServer
from http.server import SimpleHTTPRequestHandler
import os


class MyHTTPHandler(SimpleHTTPRequestHandler):
    """This handler uses server.base_path instead of always using os.getcwd()"""

    def translate_path(self, path):
        path = SimpleHTTPRequestHandler.translate_path(self, path)
        # # getcwd() returns current working directory of a process
        # # method in Python is used to get a relative filepath to the given path
        # # either from the current working directory or from the given directory.
        relpath = os.path.relpath(path, os.getcwd())
        fullpath = os.path.join(self.server.base_path, relpath)
        # print(f'init_path: {path}, \ngetcwd:{os.getcwd()}, \nrelpath: {relpath}, \nfullpath: {fullpath}')
        # # In fact, as long as the self.server.base_path is returned here, the specified directory can be achieved
        return fullpath

    def end_headers(self):
        self.send_header('Access-Control-Allow-Origin', '*')
        SimpleHTTPRequestHandler.end_headers(self)


class MyHTTPServer(HTTPServer):
    """The main server, your pass in base_path which is the path you want to serve requests from"""

    def __init__(self, base_path, server_address, request_handler_class=MyHTTPHandler):
        self.base_path = base_path
        HTTPServer.__init__(self, server_address, request_handler_class)


def start_server(file_dir, port=80):
    httpd = MyHTTPServer(file_dir, ("0.0.0.0", port))
    print('start http server!!!')
    httpd.serve_forever()


if __name__ == '__main__':
    PORT = 80
    web_dir = "../reports"
    start_server(web_dir, PORT)

```

实际上可以进行简化，只需要SimpleHTTPRequestHandler的translate_path函数能返回我们指定的目录即可，简化后的代码如下：

```python
from http.server import HTTPServer, SimpleHTTPRequestHandler
import os
import ssl
from functools import partial
from logger import logger


def start_server(port=443):
    server_address = ('0.0.0.0', port)
    cert_dir = os.environ['ROOT_DIR'] + '/utils/cert-data'

    httpd = HTTPServer(server_address, SimpleHTTPRequestHandler)
    httpd.socket = ssl.wrap_socket(httpd.socket,
                                   server_side=True,
                                   certfile=f'{cert_dir}/cert.pem',
                                   keyfile=f"{cert_dir}/key.pem",
                                   ssl_version=ssl.PROTOCOL_TLS_SERVER)  # PROTOCOL_TLSv1
    httpd.logger = logger
    print('start https server!!!')
    httpd.serve_forever()


class MyHTTPHandler(SimpleHTTPRequestHandler):
    """This handler uses server.base_path instead of always using os.getcwd()"""

    def translate_path(self, path):
        # path = SimpleHTTPRequestHandler.translate_path(self, path)
        # 返回 HTTPServer.base_path 就可以实现效果
        return self.server.base_path


def start_http_server(file_dir, port=80):
    server_address = ('0.0.0.0', port)

    # 将 HTTPServer.base_path 修改为指定的路径
    HTTPServer.base_path = file_dir
    httpd = HTTPServer(server_address, MyHTTPHandler)

    print('start http server!!!')
    httpd.serve_forever()


if __name__ == '__main__':
    file_dir = '../firmware'
    start_http_server(file_dir, 80)
```

