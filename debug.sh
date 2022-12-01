#!/usr/bin/env bash
make build-debug
dlv exec --accept-multiclient --continue --headless ./terraform-provider-example -- -debug