# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO
from enum import Enum


if getattr(kaitaistruct, 'API_VERSION', (0, 9)) < (0, 9):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

from parser import varint
class Spop(KaitaiStruct):
    """
    .. seealso::
       Source - https://github.com/haproxy/haproxy/blob/v2.7.0/doc/SPOE.txt
    """

    class FrameType(Enum):
        unset = 0
        haproxy_hello = 1
        haproxy_disconnect = 2
        notify = 3
        agent_hello = 101
        agent_disconnect = 102
        ack = 103

    class ActionType(Enum):
        set_var = 1
        unset_var = 2

    class VarScope(Enum):
        process = 0
        session = 1
        transaction = 2
        request = 3
        response = 4
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.len_frame = self._io.read_u4be()
        self._raw_frame = self._io.read_bytes(self.len_frame)
        _io__raw_frame = KaitaiStream(BytesIO(self._raw_frame))
        self.frame = Spop.Frame(_io__raw_frame, self, self._root)

    class FragmentedFrame(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.raw_data = self._io.read_bytes_full()


    class Ipv4(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.addr = self._io.read_bytes(4)


    class SpopString(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.str_len = varint.Varint(self._io)
            self.str_data = (self._io.read_bytes(self.str_len.value)).decode(u"ASCII")


    class ActionSetVar(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.var_scope = KaitaiStream.resolve_enum(Spop.VarScope, self._io.read_u1())
            self.var_name = Spop.SpopString(self._io, self, self._root)
            self.var_value = Spop.TypedData(self._io, self, self._root)


    class MetaFlags(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.reserved_flags = self._io.read_bits_int_be(30)
            self.abort_flag = self._io.read_bits_int_be(1) != 0
            self.fin_flag = self._io.read_bits_int_be(1) != 0


    class Frame(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.frame_type = KaitaiStream.resolve_enum(Spop.FrameType, self._io.read_u1())
            self.frame_meta = Spop.FrameMeta(self._io, self, self._root)
            _on = self.frame_type
            if _on == Spop.FrameType.notify:
                self.frame_payload = Spop.ListOfMessages(self._io, self, self._root)
            elif _on == Spop.FrameType.agent_hello:
                self.frame_payload = Spop.KvList(self._io, self, self._root)
            elif _on == Spop.FrameType.unset:
                self.frame_payload = Spop.FragmentedFrame(self._io, self, self._root)
            elif _on == Spop.FrameType.agent_disconnect:
                self.frame_payload = Spop.KvList(self._io, self, self._root)
            elif _on == Spop.FrameType.ack:
                self.frame_payload = Spop.ListOfActions(self._io, self, self._root)
            elif _on == Spop.FrameType.haproxy_disconnect:
                self.frame_payload = Spop.KvList(self._io, self, self._root)
            elif _on == Spop.FrameType.haproxy_hello:
                self.frame_payload = Spop.KvList(self._io, self, self._root)


    class TypedData(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.type_flags = self._io.read_bits_int_be(4)
            self.type = self._io.read_bits_int_be(4)
            self._io.align_to_byte()
            _on = self.type
            if _on == 0:
                self.type_data = Spop.SpopBool(self.type_flags, self._io, self, self._root)
            elif _on == 4:
                self.type_data = varint.Varint(self._io)
            elif _on == 6:
                self.type_data = Spop.Ipv4(self._io, self, self._root)
            elif _on == 7:
                self.type_data = Spop.Ipv6(self._io, self, self._root)
            elif _on == 1:
                self.type_data = Spop.NullType(self._io, self, self._root)
            elif _on == 3:
                self.type_data = varint.Varint(self._io)
            elif _on == 5:
                self.type_data = varint.Varint(self._io)
            elif _on == 8:
                self.type_data = Spop.SpopString(self._io, self, self._root)
            elif _on == 9:
                self.type_data = Spop.SpopString(self._io, self, self._root)
            elif _on == 2:
                self.type_data = varint.Varint(self._io)


    class ListOfMessages(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.message_name = Spop.SpopString(self._io, self, self._root)
            self.nb_args = self._io.read_u1()
            self.kv_list = []
            for i in range(self.nb_args):
                self.kv_list.append(Spop.KvList(self._io, self, self._root))



    class ActionUnsetVar(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.var_scope = KaitaiStream.resolve_enum(Spop.VarScope, self._io.read_u1())
            self.var_name = Spop.SpopString(self._io, self, self._root)


    class FrameMeta(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self._raw_meta_flags = self._io.read_bytes(4)
            _io__raw_meta_flags = KaitaiStream(BytesIO(self._raw_meta_flags))
            self.meta_flags = Spop.MetaFlags(_io__raw_meta_flags, self, self._root)
            self.stream_id = varint.Varint(self._io)
            self.frame_id = varint.Varint(self._io)


    class Ipv6(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.addr = self._io.read_bytes(16)


    class KvList(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.kv_name = Spop.SpopString(self._io, self, self._root)
            self.kv_value = Spop.TypedData(self._io, self, self._root)


    class SpopBool(KaitaiStruct):
        """bools pull their values from the type_flags."""
        def __init__(self, type_flags, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self.type_flags = type_flags
            self._read()

        def _read(self):
            pass

        @property
        def value(self):
            if hasattr(self, '_m_value'):
                return self._m_value

            self._m_value = self.type_flags
            return getattr(self, '_m_value', None)


    class NullType(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            pass


    class ListOfActions(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.action_type = KaitaiStream.resolve_enum(Spop.ActionType, self._io.read_u1())
            self.nb_args = self._io.read_u1()
            _on = self.action_type
            if _on == Spop.ActionType.set_var:
                self.action_args = Spop.ActionSetVar(self._io, self, self._root)
            elif _on == Spop.ActionType.unset_var:
                self.action_args = Spop.ActionUnsetVar(self._io, self, self._root)



