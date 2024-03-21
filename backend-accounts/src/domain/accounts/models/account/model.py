from datetime import datetime, timezone
from domain.shared.events import Event

class AccountModel:
    @staticmethod
    def create(id: str, name: str, currency: str, active: bool = True, balance: int = 0):
        return AccountModel({
            "id": id,
            "name": name,
            "active": active,
            "currency": currency,
            "balance": balance,
            "created_at": datetime.now(timezone.utc).isoformat(),
        }, is_new=True)
    
    def __init__(self, raw, is_new: bool = False):
        self._events = []

        self.id = raw["id"]
        self.name = raw["name"]
        self.active = raw["active"]
        self.currency = raw["currency"]
        self.balance = raw["balance"]
        self.created_at = raw["created_at"]

        if is_new:
            self._events.append(
                Event(topic="app.accounts.evt.account.created", payload={
                    "id": self.id,
                    "name": self.name,
                    "active": self.active,
                    "currency": self.currency,
                    "balance": self.balance,
                    "created_at": self.created_at
                })
            )
        
    def setName(self, name: str):
        if name == self.name:
            return
        
        self.name = name

        self._events.append(
            Event(topic="app.accounts.evt.account.name.changed", payload={
                "id": self.id,
                "name": self.name,
            })
        )

    def activate(self):
        if self.active:
            return

        self.active = True 

        self._events.append(
            Event(topic="app.accounts.evt.account.activated", payload={
                "id": self.id,
                "name": self.name,
            })
        )

    def inactivate(self):
        if self.active is False:
            return

        self.active = False 

        self._events.append(
            Event(topic="app.accounts.evt.account.inactivate", payload={
                "id": self.id,
                "name": self.name,
            })
        )

    def setBalance(self, balance: int):
        if self.balance == balance:
            return

        self.balance = balance 

        self._events.append(
            Event(topic="app.accounts.evt.account.balance.changed", payload={
                "id": self.id,
                "balance": self.balance,
            })
        )

    def pull_events(self):
        result, self._events[:] = self._events[:], []
        return result

    def raw(self):
        return {
            "id": self.id,
            "name": self.name,
            "active": self.active,
            "currency": self.currency,
            "balance": self.balance,
            "created_at": self.created_at
        }