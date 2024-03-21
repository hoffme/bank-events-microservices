import pymongo

from domain.accounts.models.account.model import AccountModel
from domain.accounts.models.account.repository import AccountRepository
from infrastructure.mongodb.connection import MongoConnection

class MongoAccountRepository(AccountRepository):
    def __init__(self, conn: MongoConnection, dispatcher):
        super().__init__()

        self.dispatcher = dispatcher
        self.conn = conn
        self.databaseName = "bank_accounts"
        self.collectionName = "accounts"

    def _collection(self):
        return self.conn.client.get_database(self.databaseName).get_collection(self.collectionName)
    
    def _encode(self, doc: any) -> AccountModel:
        return AccountModel({
            "id": doc["_id"],
            "name": doc["name"],
            "currency": doc["currency"],
            "active": doc["active"],
            "balance": doc["balance"],
            "created_at": doc["created_at"],
        })

    def _decode(self, model: AccountModel) -> any:
        doc = model.raw()
        doc["_id"] = doc["id"]
        del doc["id"]
        return doc

    def find(self, id: str) -> AccountModel:
        doc = self._collection().find_one({ "_id": id })
        if doc is None:
            return None

        return self._encode(doc)

    def search(self, filter):
        mongo_filter = {}

        if "query" in filter:
            if "$or" not in mongo_filter:
                mongo_filter["$or"] = []

            mongo_filter["$or"].append({ "$text": { "$search": filter["query"] } })

        if "where" in filter:
            if "$or" not in mongo_filter:
                mongo_filter["$or"] = []
                
            mongo_filter["$or"].append(filter["where"])
        
        count = self._collection().count_documents(mongo_filter)

        qry = self._collection().find(mongo_filter)

        if "order_by" in filter:
            dir = pymongo.ASCENDING
            if "order_dir" in filter and filter["order_dir"] == "desc":
                dir = pymongo.DESCENDING

            qry = qry.sort(filter["order_by"], dir)

        limit = filter["limit"] if "limit" in filter else 10
        qry = qry.limit(limit)

        skip = filter["skip"] if "skip" in filter else 0
        qry = qry.skip(skip)

        result = {
            "data": [],
            "count": count,
            "limit": limit,
            "skip": skip
        }

        for doc in qry:
            result["data"].append(self._encode(doc))

        return result

    def save(self, account: AccountModel):
        raw = self._decode(account)
        self._collection().update_one({ "_id": raw["_id"] }, { "$set": raw }, upsert=True)

        if self.dispatcher is not None:
            self.dispatcher.dispatch(account.pull_events())

    