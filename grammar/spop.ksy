meta:
  id: spop
  file-extension: spop
  endian: be
  bit-endian: be
  imports:
    - varint
doc-ref: https://github.com/haproxy/haproxy/blob/v2.7.0/doc/SPOE.txt
seq:
  - id: len_frame
    type: u4
  - id: frame
    type: frame
    size: len_frame
types:
  frame:
    seq:
      - id: frame_type
        type: u1
        enum: frame_type
      - id: frame_meta
        type: frame_meta
      - id: frame_payload
        type:
          switch-on: frame_type
          cases:
            'frame_type::unset': fragmented_frame
            'frame_type::haproxy_hello': kv_list
            'frame_type::haproxy_disconnect': kv_list
            'frame_type::notify': list_of_messages
            'frame_type::agent_hello': kv_list
            'frame_type::agent_disconnect': kv_list
            'frame_type::ack': list_of_actions
  fragmented_frame:
    seq:
      # When encountering a raw frame, calling code needs to
      # concat all the raw data and re-parse
      - id: raw_data
        size-eos: true
  frame_meta:
    seq:
      - id: meta_flags
        type: meta_flags
        size: 4
      - id: stream_id
        type: varint
      - id: frame_id
        type: varint
  meta_flags:
    seq:
      - id: reserved_flags
        type: b30
      - id: abort_flag
        type: b1
      - id: fin_flag
        type: b1
  list_of_actions:
    seq:
      - id: actions
        type: action
        repeat: eos
  action:
    seq:
      - id: action_type
        type: u1
        enum: action_type
      - id: nb_args
        type: u1
      - id: action_args
        # Note: SPOE.txt has a mistake; this is an action, not typed data
        type:
          switch-on: action_type
          cases:
            'action_type::set_var': action_set_var
            'action_type::unset_var': action_unset_var
  list_of_messages:
    seq:
      - id: messages
        type: message
        repeat: eos
  message:
    seq:
      - id: message_name
        type: spop_string
      - id: nb_args
        type: u1
      - id: kvs
        type: kv
        repeat: expr
        repeat-expr: nb_args
  kv_list:
    seq:
      - id: kvs
        type: kv
        repeat: eos
  kv:
    seq:
      - id: kv_name
        type: spop_string
      - id: kv_value
        type: typed_data
  action_set_var:
    seq:
      - id: var_scope
        type: u1
        enum: var_scope
      - id: var_name
        type: spop_string
      - id: var_value
        type: typed_data
  action_unset_var:
    seq:
      - id: var_scope
        type: u1
        enum: var_scope
      - id: var_name
        type: spop_string
  typed_data:
    seq:
      # TODO: Figure out why we have to swap the order here
      - id: type_flags
        type: b4
      - id: type
        type: b4
      - id: type_data
        type:
          switch-on: type
          cases:
            0: spop_bool(type_flags)
            1: null_type
            2: varint
            3: varint
            4: varint
            5: varint
            6: ipv4
            7: ipv6
            8: spop_string
            # Technically binary, but same type under the hood
            9: spop_string
  spop_string:
    seq:
      - id: str_len
        type: varint
      - id: str_data
        type: str
        encoding: ASCII
        size: str_len.value
  spop_bool:
    params:
      - id: type_flags
        type: b4
    instances:
      value:
        value: type_flags
    doc: "bools pull their values from the type_flags"
  null_type: {}
  ipv6:
    seq:
      - id: addr
        size: 16
  ipv4:
    seq:
      - id: addr
        size: 4
enums:
  frame_type:
    0: unset
    1: haproxy_hello
    2: haproxy_disconnect
    3: notify
    101: agent_hello
    102: agent_disconnect
    103: ack
  action_type:
    1: set_var
    2: unset_var
  var_scope:
    0: process
    1: session
    2: transaction
    3: request
    4: response
