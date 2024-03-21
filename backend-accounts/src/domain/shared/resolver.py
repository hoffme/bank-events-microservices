from uuid import uuid4
from domain.shared.events import Event

class EventResolver:
    def __init__(self, name: str, topics: list, id: str = str(uuid4())):
        self.name = name
        self.topics = topics
        self.id = id

    def resolve(self, event: Event):
        pass