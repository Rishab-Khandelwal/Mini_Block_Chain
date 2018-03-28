# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: db.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='db.proto',
  package='miniblockchaindb',
  syntax='proto2',
  serialized_pb=_b('\n\x08\x64\x62.proto\x12\x0c\x62lockchaindb\"$\n\x0fGetBlockRequest\x12\x11\n\tBlockHash\x18\x01 \x02(\t\"\x1c\n\nGetRequest\x12\x0e\n\x06UserID\x18\x01 \x02(\t\"\x1c\n\x0bGetResponse\x12\r\n\x05Value\x18\x01 \x02(\x05\"5\n\x11GetHeightResponse\x12\x0e\n\x06Height\x18\x01 \x02(\x05\x12\x10\n\x08LeafHash\x18\x02 \x02(\t\"\"\n\x0f\x42ooleanResponse\x12\x0f\n\x07Success\x18\x01 \x02(\x08\"\x8c\x01\n\x0eVerifyResponse\x12\x34\n\x06Result\x18\x01 \x02(\x0e\x32$.miniblockchaindb.VerifyResponse.Results\x12\x11\n\tBlockHash\x18\x02 \x02(\t\"1\n\x07Results\x12\n\n\x06\x46\x41ILED\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\r\n\tSUCCEEDED\x10\x02\"\x06\n\x04Null\"\xae\x01\n\x0bTransaction\x12-\n\x04Type\x18\x01 \x02(\x0e\x32\x1f.miniblockchaindb.Transaction.Types\x12\x0e\n\x06\x46romID\x18\x03 \x02(\t\x12\x0c\n\x04ToID\x18\x04 \x02(\t\x12\r\n\x05Value\x18\x05 \x02(\x05\x12\x11\n\tMiningFee\x18\x06 \x02(\x05\x12\x0c\n\x04UUID\x18\x07 \x02(\t\"\"\n\x05Types\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0c\n\x08TRANSFER\x10\x05\"\x1f\n\x0fJsonBlockString\x12\x0c\n\x04Json\x18\x01 \x02(\t\"{\n\x05\x42lock\x12\x0f\n\x07\x42lockID\x18\x01 \x02(\x05\x12\x10\n\x08PrevHash\x18\x02 \x02(\t\x12/\n\x0cTransactions\x18\x03 \x03(\x0b\x32\x19.miniblockchaindb.Transaction\x12\x0f\n\x07MinerID\x18\x04 \x02(\t\x12\r\n\x05Nonce\x18\x05 \x02(\t2\xf2\x03\n\x0f\x42lockChainMiner\x12<\n\x03Get\x12\x18.miniblockchaindb.GetRequest\x1a\x19.miniblockchaindb.GetResponse\"\x00\x12\x46\n\x08Transfer\x12\x19.miniblockchaindb.Transaction\x1a\x1d.miniblockchaindb.BooleanResponse\"\x00\x12\x43\n\x06Verify\x12\x19.miniblockchaindb.Transaction\x1a\x1c.miniblockchaindb.VerifyResponse\"\x00\x12\x42\n\tGetHeight\x12\x12.miniblockchaindb.Null\x1a\x1f.miniblockchaindb.GetHeightResponse\"\x00\x12J\n\x08GetBlock\x12\x1d.miniblockchaindb.GetBlockRequest\x1a\x1d.miniblockchaindb.JsonBlockString\"\x00\x12@\n\tPushBlock\x12\x1d.miniblockchaindb.JsonBlockString\x1a\x12.miniblockchaindb.Null\"\x00\x12\x42\n\x0fPushTransaction\x12\x19.miniblockchaindb.Transaction\x1a\x12.miniblockchaindb.Null\"\x00')
)



_VERIFYRESPONSE_RESULTS = _descriptor.EnumDescriptor(
  name='Results',
  full_name='miniblockchaindb.VerifyResponse.Results',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FAILED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PENDING', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUCCEEDED', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=307,
  serialized_end=356,
)
_sym_db.RegisterEnumDescriptor(_VERIFYRESPONSE_RESULTS)

_TRANSACTION_TYPES = _descriptor.EnumDescriptor(
  name='Types',
  full_name='miniblockchaindb.Transaction.Types',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TRANSFER', index=1, number=5,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=507,
  serialized_end=541,
)
_sym_db.RegisterEnumDescriptor(_TRANSACTION_TYPES)


_GETBLOCKREQUEST = _descriptor.Descriptor(
  name='GetBlockRequest',
  full_name='miniblockchaindb.GetBlockRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='BlockHash', full_name='miniblockchaindb.GetBlockRequest.BlockHash', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=26,
  serialized_end=62,
)


