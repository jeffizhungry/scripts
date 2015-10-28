#!/usr/bin/bash
# *****************************************************************************
# Setup Go Workspace 
# *****************************************************************************
GOPATH=$HOME/go
mkdir -p $GOPATH

# Create bin, pkg, src if they don't already exist
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg
mkdir -p $GOPATH/src

# Export go workspace varaibles
export GOPATH=$GOPATH
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

# Check if go cmd exists

# Install godeps, cuz everyone should have godeps

# Install other useful go packages

