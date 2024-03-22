from uuid import UUID
from sanic import Request, Sanic, json
from context import AppContext
from domain.accounts.models.account.model import AccountModel
from application.api.accounts.validations import SearchAccountQueryDto, CreateAccountBodyDto, ModifiyAccountBodyDto

class AccountsControllers:
    def __init__(self, ctx: AppContext):
        self.ctx = ctx
        self.path = "/api/v1/accounts"

    def create_routes(self, app: Sanic):
        app.add_route(self._search, f"{self.path}", methods=["GET"])
        app.add_route(self._find, f"{self.path}/<account_id:uuid>", methods=["GET"])
        app.add_route(self._save, f"{self.path}/<account_id:uuid>", methods=["PUT"])
        app.add_route(self._activate, f"{self.path}/<account_id:uuid>/activate", methods=["PUT"])
        app.add_route(self._inactivate, f"{self.path}/<account_id:uuid>/inactivate", methods=["PUT"])

    def _search(self, request: Request):
        filter = SearchAccountQueryDto(**dict(request.query_args))

        result = self.ctx.account_repository.search(filter)

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
        id = str(account_id)

        result = self.ctx.account_repository.find(id)
        if result is None:
            return json({ "error": { "code": "ACCOUNT:NOT_FOUND" } }, status=404)

        return json({ "result": result.raw() })

    def _save(self, request: Request, account_id: UUID):
        id = str(account_id)

        account = self.ctx.account_repository.find(id)
        if account is None:
            params = CreateAccountBodyDto(**request.json)

            account = AccountModel.create(
                id=id, 
                name=params.name, 
                currency=params.currency,
                balance=params.balance
            )
        else:
            params = ModifiyAccountBodyDto(**request.json)
            
            account.setName(params.name)

        self.ctx.account_repository.save(account)

        return json({ "result": account.raw() })
    
    def _inactivate(self, request: Request, account_id: UUID):
        body = request.json

        account = self.ctx.account_repository.find(str(account_id))
        if account is None:
            return json({ "error": { "code": "ACCOUNT:NOT_FOUND" } }, status=404)

        account.inactivate()

        self.ctx.account_repository.save(account)

        return json({ "result": account.raw() })
    
    def _activate(self, request: Request, account_id: UUID):
        body = request.json

        account = self.ctx.account_repository.find(str(account_id))
        if account is None:
            return json({ "error": { "code": "ACCOUNT:NOT_FOUND" } }, status=404)

        account.activate()

        self.ctx.account_repository.save(account)

        return json({ "result": account.raw() })
