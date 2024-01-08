import requests
from utils.database_operate import ExecMongoDb
from utils.update_zta_server_code import UpdateZtaInDocker, UpdateZtaIn238

from flask import request, Blueprint
import logging
import urllib3

urllib3.disable_warnings()

# app = Flask(__name__)
logger = logging.getLogger('app_log.xbu_api')
xbu_api = Blueprint('xbu_api', __name__)

cc_mongo_data = {}


@xbu_api.route('/--/<email>', methods=['GET'])
def get_account_id(email):
    mongo = ExecMongoDb(cc_mongo_data)

    for i in mongo.db['users'].find():
        if i['email'] == email:
            if len(i['accountRoles']) == 0:
                mongo.close_db()
                return f"The 'accountroles' value of the email login user({email}) is empty " \
                       f"and there is no 'account'.", 404
            else:
                for j in mongo.db['accounts'].find():
                    if j['_id'] == i['accountRoles'][0]['account']:
                        account_id = j['accountId']
                        mongo.close_db()
                        return account_id
    mongo.close_db()
    return f"The '{email}' email login user was not queried!", 404


docker_zta = UpdateZtaInDocker()


@xbu_api.route('/operate_ztna_server', methods=['POST'])
def operate_ztna_server():
    data = request.get_json(force=True)
    restart_way = data.get('restart_way')
    action = data.get('action')
    zta_port = '9000'
    logger.info(f"request_body: {data}")
    if None in [restart_way, action]:
        return 'The request body is missing a required parameter.', 404
    if restart_way == "docker":
        ztna = docker_zta
        update_zta = UpdateZtaInDocker
    else:
        zta_port = data.get('zta_port')
        ztna = ser238_zta
        update_zta = UpdateZtaIn238

    def start_zta(zta, port):
        zta.set_cli_list(port)
        zta.ssh_connect()
        zta.start_server()
        stdout_info = f'start ZTNA port is 9000 in {restart_way}\n\n' + ''.join(update_zta.stdout_info)
        return stdout_info

    def close_client(zta):
        logger.info('\nrequest zta_close_client')
        zta.close_client()
        return 'Close paramiko client which connect to ZTNA server'

    if action == 'start':
        return start_zta(ztna, zta_port)
    elif action == 'close':
        return close_client(ztna)


@xbu_api.route('/docker_operate_zta/<symbol>', methods=['GET'])
def docker_operate_zta(symbol):
    def start_zta(ztna):
        ztna.set_cli_list()
        ztna.ssh_connect()
        ztna.start_server()
        stdout_info = f'start zta port is 9000 in docker\n\n' + ''.join(UpdateZtaInDocker.stdout_info)
        return stdout_info

    def close_client(ztna):
        logger.info('\nrequest zta_close_client')
        ztna.close_client()
        return 'Close paramiko client which connect to ZTA server'

    if symbol == 'start':
        return start_zta(docker_zta)
    elif symbol == 'close':
        return close_client(docker_zta)


ser238_zta = UpdateZtaIn238()


@xbu_api.route('/ser238_operate_zta/<symbol>-<port>', methods=['GET'])
def ser238_operate_zta(symbol, port):
    def start_zta(zta, zta_port):
        zta.set_cli_list(zta_port)
        zta.ssh_connect()
        zta.start_server()
        stdout_info = f'start zta port is {zta_port} in 238 server\n\n' + ''.join(UpdateZtaIn238.stdout_info)
        return stdout_info

    def close_client(zta):
        logger.info('\nrequest zta_close_client')
        zta.close_client()
        return 'Close paramiko client which connect to ZTA server'

    if symbol == 'start':
        stdout = start_zta(ser238_zta, zta_port=port)
        return stdout
    elif symbol == 'close':
        return close_client(ser238_zta)


def get_zta_account(email):
    user = {}
    if email in user:
        return user[email]
    else:
        return 404


@xbu_api.route('/zta_login', methods=['POST'])
def zta_login_page():
    data = request.get_json(force=True)
    email = data.get('email')
    endpoint = data.get('login_endpoint')
    logger.info(f"email: {email}, login_endpoint: {endpoint}")
    if None in [email, endpoint]:
        return 'The request body is missing a required parameter.', 404

    password = get_zta_account(email)
    if password == 404:
        return f"The user of '{email}' email was not queried!", 404

    request_body = {"password": f"{password}", "email": email}
    response_content, status_code = login_or_logout(endpoint, "login", request_body)
    return response_content, status_code


def login_or_logout(endpoint, target, request_body, headers=None):
    url = f'{endpoint}/api/{target}'
    response = requests.post(url=url, data=request_body, headers=headers, verify=False)
    status_code = response.status_code
    if status_code == 200:
        return response.content, status_code
    else:
        agent_url = f'http://10.202.6.112:3007/{target}'
        logger.info(f'url[{url}] direct {target} failed\n using agent url[{agent_url}] {target}.')
        agent_response = requests.post(url=agent_url, data=request_body, headers=headers, verify=False)
        agent_status_code = agent_response.status_code
        if agent_status_code == 200:
            return agent_response.content, agent_status_code
        else:
            return_body = {
                "use agent login failed message":
                    {
                        "agent_url": agent_url,
                        "status_code": agent_status_code,
                        "response_body": agent_response.text
                    },
                "direct login failed message":
                    {
                        "login_url": url,
                        "status_code": status_code,
                        "response_body": response.text
                    }

            }
            return f'{return_body}', agent_status_code


@xbu_api.route('/zta_logout', methods=['POST'])
def zta_logout():
    data = request.get_json(force=True)
    email = data.get('email')
    cac_token = data.get('cac_token')
    endpoint = data.get('endpoint')
    logger.info(f"email: {email}, endpoint: {endpoint}, cac_token: {cac_token}")
    if None in [email, cac_token, endpoint]:
        return 'The request body is missing a required parameter.', 404

    password = get_zta_account(email)
    if password == 404:
        return f"The user of '{email}' email was not queried!", 404

    request_body = {"password": f"{password}", "email": email}
    headers = {"Authorization": cac_token}
    response_content, status_code = login_or_logout(endpoint, "logout", request_body, headers)
    return response_content, status_code


@xbu_api.route('/cac_token/<email>', methods=['GET'])
def get_cac_token(email):
    token_dic = {}
    if email in token_dic:
        return token_dic[email]
    else:
        return f"The cac_token of '{email}' email was not queried!", 404
