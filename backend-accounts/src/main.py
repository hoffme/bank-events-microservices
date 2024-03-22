from config import AppConfig
from context import AppContext

from application.api.api import Api
from application.queues.main import Queues
from infrastructure.infrastructure import Infrastructure

config = AppConfig()

infra = Infrastructure(config=config)

ctx = AppContext(dependencies=infra.dependencies())

api = Api(config, ctx)
queues = Queues(ctx=ctx)

if __name__ == '__main__':
    queues.run()
    api.run()
    queues.close()
    