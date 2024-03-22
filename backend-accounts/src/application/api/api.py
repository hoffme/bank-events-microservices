from pydantic import ValidationError
from sanic import Sanic, json
from config import AppConfig
from context import AppContext
from application.api.accounts.controller import AccountsControllers

class Api:
    def __init__(self, config: AppConfig, ctx: AppContext):
        self.app = Sanic("bank-accounts", config=config, ctx=ctx)

        self.app.error_handler.add(handler=self._error_handler,exception=Exception)

        self.controllers = [
            AccountsControllers(ctx=ctx)
        ]

        self._start()

    def _error_handler(self, request, exception):
        if self.app.config["DEV"]:
            print(exception)

        if isinstance(exception, ValidationError):
            return json({"errors": exception.errors(), "status": 400}, status=400)

        return json({"errors": [{ "code": "internal" }], "status": 500}, status=500)
    
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