_GETREQUEST = _descriptor.Descriptor(
  name='GetRequest',
  full_name='miniblockchaindb.GetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UserID', full_name='miniblockchaindb.GetRequest.UserID', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=64,
  serialized_end=92,
)


_GETRESPONSE = _descriptor.Descriptor(
  name='GetResponse',
  full_name='miniblockchaindb.GetResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Value', full_name='miniblockchaindb.GetResponse.Value', index=0,
      number=1, type=5, cpp_type=1, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=94,
  serialized_end=122,
)


_GETHEIGHTRESPONSE = _descriptor.Descriptor(
  name='GetHeightResponse',
  full_name='miniblockchaindb.GetHeightResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Height', full_name='miniblockchaindb.GetHeightResponse.Height', index=0,
      number=1, type=5, cpp_type=1, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='LeafHash', full_name='miniblockchaindb.GetHeightResponse.LeafHash', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=124,
  serialized_end=177,
)


_BOOLEANRESPONSE = _descriptor.Descriptor(
  name='BooleanResponse',
  full_name='miniblockchaindb.BooleanResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Success', full_name='miniblockchaindb.BooleanResponse.Success', index=0,
      number=1, type=8, cpp_type=7, label=2,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=179,
  serialized_end=213,
)


_VERIFYRESPONSE = _descriptor.Descriptor(
  name='VerifyResponse',
  full_name='miniblockchaindb.VerifyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Result', full_name='miniblockchaindb.VerifyResponse.Result', index=0,
      number=1, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='BlockHash', full_name='miniblockchaindb.VerifyResponse.BlockHash', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _VERIFYRESPONSE_RESULTS,
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=216,
  serialized_end=356,
)


_NULL = _descriptor.Descriptor(
  name='Null',
  full_name='miniblockchaindb.Null',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=358,
  serialized_end=364,
)


_TRANSACTION = _descriptor.Descriptor(
  name='Transaction',
  full_name='miniblockchaindb.Transaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Type', full_name='miniblockchaindb.Transaction.Type', index=0,
      number=1, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='FromID', full_name='miniblockchaindb.Transaction.FromID', index=1,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ToID', full_name='miniblockchaindb.Transaction.ToID', index=2,
      number=4, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Value', full_name='miniblockchaindb.Transaction.Value', index=3,
      number=5, type=5, cpp_type=1, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='MiningFee', full_name='miniblockchaindb.Transaction.MiningFee', index=4,
      number=6, type=5, cpp_type=1, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UUID', full_name='miniblockchaindb.Transaction.UUID', index=5,
      number=7, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TRANSACTION_TYPES,
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=367,
  serialized_end=541,
)


_JSONBLOCKSTRING = _descriptor.Descriptor(
  name='JsonBlockString',
  full_name='miniblockchaindb.JsonBlockString',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Json', full_name='miniblockchaindb.JsonBlockString.Json', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=543,
  serialized_end=574,
)


_BLOCK = _descriptor.Descriptor(
  name='Block',
  full_name='miniblockchaindb.Block',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='BlockID', full_name='miniblockchaindb.Block.BlockID', index=0,
      number=1, type=5, cpp_type=1, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PrevHash', full_name='miniblockchaindb.Block.PrevHash', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Transactions', full_name='miniblockchaindb.Block.Transactions', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='MinerID', full_name='miniblockchaindb.Block.MinerID', index=3,
      number=4, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Nonce', full_name='miniblockchaindb.Block.Nonce', index=4,
      number=5, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=576,
  serialized_end=699,
)

_VERIFYRESPONSE.fields_by_name['Result'].enum_type = _VERIFYRESPONSE_RESULTS
_VERIFYRESPONSE_RESULTS.containing_type = _VERIFYRESPONSE
_TRANSACTION.fields_by_name['Type'].enum_type = _TRANSACTION_TYPES
_TRANSACTION_TYPES.containing_type = _TRANSACTION
_BLOCK.fields_by_name['Transactions'].message_type = _TRANSACTION
DESCRIPTOR.message_types_by_name['GetBlockRequest'] = _GETBLOCKREQUEST
DESCRIPTOR.message_types_by_name['GetRequest'] = _GETREQUEST
DESCRIPTOR.message_types_by_name['GetResponse'] = _GETRESPONSE
DESCRIPTOR.message_types_by_name['GetHeightResponse'] = _GETHEIGHTRESPONSE
DESCRIPTOR.message_types_by_name['BooleanResponse'] = _BOOLEANRESPONSE
DESCRIPTOR.message_types_by_name['VerifyResponse'] = _VERIFYRESPONSE
DESCRIPTOR.message_types_by_name['Null'] = _NULL
DESCRIPTOR.message_types_by_name['Transaction'] = _TRANSACTION
DESCRIPTOR.message_types_by_name['JsonBlockString'] = _JSONBLOCKSTRING
DESCRIPTOR.message_types_by_name['Block'] = _BLOCK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetBlockRequest = _reflection.GeneratedProtocolMessageType('GetBlockRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETBLOCKREQUEST,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.GetBlockRequest)
  ))
