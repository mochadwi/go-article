#!/usr/bin/env bash

dep ensure
make
gin -a 3333 article_clean # auto-reload