import asyncio
import logging

import grpc
import proto.helloworld_pb2 as helloworld_pb2
import proto.helloworld_pb2_grpc as helloworld_pb2_grpc


class Client:
    def main(self, address: str):
        asyncio.run(run(address))


async def run(address: str) -> None:
    async with grpc.aio.insecure_channel(address) as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        response_future = stub.SayHello(helloworld_pb2.HelloRequest(name="you"))
        await asyncio.sleep(1)
        response_future.cancel()
    try:
        response = await response_future
        logging.info(f"Greeter client received code: {response.code()}")
        logging.info(f"Greeter client received message: {response.message}")
    except asyncio.exceptions.CancelledError:
        logging.info("canceled")
