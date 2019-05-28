#!/usr/bin/env bash

dep ensure
make
gin -a 3000 article_clean # auto-reload