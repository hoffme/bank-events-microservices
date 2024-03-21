from config import AppConfig
from infrastructure.mongodb.account import MongoAccountRepository
from infrastructure.mongodb.connection import MongoConnection
from infrastructure.rabbitmq.bus import RabbitMQEventBus

class Infrastructure:
    def __init__(self, config: AppConfig):
        self.mongo_conn = MongoConnection(uri=config["MONGO_URI"])
        self.rabbitmq = RabbitMQEventBus(
            host=config["RABBIT_HOST"],
            port=config["RABBIT_PORT"],
            user=config["RABBIT_USER"],
            password=config["RABBIT_PASS"],
            exchange=config["RABBIT_EXCHANGE"],
        )

    def dependencies(self):
        return {
            "repositories": {
                "account": MongoAccountRepository(conn=self.mongo_conn, dispatcher=self.rabbitmq)
            },
            "ports": {
                "event_bus": self.rabbitmq
            }
        }