import logging

import app.client.client as client
import app.server.hello_world as hello_world
import app.server.other as other
import fire


class Commands:
    def client(self, address: str = "localhost:50051"):
        client.Client().main(address)

    def hello_world(self, my_address: str = "localhost:50051", other_address: str = "localhost:50052"):
        hello_world.HelloWorld().main(my_address, other_address)

    def other(self, address: str = "localhost:50052", sleep_sec: int = 3):
        other.Other().main(address, sleep_sec)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    fire.Fire(Commands)
