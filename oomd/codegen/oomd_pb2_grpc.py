# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import oomd_pb2 as oomd__pb2


class OomDStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.OnlineGet = channel.unary_unary(
                '/oomd.OomD/OnlineGet',
                request_serializer=oomd__pb2.OnlineGetRequest.SerializeToString,
                response_deserializer=oomd__pb2.OnlineGetResponse.FromString,
                )
        self.OnlineMultiGet = channel.unary_unary(
                '/oomd.OomD/OnlineMultiGet',
                request_serializer=oomd__pb2.OnlineMultiGetRequest.SerializeToString,
                response_deserializer=oomd__pb2.OnlineMultiGetResponse.FromString,
                )
        self.Sync = channel.unary_unary(
                '/oomd.OomD/Sync',
                request_serializer=oomd__pb2.SyncRequest.SerializeToString,
                response_deserializer=oomd__pb2.SyncResponse.FromString,
                )
        self.Import = channel.stream_unary(
                '/oomd.OomD/Import',
                request_serializer=oomd__pb2.ImportRequest.SerializeToString,
                response_deserializer=oomd__pb2.ImportResponse.FromString,
                )
        self.Join = channel.unary_stream(
                '/oomd.OomD/Join',
                request_serializer=oomd__pb2.JoinRequest.SerializeToString,
                response_deserializer=oomd__pb2.JoinResponse.FromString,
                )
        self.ImportByFile = channel.unary_unary(
                '/oomd.OomD/ImportByFile',
                request_serializer=oomd__pb2.ImportByFileRequest.SerializeToString,
                response_deserializer=oomd__pb2.ImportResponse.FromString,
                )
        self.JoinByFile = channel.unary_unary(
                '/oomd.OomD/JoinByFile',
                request_serializer=oomd__pb2.JoinRequest.SerializeToString,
                response_deserializer=oomd__pb2.JoinByFileResponse.FromString,
                )


class OomDServicer(object):
    """Missing associated documentation comment in .proto file."""

    def OnlineGet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def OnlineMultiGet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Sync(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Import(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Join(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ImportByFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def JoinByFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OomDServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'OnlineGet': grpc.unary_unary_rpc_method_handler(
                    servicer.OnlineGet,
                    request_deserializer=oomd__pb2.OnlineGetRequest.FromString,
                    response_serializer=oomd__pb2.OnlineGetResponse.SerializeToString,
            ),
            'OnlineMultiGet': grpc.unary_unary_rpc_method_handler(
                    servicer.OnlineMultiGet,
                    request_deserializer=oomd__pb2.OnlineMultiGetRequest.FromString,
                    response_serializer=oomd__pb2.OnlineMultiGetResponse.SerializeToString,
            ),
            'Sync': grpc.unary_unary_rpc_method_handler(
                    servicer.Sync,
                    request_deserializer=oomd__pb2.SyncRequest.FromString,
                    response_serializer=oomd__pb2.SyncResponse.SerializeToString,
            ),
            'Import': grpc.stream_unary_rpc_method_handler(
                    servicer.Import,
                    request_deserializer=oomd__pb2.ImportRequest.FromString,
                    response_serializer=oomd__pb2.ImportResponse.SerializeToString,
            ),
            'Join': grpc.unary_stream_rpc_method_handler(
                    servicer.Join,
                    request_deserializer=oomd__pb2.JoinRequest.FromString,
                    response_serializer=oomd__pb2.JoinResponse.SerializeToString,
            ),
            'ImportByFile': grpc.unary_unary_rpc_method_handler(
                    servicer.ImportByFile,
                    request_deserializer=oomd__pb2.ImportByFileRequest.FromString,
                    response_serializer=oomd__pb2.ImportResponse.SerializeToString,
            ),
            'JoinByFile': grpc.unary_unary_rpc_method_handler(
                    servicer.JoinByFile,
                    request_deserializer=oomd__pb2.JoinRequest.FromString,
                    response_serializer=oomd__pb2.JoinByFileResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'oomd.OomD', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OomD(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def OnlineGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomd.OomD/OnlineGet',
            oomd__pb2.OnlineGetRequest.SerializeToString,
            oomd__pb2.OnlineGetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def OnlineMultiGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomd.OomD/OnlineMultiGet',
            oomd__pb2.OnlineMultiGetRequest.SerializeToString,
            oomd__pb2.OnlineMultiGetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Sync(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomd.OomD/Sync',
            oomd__pb2.SyncRequest.SerializeToString,
            oomd__pb2.SyncResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Import(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/oomd.OomD/Import',
            oomd__pb2.ImportRequest.SerializeToString,
            oomd__pb2.ImportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Join(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/oomd.OomD/Join',
            oomd__pb2.JoinRequest.SerializeToString,
            oomd__pb2.JoinResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ImportByFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomd.OomD/ImportByFile',
            oomd__pb2.ImportByFileRequest.SerializeToString,
            oomd__pb2.ImportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def JoinByFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomd.OomD/JoinByFile',
            oomd__pb2.JoinRequest.SerializeToString,
            oomd__pb2.JoinByFileResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
