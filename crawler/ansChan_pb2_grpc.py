# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import ansChan_pb2 as ansChan__pb2


class AnsServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ProcessAnsList = channel.unary_unary(
                '/main.AnsService/ProcessAnsList',
                request_serializer=ansChan__pb2.AnsList.SerializeToString,
                response_deserializer=ansChan__pb2.Ans.FromString,
                )


class AnsServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ProcessAnsList(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AnsServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ProcessAnsList': grpc.unary_unary_rpc_method_handler(
                    servicer.ProcessAnsList,
                    request_deserializer=ansChan__pb2.AnsList.FromString,
                    response_serializer=ansChan__pb2.Ans.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'main.AnsService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AnsService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ProcessAnsList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.AnsService/ProcessAnsList',
            ansChan__pb2.AnsList.SerializeToString,
            ansChan__pb2.Ans.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
