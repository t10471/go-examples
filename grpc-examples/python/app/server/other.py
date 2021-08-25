import asyncio
import logging

import grpc
import proto.other_pb2 as other_pb2
import proto.other_pb2_grpc as other_pb2_grpc


class Other:
    def main(self, my_address: str, sleep_sec: int):
        asyncio.run(serve(my_address, sleep_sec))


class OtherCallee(other_pb2_grpc.OtherServicer):
    def __init__(self, sleep_sec: int):
        self.sleep_sec = sleep_sec

    async def CallOther(
        self, request: other_pb2.OtherRequest, context: grpc.aio.ServicerContext
    ) -> other_pb2.OtherReply:
        logging.info("start sleep")
        await asyncio.sleep(self.sleep_sec)
        logging.info("end sleep")
        return other_pb2.OtherReply(message="Hello, %s!" % request.name)

    async def CallOtherV2(
        self, request: other_pb2.OtherRequest, context: grpc.aio.ServicerContext
    ) -> other_pb2.OtherReply:
        return other_pb2.OtherReply(message="Hello, %s!" % request.name)


async def serve(my_address: str, sleep_sec: int) -> None:
    server = grpc.aio.server()
    other_pb2_grpc.add_OtherServicer_to_server(OtherCallee(sleep_sec), server)
    listen_addr = my_address
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()
