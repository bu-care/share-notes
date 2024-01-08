# import subprocess
import time
from datetime import datetime

import paramiko
from collections import namedtuple
import logging
# import os

cur_time = datetime.now().strftime('%Y-%m-%d_%H-%M-%S')
logger = logging.getLogger('app_log.update_zta')


def paramiko_ssh_connect(connect_data):
    user = connect_data['user']
    passwd = connect_data['passwd']
    client = paramiko.SSHClient()
    client.load_system_host_keys()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect("10.103.12.238", connect_data['port'], user, passwd, timeout=5)

    return client


class UpdateZtaInDocker:
    stdout_info = None

    def __init__(self):
        self.zta_dir = '/root/code/zta/'
        self.activate_zta = 'source /root/python_env/ztna/bin/activate'
        self.web_zta_dir = '/root/code/web-zta/'

    def set_cli_list(self, port='9000'):
        UpdateZtaInDocker.stdout_info = ['******start server******\n']
        self.port = str(port)

        CliCommand = namedtuple('CliCommand', ['fetch_rebase', 'install_requirements',
                                               'installTools', 'start_zta_server'])
        cmd_list = [f"cd {self.zta_dir}; git fetch; git rebase",
                    f"{self.activate_zta}; cd {self.zta_dir}/awsApiGatewaySimulator; pip install -r requirements.txt",
                    f"cd {self.zta_dir}/tools; ./installTools.sh",
                    f"{self.activate_zta}; cd {self.zta_dir}/awsApiGatewaySimulator;"
                    f"nohup python3.8 main.py --port {port}"] # >/root/code/log_zta/{cur_time}_main.log &
        self.cli_cmd = CliCommand(*cmd_list)

        WebZtaCliCommand = namedtuple('WebZtaCliCommand', ['fetch_rebase', 'npm_install', 'start_web_zta_server'])
        web_cmd_list = [f"cd {self.web_zta_dir}; git stash; git fetch; git rebase",
                        f"cd {self.web_zta_dir}; npm install",
                        f"cd {self.web_zta_dir}; ZTNA_API_BASE_URL=http://10.103.12.238:{port} npm run dev"]

        self.web_cli_cmd = WebZtaCliCommand(*web_cmd_list)

    def ssh_connect(self):
        connect_data = {'user': 'root', 'passwd': 'wireless_dev', 'port': 222}
        self.client = paramiko_ssh_connect(connect_data)
        return self.client

    def check_zta_process(self, port):
        pid = self.client.exec_command(f'lsof -i:{port} -t')[1].read().decode('utf-8')
        info = f'pid of port {port} is {pid}'.strip('\n')
        if pid:
            logger.info(info)
        else:
            print(info)
        status_code = self.client.exec_command('echo $?')[1].read().decode('utf-8').strip()
        return status_code, pid

    def stop_server(self, port):
        status_code, pid = self.check_zta_process(port)
        logger.info(f"status code of linux shell with 'lsof -i:{port} -t': {status_code}")
        if status_code == '0' and pid:
            logger.info('process is alive.')
            self.client.exec_command(f'kill -9 $(lsof -i:{port} -t)')
            logger.info('kill the process.')
        else:
            logger.info('process is dead.')

    def update_code(self):
        _, stdout, _ = self.client.exec_command(self.cli_cmd.fetch_rebase)
        self.client_stdout(stdout)
        _, stdout, _ = self.client.exec_command(self.cli_cmd.install_requirements)
        self.client_stdout(stdout)

    def build_code(self):
        _, stdout, _ = self.client.exec_command(self.cli_cmd.installTools)
        self.client_stdout(stdout)

    def start_zta_server(self):
        self.client.exec_command(self.cli_cmd.start_zta_server)

        # with ThreadPoolExecutor(5) as executor:
        #     executor.submit(self.client.exec_command, self.cli_cmd.start_zta_server)

        # # start server
        # pool = Pool(1)
        # pool.apply_async(self.client.exec_command, (self.cli_cmd.start_zta_server,))
        # pool.close()  # Close the process pool and do not accept new processes.
        # # # pool.join() # The main process is blocked, waiting for the child process to exit.
        logger.info('The ZTA process is starting......')
        wait_flag = True
        while wait_flag:
            status_code, pid = self.check_zta_process(self.port)
            if status_code == '0' and pid:
                logger.info('The ZTA process has started!\n')
                time.sleep(2)
                wait_flag = False
            else:
                print('The ZTA process has not started yet and is waiting to start......')

    # def start_zta_frontend_server(self):
    #     self.stop_server('8080')
    #     _, stdout, _ = self.client.exec_command(self.web_cli_cmd.fetch_rebase)
    #     self.client_stdout(stdout, is_backend=False)
    #     _, stdout, _ = self.client.exec_command(self.web_cli_cmd.npm_install, get_pty=True)
    #     while not stdout.channel.exit_status_ready():  # Monitoring terminal output
    #         console_output = stdout.readline()
    #         print("\n----------------------\n", console_output, "^^^^^^^^^^^^^^^^^^^^^^^^")
    #
    #     _, stdout, _ = self.client.exec_command(self.web_cli_cmd.start_web_zta_server)
    #     logger.info('The ZTA WEB process is starting......')
    #
    #     wait_flag = True
    #     start_time = time.time()
    #     end_time = time.time()
    #     while wait_flag and (end_time - start_time) < 20:
    #         status_code, pid = self.check_zta_process('8080')
    #         if status_code == '0' and pid:
    #             logger.info(f'The ZTA WEB process has started and spend time is {end_time - start_time} s.')
    #             time.sleep(2)
    #             wait_flag = False
    #         else:
    #             logger.info('The ZTA WEB process has not started yet and is waiting to start......')
    #             end_time = time.time()

    def start_server(self):
        self.stop_server(self.port)
        self.update_code()
        self.build_code()
        self.start_zta_server()
        # self.start_zta_frontend_server()

    @staticmethod
    def client_stdout(stdout, is_backend=True):
        stdout_lines = stdout.readlines()
        if is_backend:
            UpdateZtaInDocker.stdout_info.extend(stdout_lines)
        print("\n----------------------\n")
        for console_output in stdout_lines:
            print(console_output)
        print("^^^^^^^^^^^^^^^^^^^^^^^^")

    @staticmethod
    def get_stdout_info():
        # print(UpdateZtaInDocker.stdout_info)
        return UpdateZtaInDocker.stdout_info

    def close_client(self):
        # self.stop_server()
        self.client.close()
        logger.info('stop zta server and close paramiko client.\n\n\n')


