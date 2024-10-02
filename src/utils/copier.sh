#!/bin/bash
rm -rf src/models/ent/schema
rsync -av --exclude='*.resources.go' src/models/schema src/models/ent/