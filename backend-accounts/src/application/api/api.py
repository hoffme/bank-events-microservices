from sanic import Sanic
from config import AppConfig
from context import AppContext
from application.api.accounts import AccountsControllers

class Api:
    def __init__(self, config: AppConfig, ctx: AppContext):
        self.app = Sanic("bank-accounts", config=config, ctx=ctx)

        self.controllers = [
            AccountsControllers(ctx=ctx)
        ]

        self._start()

    def _start(self):
        for controller in self.controllers:
            controller.create_routes(self.app)

    def run(self):
        self.app.run(
            host="0.0.0.0",
            port=int(self.app.config.PORT),
            dev=self.app.config.DEV,
            auto_reload=self.app.config.DEV
        )