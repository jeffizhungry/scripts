#!/bin/bash
# *****************************************************************************
# Setup Go Workspace with this scipt's path as GOPATH
#  - Note: GOPATH is where all non-standard libraries are stored
# *****************************************************************************

# Get the path of the directory containing this script
CURR_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Create bin, pkg, src if they don't already exist
if [ ! -d $CURR_PATH/bin ]; then
    mkdir bin
fi
if [ ! -d $CURR_PATH/pkg ]; then
    mkdir pkg
fi
if [ ! -d $CURR_PATH/src ]; then
    mkdir src
fi

# Export go workspace varaibles
export GOPATH=$CURR_PATH
export GOBIN=$CURR_PATH/bin
export PATH=$PATH:$GOBIN