class UpdateZtaIn238:
    stdout_info = None

    def __init__(self):
        self.zta_dir = '/home/xbu/files/code/zta/'
        self.web_zta_dir = '/home/xbu/files/code/web-zta/'

    def set_cli_list(self, port='9000'):
        UpdateZtaIn238.stdout_info = ['******start server******\n']
        self.port = str(port)
        # conda_path = '/home/xbu/Programs/anaconda3/bin/conda'
        CliCommand = namedtuple('CliCommand', ['fetch_rebase', 'install_requirements',
                                               'installTools', 'start_zta_server'])
        cmd_list = [f"cd {self.zta_dir}; git fetch; git rebase",
                    f"conda activate ztna; cd {self.zta_dir}/awsApiGatewaySimulator; pip install -r requirements.txt",
                    f"cd {self.zta_dir}/tools; ./installTools.sh",
                    f"conda activate ztna; cd {self.zta_dir}/awsApiGatewaySimulator;"
                    f"nohup python3.8 main.py --port {port}"] # >/home/xbu/files/code/log_zta/{cur_time}_main.log &
        self.cli_cmd = CliCommand(*cmd_list)

        # WebZtaCliCommand = namedtuple('WebZtaCliCommand', ['fetch_rebase', 'start_web_zta_server'])
        # web_cmd_list = [f"cd {self.web_zta_dir}; git fetch; git rebase",
        #                 f"cd {self.web_zta_dir}; npm install; ZTNA_API_BASE_URL=http://10.103.12.238:{port} npm run dev"]
        # #  >/home/xbu/files/code/log_web_zta/{cur_time}_main.log &
        WebZtaCliCommand = namedtuple('WebZtaCliCommand', ['fetch_rebase', 'npm_install', 'start_web_zta_server'])
        web_cmd_list = [f"cd {self.web_zta_dir}; git stash; git fetch; git rebase",
                        f"cd {self.web_zta_dir}; npm install",
                        f"cd {self.web_zta_dir}; ZTNA_API_BASE_URL=http://10.103.12.238:{port} npm run dev"]

        self.web_cli_cmd = WebZtaCliCommand(*web_cmd_list)

    def ssh_connect(self):
        connect_data = {'user': 'xbu', 'passwd': 'wireless_dev', 'port': 22}
        self.client = paramiko_ssh_connect(connect_data)
        return self.client

    def check_zta_process(self, port):
        pid = self.client.exec_command(f'lsof -i:{port} -t')[1].read().decode('utf-8')
        info = f'pid of port {port} is {pid}'.strip('\n')
        if pid:
            logger.info(info)
        else:
            print(info)
        status_code = self.client.exec_command('echo $?')[1].read().decode('utf-8').strip()
        return status_code, pid

    def stop_server(self, port):
        status_code, pid = self.check_zta_process(port)
        logger.info(f"status code of linux shell with 'lsof -i:{port} -t': {status_code}.")
        if status_code == '0' and pid:
            logger.info('process is alive.')
            self.client.exec_command(f'kill -9 $(lsof -i:{port} -t)')
            logger.info('kill the process.')
        else:
            logger.info('process is dead.')
        # self.client.exec_command(f'kill -9 $(lsof -i:{self.port} -t)')

    def update_code(self):
        _, stdout, _ = self.client.exec_command(self.cli_cmd.fetch_rebase)
        self.client_stdout(stdout)
        _, stdout, _ = self.client.exec_command(self.cli_cmd.install_requirements)
        self.client_stdout(stdout)

    def build_code(self):
        _, stdout, _ = self.client.exec_command(self.cli_cmd.installTools)
        self.client_stdout(stdout)

    def start_zta_server(self):
        self.client.exec_command(self.cli_cmd.start_zta_server)

        # with ThreadPoolExecutor(5) as executor:
        #     executor.submit(self.client.exec_command, self.cli_cmd.start_zta_server)

        # # start server
        # pool = Pool(1)
        # pool.apply_async(self.client.exec_command, (self.cli_cmd.start_zta_server,))
        # pool.close()  # Close the process pool and do not accept new processes.
        # # # pool.join() # The main process is blocked, waiting for the child process to exit.
        wait_flag = True
        while wait_flag:
            status_code, pid = self.check_zta_process(self.port)
            if status_code == '0' and pid:
                logger.info('The ZTA process has started!\n')
                time.sleep(2)
                wait_flag = False
            else:
                print('The ZTA process has not started yet and is waiting to start......')

    def start_zta_frontend_server(self):
        self.stop_server('8080')
        _, stdout, _ = self.client.exec_command(self.web_cli_cmd.fetch_rebase)
        self.client_stdout(stdout, is_backend=False)
        _, stdout, _ = self.client.exec_command(self.web_cli_cmd.npm_install, get_pty=True)
        while not stdout.channel.exit_status_ready():  # Monitoring terminal output
            console_output = stdout.readline()
            print("\n----------------------\n", console_output, "^^^^^^^^^^^^^^^^^^^^^^^^")

        _, stdout, _ = self.client.exec_command(self.web_cli_cmd.start_web_zta_server)
        logger.info('The ZTA WEB process is starting......')

        wait_flag = True
        start_time = time.time()
        end_time = time.time()
        while wait_flag and (end_time - start_time) < 30:
            status_code, pid = self.check_zta_process('8080')
            if status_code == '0' and pid:
                logger.info(f'The ZTA WEB process has started and spend time is {end_time - start_time} s.\n')
                time.sleep(2)
                wait_flag = False
            else:
                print('The ZTA WEB process has not started yet and is waiting to start......')
                end_time = time.time()

    def start_server(self):
        self.stop_server(self.port)
        self.update_code()
        self.build_code()
        self.start_zta_server()
        self.start_zta_frontend_server()

    @staticmethod
    def client_stdout(stdout, is_backend=True):
        stdout_lines = stdout.readlines()
        if is_backend:
            UpdateZtaIn238.stdout_info.extend(stdout_lines)
        print("\n----------------------\n")
        for console_output in stdout_lines:
            print(console_output)
        print("^^^^^^^^^^^^^^^^^^^^^^^^")

    @staticmethod
    def get_stdout_info():
        # print(UpdateZtaIn238.stdout_info)
        return UpdateZtaIn238.stdout_info

    def close_client(self):
        # self.stop_server()
        self.client.close()
        logger.info('stop zta server and close paramiko client.\n\n\n')


