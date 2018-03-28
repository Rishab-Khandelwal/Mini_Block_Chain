# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc
import imp
import db_pb2 as db__pb2
# db__pb2 = imp.load_source('db_pb2', './protobuf/db_pb2.py')

class BlockChainMinerStub(object):
  """Interface exported by the server.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Get = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/Get',
        request_serializer=db__pb2.GetRequest.SerializeToString,
        response_deserializer=db__pb2.GetResponse.FromString,
        )
    self.Transfer = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/Transfer',
        request_serializer=db__pb2.Transaction.SerializeToString,
        response_deserializer=db__pb2.BooleanResponse.FromString,
        )
    self.Verify = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/Verify',
        request_serializer=db__pb2.Transaction.SerializeToString,
        response_deserializer=db__pb2.VerifyResponse.FromString,
        )
    self.GetHeight = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/GetHeight',
        request_serializer=db__pb2.Null.SerializeToString,
        response_deserializer=db__pb2.GetHeightResponse.FromString,
        )
    self.GetBlock = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/GetBlock',
        request_serializer=db__pb2.GetBlockRequest.SerializeToString,
        response_deserializer=db__pb2.JsonBlockString.FromString,
        )
    self.PushBlock = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/PushBlock',
        request_serializer=db__pb2.JsonBlockString.SerializeToString,
        response_deserializer=db__pb2.Null.FromString,
        )
    self.PushTransaction = channel.unary_unary(
        '/miniblockchaindb.BlockChainMiner/PushTransaction',
        request_serializer=db__pb2.Transaction.SerializeToString,
        response_deserializer=db__pb2.Null.FromString,
        )


class BlockChainMinerServicer(object):
  """Interface exported by the server.
  """

  def Get(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Transfer(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Verify(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetHeight(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetBlock(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PushBlock(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PushTransaction(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_BlockChainMinerServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Get': grpc.unary_unary_rpc_method_handler(
          servicer.Get,
          request_deserializer=db__pb2.GetRequest.FromString,
          response_serializer=db__pb2.GetResponse.SerializeToString,
      ),
      'Transfer': grpc.unary_unary_rpc_method_handler(
          servicer.Transfer,
          request_deserializer=db__pb2.Transaction.FromString,
          response_serializer=db__pb2.BooleanResponse.SerializeToString,
      ),
      'Verify': grpc.unary_unary_rpc_method_handler(
          servicer.Verify,
          request_deserializer=db__pb2.Transaction.FromString,
          response_serializer=db__pb2.VerifyResponse.SerializeToString,
      ),
      'GetHeight': grpc.unary_unary_rpc_method_handler(
          servicer.GetHeight,
          request_deserializer=db__pb2.Null.FromString,
          response_serializer=db__pb2.GetHeightResponse.SerializeToString,
      ),
      'GetBlock': grpc.unary_unary_rpc_method_handler(
          servicer.GetBlock,
          request_deserializer=db__pb2.GetBlockRequest.FromString,
          response_serializer=db__pb2.JsonBlockString.SerializeToString,
      ),
      'PushBlock': grpc.unary_unary_rpc_method_handler(
          servicer.PushBlock,
          request_deserializer=db__pb2.JsonBlockString.FromString,
          response_serializer=db__pb2.Null.SerializeToString,
      ),
      'PushTransaction': grpc.unary_unary_rpc_method_handler(
          servicer.PushTransaction,
          request_deserializer=db__pb2.Transaction.FromString,
          response_serializer=db__pb2.Null.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'miniblockchaindb.BlockChainMiner', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))