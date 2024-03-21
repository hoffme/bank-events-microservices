import base64
import json
from uuid import uuid4
from datetime import datetime, timezone

class Event:
    @staticmethod
    def from_raw(raw):
        return Event(
            topic=raw["header"]["topic"],
            id=raw["header"]["id"],
            timestamp=datetime.fromisoformat(raw["header"]["timestamp"]),
            payload=json.loads(base64.b64decode(raw["payload"]))
        )

    def __init__(self, topic: str, id: str = str(uuid4()), timestamp: datetime = datetime.now(timezone.utc), payload = None):
        self.id = id
        self.topic = topic
        self.timestamp = timestamp
        self.payload = payload

    def raw(self):
        return {
            "header": {
                "id": self.id,
                "topic": self.topic,
                "timestamp": self.timestamp.isoformat()
            },
            "payload": base64.b64encode(json.dumps(self.payload).encode('utf-8')).decode('ascii')
        }