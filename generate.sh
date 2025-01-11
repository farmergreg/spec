#!/usr/bin/env sh

# Run this after updating the spec TSV files in src/pkg/spec

cd src/cmd/specgen
go build && ./specgen
