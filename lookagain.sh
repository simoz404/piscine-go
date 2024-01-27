#! /bin/bash
find . -name '*.sh' | sed 's/\.sh$//' | rev | cut -d '/' -f 1 | rev | sort -r
