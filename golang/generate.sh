#!/bin/sh
set -e
kaitai-struct-compiler --target=go ../grammar/spop.ksy --go-package=parser --verbose=file
