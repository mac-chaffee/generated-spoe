# generated-spoe

This repo contains a parser for the Stream Processing Offload Protocol (SPOP), generated via kaitai.

Official protocol description: https://github.com/haproxy/haproxy/blob/v2.7.0/doc/SPOE.txt

## Generation

1. Follow these instructions to install kaitai: https://kaitai.io/#quick-start
2. In the root of this repo, run `go generate ./...`

Any changes to `grammar/*.ksy` requires regeneration. Otherwise, do not edit the generated code in the `parser` folder!

## Testing

```
go test ./...
```
