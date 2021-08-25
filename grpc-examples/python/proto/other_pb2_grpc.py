# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import other_pb2 as other__pb2


class OtherStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CallOther = channel.unary_unary(
                '/other.Other/CallOther',
                request_serializer=other__pb2.OtherRequest.SerializeToString,
                response_deserializer=other__pb2.OtherReply.FromString,
                )
        self.CallOtherV2 = channel.unary_unary(
                '/other.Other/CallOtherV2',
                request_serializer=other__pb2.OtherRequest.SerializeToString,
                response_deserializer=other__pb2.OtherReply.FromString,
                )


class OtherServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CallOther(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CallOtherV2(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OtherServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CallOther': grpc.unary_unary_rpc_method_handler(
                    servicer.CallOther,
                    request_deserializer=other__pb2.OtherRequest.FromString,
                    response_serializer=other__pb2.OtherReply.SerializeToString,
            ),
            'CallOtherV2': grpc.unary_unary_rpc_method_handler(
                    servicer.CallOtherV2,
                    request_deserializer=other__pb2.OtherRequest.FromString,
                    response_serializer=other__pb2.OtherReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'other.Other', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Other(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CallOther(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/other.Other/CallOther',
            other__pb2.OtherRequest.SerializeToString,
            other__pb2.OtherReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CallOtherV2(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/other.Other/CallOtherV2',
            other__pb2.OtherRequest.SerializeToString,
            other__pb2.OtherReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)