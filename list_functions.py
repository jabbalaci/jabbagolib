#!/usr/bin/env python3

"""
I tend to forget what is where...
This script prints the package name and
shows the available functions inside that package.

requirements:

$ pip install yachalk
"""

import os
import re
from glob import glob

from yachalk import chalk

IGNORE_DIRS = (".git", ".old", "examples")
DOT = "•"

TL = "┏"  # top left
TR = "┓"  # top right
HOR = "━"  # horizontal
VER = "┃"  # vertical
BL = "┗"  # bottom left
BR = "┛"  # bottom right


def colorize(sign):
    result = sign
    m = re.search(r"func ([^([]+)", sign)
    if m:
        fn_name = m.group(1)
        new_value = chalk.bold.blue(fn_name)
        result = result.replace(fn_name, new_value)
    #
    return result


def render_fn_signature(fn_signatures):
    for sign in fn_signatures:
        sign = colorize(sign)
        print(f"    {DOT} {sign}")


def collect_fn_signatures(fname):
    result = []
    with open(fname) as f:
        for line in f:
            m = re.search(r"func [A-Z]+", line)
            if m:
                result.append(line.rstrip(" {\n"))
            #
        #
    #
    return result


def process_files_inside(dname):
    files = sorted(glob(f"{dname}/*.go"))
    for fname in files:
        if fname.endswith("_test.go"):
            continue
        #
        fn_signatures = collect_fn_signatures(fname)
        render_fn_signature(fn_signatures)


def render_dirname(dname):
    # print(chalk.bold.green.bg_gray(dname))
    text = chalk.bold(dname)
    print("{0}{1}{2}".format(TL, HOR * len(dname), TR))
    print("{0}{1}{2}".format(VER, text, VER))
    print("{0}{1}{2}".format(BL, HOR * len(dname), BR))


def process_dir(dname):
    render_dirname(dname)
    process_files_inside(dname)


def main():
    entries = sorted(os.listdir("."))
    for e in entries:
        if os.path.isdir(e):
            if e not in IGNORE_DIRS:
                process_dir(e)
            #
        #
    #


##############################################################################

if __name__ == "__main__":
    main()
