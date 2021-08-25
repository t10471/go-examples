import asyncio
import logging

import grpc
import proto.helloworld_pb2 as helloworld_pb2
import proto.helloworld_pb2_grpc as helloworld_pb2_grpc
import proto.other_pb2 as other_pb2
import proto.other_pb2_grpc as other_pb2_grpc


class HelloWorld:
    def main(self, my_address: str, other_address: str):
        asyncio.run(serve(my_address, other_address))


class Greeter(helloworld_pb2_grpc.GreeterServicer):
    def __init__(self, other_address: str):
        self.other_address = other_address

    async def SayHello(
        self, request: helloworld_pb2.HelloRequest, context: grpc.aio.ServicerContext
    ) -> helloworld_pb2.HelloReply:
        async with grpc.aio.insecure_channel(self.other_address) as channel:
            stub = other_pb2_grpc.OtherStub(channel)
            response = await stub.CallOther(other_pb2.OtherRequest(name="other"))
            logging.info(f"CallOther response {response.message}")
        async with grpc.aio.insecure_channel(self.other_address) as channel:
            stub = other_pb2_grpc.OtherStub(channel)
            response = await stub.CallOtherV2(other_pb2.OtherRequest(name="other v2"))
            logging.info(f"CallOtherV2 response {response.message}")
        res = helloworld_pb2.HelloReply(message="Hello, %s!" % request.name)
        return res


async def serve(my_address: str, other_address: str) -> None:
    server = grpc.aio.server()
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(other_address), server)
    listen_addr = my_address
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()
