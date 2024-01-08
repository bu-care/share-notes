import argparse


def get_args_parser():
    parser = argparse.ArgumentParser('xbu api Config')
    parser.add_argument('-p', '--port', default=80, type=int, help='api port of xbu-api.')
    parser.add_argument('--ssl', dest="ssl", action="store_true", default=False, help='use https instead of http')
    parser.add_argument('--debug', default=False)

    args = parser.parse_args()

    return args


if __name__ == '__main__':
    args = get_args_parser()

    from app import app, start_app

    if args.debug:
        app.config["DEBUG"] = True

    start_app(args.ssl, args.port)

