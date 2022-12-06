#!/bin/sh
set -e
kaitai-struct-compiler --target=python ../grammar/spop.ksy --python-package=parser --verbose=file --outdir=parser
