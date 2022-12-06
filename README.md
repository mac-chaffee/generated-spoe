# generated-spoe

This repo contains a parser for the Stream Processing Offload Protocol (SPOP), generated via kaitai.

Official protocol description: https://github.com/haproxy/haproxy/blob/v2.7.0/doc/SPOE.txt

See the README in each language-specific subfolder for more info.

## Folder Structure

* `golang/`: Contains the generated parser and tests for Go.
* `grammar/`: Contains the Kaitai files that formally define the SPOProtocol.
* `python/`: Contains the generated parser and tests for Python.
* `resources/`: Contains SPOP example files for use in writing tests.

## Contributing

The best way you can contribute at the moment is by providing more SPOP binary files for testing.

You can obtain an SPOP binary file like so:

1. Use `tcpdump` to capture SPOP TCP packets between an HAProxy instance and an existing SPOAgent
2. Load the packet capture into WireShark
3. Select one of the SPOP packets
4. In the bottom box, find the "Data" section of the packet. Right click > Copy > ...as a Hex Stream
    - This is the raw SPOP data that this repo parses.
5. In your terminal, use `xxd` to convert the hex stream into binary and save it to a file:

```
echo -n "paste hex stream here" | xxd -r -ps - my_spop_frame.bin
```
