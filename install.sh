#!/usr/bin/env bash

dep ensure
make
gin -a 9090 article_clean # auto-reload