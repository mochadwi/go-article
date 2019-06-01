#!/usr/bin/env bash

pg_ctl -D /path/to/postgres start
dep ensure
make
gin -p 9091 -a 9090 article_clean