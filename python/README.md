# Generated SPOE Parser for Python

This language is still a work in progress. I personally only use the golang version, so this version is currently untested and only provided as an example.

## Usage

```python

from parser.spop import Spop

frame = Spop.from_file("../resources/ack_frame.bin")
print(frame.len_frame)
```
