# Downloaded from: https://formats.kaitai.io/vlq_base128_be/
# With modifications to support the differences in SPOP's varint
meta:
  id: varint
  title: Variable length quantity, unsigned integer, base128, big-endian
  license: CC0-1.0
  ks-version: 0.7
doc: |
  A variable-length unsigned integer using base128 encoding. 1-byte groups
  consist of 1-byte flag of continuation and 7-bit value chunk, and are ordered
  "most significant group first", i.e. in "big-endian" manner.

  This particular encoding is specified and used in:

  * Standard MIDI file format
  * ASN.1 BER encoding
  * RAR 5.0 file format

  More information on this encoding is available at
  https://en.wikipedia.org/wiki/Variable-length_quantity

  This particular implementation supports serialized values to up 8 bytes long.
-webide-representation: '{value:dec}'
seq:
  - id: groups
    type: group(_io.pos)
    repeat: until
    repeat-until: not _.has_next
types:
  group:
    -webide-representation: '{b}'
    doc: |
      One byte group, clearly divided into 7-bit "value" chunk and 1-bit "continuation" flag.
    params:
      - id: pos
        type: s8
    seq:
      - id: b
        type: u1
    instances:
      has_next:
        value: "pos == 0 ? b >= 240 : (b & 0b1000_0000) != 0"
        doc: If true, then we have more bytes to read
      # TODO: Can get rid of this if golang target supports .as<u1>
      value:
        value: "b + 0"
instances:
  value:
    value: >-
      groups[0].value
      + ( groups.size >= 2 ? (groups[1].value << 4) : 0 )
      + ( groups.size >= 3 ? (groups[2].value << 11) : 0 )
      + ( groups.size >= 4 ? (groups[3].value << 18) : 0 )
      + ( groups.size >= 5 ? (groups[4].value << 25) : 0 )
      + ( groups.size >= 6 ? (groups[5].value << 32) : 0 )
      + ( groups.size >= 7 ? (groups[6].value << 39) : 0 )
      + ( groups.size >= 8 ? (groups[7].value << 46) : 0 )
    doc: Resulting value as normal integer
