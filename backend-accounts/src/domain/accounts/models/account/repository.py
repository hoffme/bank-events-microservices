from domain.accounts.models.account.model import AccountModel


class AccountRepository:
    def find(self, id: str) -> AccountModel:
        pass

    def search(self, id: str) -> [AccountModel]: # type: ignore
        pass

    def save(self, account: AccountModel):
        pass

    