_sym_db.RegisterMessage(GetBlockRequest)

GetRequest = _reflection.GeneratedProtocolMessageType('GetRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETREQUEST,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.GetRequest)
  ))
_sym_db.RegisterMessage(GetRequest)

GetResponse = _reflection.GeneratedProtocolMessageType('GetResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETRESPONSE,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.GetResponse)
  ))
_sym_db.RegisterMessage(GetResponse)

GetHeightResponse = _reflection.GeneratedProtocolMessageType('GetHeightResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETHEIGHTRESPONSE,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.GetHeightResponse)
  ))
_sym_db.RegisterMessage(GetHeightResponse)

BooleanResponse = _reflection.GeneratedProtocolMessageType('BooleanResponse', (_message.Message,), dict(
  DESCRIPTOR = _BOOLEANRESPONSE,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.BooleanResponse)
  ))
_sym_db.RegisterMessage(BooleanResponse)

VerifyResponse = _reflection.GeneratedProtocolMessageType('VerifyResponse', (_message.Message,), dict(
  DESCRIPTOR = _VERIFYRESPONSE,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.VerifyResponse)
  ))
_sym_db.RegisterMessage(VerifyResponse)

Null = _reflection.GeneratedProtocolMessageType('Null', (_message.Message,), dict(
  DESCRIPTOR = _NULL,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.Null)
  ))
_sym_db.RegisterMessage(Null)

Transaction = _reflection.GeneratedProtocolMessageType('Transaction', (_message.Message,), dict(
  DESCRIPTOR = _TRANSACTION,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.Transaction)
  ))
_sym_db.RegisterMessage(Transaction)

JsonBlockString = _reflection.GeneratedProtocolMessageType('JsonBlockString', (_message.Message,), dict(
  DESCRIPTOR = _JSONBLOCKSTRING,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.JsonBlockString)
  ))
_sym_db.RegisterMessage(JsonBlockString)

Block = _reflection.GeneratedProtocolMessageType('Block', (_message.Message,), dict(
  DESCRIPTOR = _BLOCK,
  __module__ = 'db_pb2'
  # @@protoc_insertion_point(class_scope:miniblockchaindb.Block)
  ))
_sym_db.RegisterMessage(Block)



_BLOCKCHAINMINER = _descriptor.ServiceDescriptor(
  name='BlockChainMiner',
  full_name='miniblockchaindb.BlockChainMiner',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=702,
  serialized_end=1200,
  methods=[
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='miniblockchaindb.BlockChainMiner.Get',
    index=0,
    containing_service=None,
    input_type=_GETREQUEST,
    output_type=_GETRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Transfer',
    full_name='miniblockchaindb.BlockChainMiner.Transfer',
    index=1,
    containing_service=None,
    input_type=_TRANSACTION,
    output_type=_BOOLEANRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Verify',
    full_name='miniblockchaindb.BlockChainMiner.Verify',
    index=2,
    containing_service=None,
    input_type=_TRANSACTION,
    output_type=_VERIFYRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetHeight',
    full_name='miniblockchaindb.BlockChainMiner.GetHeight',
    index=3,
    containing_service=None,
    input_type=_NULL,
    output_type=_GETHEIGHTRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetBlock',
    full_name='miniblockchaindb.BlockChainMiner.GetBlock',
    index=4,
    containing_service=None,
    input_type=_GETBLOCKREQUEST,
    output_type=_JSONBLOCKSTRING,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PushBlock',
    full_name='miniblockchaindb.BlockChainMiner.PushBlock',
    index=5,
    containing_service=None,
    input_type=_JSONBLOCKSTRING,
    output_type=_NULL,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PushTransaction',
    full_name='miniblockchaindb.BlockChainMiner.PushTransaction',
    index=6,
    containing_service=None,
    input_type=_TRANSACTION,
    output_type=_NULL,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_BLOCKCHAINMINER)

DESCRIPTOR.services_by_name['BlockChainMiner'] = _BLOCKCHAINMINER

# @@protoc_insertion_point(module_scope)
