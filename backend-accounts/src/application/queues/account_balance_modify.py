from context import AppContext
from domain.shared.events import Event
from domain.shared.resolver import EventResolver

class ResolverAccountBalanceModify(EventResolver):
    def __init__(self, ctx: AppContext):
        super().__init__(
            name="app.accounts.qeu.account.update.balance",
            topics=["app.transactions.evt.transactions.account.balance.changed"]
        )
        self.ctx = ctx

    def resolve(self, event: Event):
        print(self.id, "resolve", event)

        account = self.ctx.account_repository.find(event.payload["id"])
        if account is None:
            return

        account.setBalance(event.payload["balance"])

        self.ctx.account_repository.save(account)