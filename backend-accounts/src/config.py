from os import getenv
from dotenv import load_dotenv
from sanic.config import Config

class AppConfig(Config):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        
        load_dotenv()
        
        self.DEV = getenv("MODE") != "production"
        self.PORT = getenv("PORT")
        self.MONGO_URI = getenv("MONGO_URI")
        self.RABBIT_HOST = getenv("RABBIT_HOST")
        self.RABBIT_PORT = getenv("RABBIT_PORT")
        self.RABBIT_USER = getenv("RABBIT_USER")
        self.RABBIT_PASS = getenv("RABBIT_PASS")
        self.RABBIT_EXCHANGE = getenv("RABBIT_EXCHANGE")