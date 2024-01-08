from pymongo import MongoClient
# from urllib.parse import quote_plus


class ExecMongoDb:
    def __init__(self, connect_db_data):
        host = connect_db_data['host']
        port = connect_db_data['port']
        username = connect_db_data['user']
        password = connect_db_data['password']
        db_name = connect_db_data['db_name']
        # client = MongoClient(host=host, port=port, username=username, password=password, authSource=db_name)
        # self.db = client.db_name

        # client = MongoClient(f'mongodb://{username}:{password}@{host}:{port}/?authSource={db_name}')
        # self.db = client.db_name

        self.client = MongoClient(host=host, port=int(port))
        self.db = self.client[db_name]
        self.db.authenticate(username, password)

    def get_collection_names(self):
        return self.db.list_collection_names(session=None)

    def close_db(self):
        self.client.close()
        print('Cleanup client resources and disconnect from MongoDB.')


