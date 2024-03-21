class AppContext:
    def __init__(self, dependencies):
        self.account_repository = dependencies["repositories"]["account"]
        self.event_bus = dependencies["ports"]["event_bus"]