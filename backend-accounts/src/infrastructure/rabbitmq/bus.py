import json
import threading
import pika
from domain.shared.events import Event
from domain.shared.resolver import EventResolver

class RabbitMQEventBus:
    def __init__(self, host: str, port: str, user: str, password: str, exchange: str):
        self.exchange = exchange
        self.conn = pika.BlockingConnection(pika.ConnectionParameters(
            host=host,
            port=port,
            credentials=pika.PlainCredentials(user, password)
        ))
        self.channel = self.conn.channel()

        self.channel.exchange_declare(
            self.exchange, 
            exchange_type="topic",
            durable=True,
            auto_delete=False,
            internal=False
        )

    def dispatch(self, events: list):
        for event in events:
            self.channel.basic_publish(
                exchange=self.exchange,
                routing_key=event.topic,
                body=json.dumps(event.raw()),
                properties=pika.BasicProperties(
                    message_id=event.id,
                )
            )

    def subscribe(self, resolver: EventResolver):
        if len(resolver.topics) == 0:
            return

        self.channel.queue_declare(
            queue=resolver.name,
            durable=True,
        )

        for topic in resolver.topics:
            self.channel.queue_bind(
                queue=resolver.name,
                exchange=self.exchange,
                routing_key=topic
            )
            

        self.channel.basic_consume(
            queue=resolver.name, 
            auto_ack=True, 
            on_message_callback=lambda c, m, p, body: resolver.resolve(Event.from_raw(json.loads(body.decode('ascii')))),
        )

    def consume(self):
        self.consumer_thread = threading.Thread(target=lambda: self.channel.start_consuming())
        self.consumer_thread.start()

    def close(self):
        self.channel.stop_consuming()
        