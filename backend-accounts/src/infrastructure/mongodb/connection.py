from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi

class MongoConnection:
    def __init__(self, uri: str):
        self.uri = uri
        self.client = MongoClient(uri, server_api=ServerApi('1'))

    def connect(self):
        self.client.admin.command('ping')