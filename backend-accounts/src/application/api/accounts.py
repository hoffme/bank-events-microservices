from uuid import UUID
from sanic import Request, Sanic, json
from context import AppContext
from domain.accounts.models.account.model import AccountModel

class AccountsControllers:
    def __init__(self, ctx: AppContext):
        self.ctx = ctx
        self.path = "/api/v1/accounts"

    def create_routes(self, app: Sanic):
        app.add_route(self._search, f"{self.path}", methods=["GET"])
        app.add_route(self._find, f"{self.path}/<account_id:uuid>", methods=["GET"])
        app.add_route(self._save, f"{self.path}/<account_id:uuid>", methods=["PUT"])

    def _search(self, request: Request):
        result = self.ctx.account_repository.search({})

        data = list(map(lambda x: x.raw(), result["data"]))

        return json({
            "result": {
                "data": data,
                "count": result["count"],
                "limit": result["limit"],
                "skip": result["skip"],
            }
        })

    def _find(self, request: Request, account_id: UUID):
        result = self.ctx.account_repository.find(str(account_id))
        if result is None:
            return json({ "error": { "code": "ACCOUNT:NOT_FOUND" } }, status=404)

        return json({ "result": result.raw() })

    def _save(self, request: Request, account_id: UUID):
        body = request.json

        account = self.ctx.account_repository.find(str(account_id))
        if account is None:
            account = AccountModel.create(
                id=str(account_id), 
                name=body["name"], 
                currency=body["currency"],
                balance=body["balance"]
            )

        account.setName(body["name"])

        self.ctx.account_repository.save(account)

        return json({ "result": account.raw() })
