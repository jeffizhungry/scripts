#!/bin/bash
VERSION='1.7'
   
function main {
   if [[ `uname` != 'Linux' && `uname` != 'Darwin' ]]; then
      echo "Error: This install script is only for linux/darwin systems"
      exit 1
   fi
   if [[ `whoami` != 'root' ]]; then
      echo "Error: Please run as root"
      exit 1
   fi

   OPSYS="`uname` | awk '{print tolower($0)}'"

   # Get correct Go distribution
   GO_DIST=""
   if [[ `uname -m` -eq 'x86_64' ]]; then
      GO_DIST="go$VERSION.linux-amd64.tar.gz"
   elif [[ `uname -m` -eq 'i686' ]]; then
      GO_DIST="go$VERSION.linux-386.tar.gz"
   else
      echo "Error: This script does not handle " `uname -m`
   fi

   GO_INSTALL_URL=https://storage.googleapis.com/golang/$GO_DIST
   GO_INSTALL_TMP=/tmp/$GO_DIST
   GO_INSTALL_DIR=/usr/local/go

   echo "---------------------------"
   echo "| Go Linux Install Script |"
   echo "---------------------------"

   # Downlaod go distribution from golang website
   echo "-- Downloading Go Distribution..." $GO_DIST
   curl -so $GO_INSTALL_TMP $GO_INSTALL_URL

   # Install go distribution
   echo "-- Installing go under.... $GO_INSTALL_DIR"
   tar xf $GO_INSTALL_TMP -C '/usr/local' 

   # Create symlink to tool chain
   rm -r '/usr/local/bin/go'
   echo "-- Creating symbolic link for go tool chain"
   ln -sf "$GO_INSTALL_DIR/bin/go" '/usr/local/bin/go'
   echo "-- Creating symbolic link for gofmt"
   ln -sf "$GO_INSTALL_DIR/bin/gofmt" '/usr/local/bin/gofmt'
   echo "-- Creating symbolic link for godoc"
   ln -sf "$GO_INSTALL_DIR/bin/godoc" '/usr/local/bin/godoc'

   # Remove tar
   rm -rf $GO_INSTALL_TMP
}

main
