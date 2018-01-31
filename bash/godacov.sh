#!/bin/bash

godacov -t $CODACY_TOKEN -r ./coverage.txt -c $TRAVIS_COMMIT