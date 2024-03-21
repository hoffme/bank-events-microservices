from context import AppContext

from application.queues.account_balance_modify import ResolverAccountBalanceModify

class Queues:
    def __init__(self, ctx: AppContext):
        self.ctx = ctx

        self.subscribe()

    def subscribe(self):
        self.ctx.event_bus.subscribe(ResolverAccountBalanceModify(self.ctx))
        pass