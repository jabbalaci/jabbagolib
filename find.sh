#!/usr/bin/env bash

if [ -z "$1" ]
  then
    python3 list_functions.py
    echo
    echo "Warning: you didn't provide a search term"
    exit
fi

python3 list_functions.py | ack --passthru -i "$1"
echo
echo "with grep:"
python3 list_functions.py | grep --color -i "$1"
