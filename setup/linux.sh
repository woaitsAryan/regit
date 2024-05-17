#!/bin/bash

latest_release=$(curl --silent "https://api.github.com/repos/woaitsAryan/regit/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

curl -LO "https://github.com/woaitsAryan/regit/releases/download/${latest_release}/regit_MacOS_all.tar.gz"

tar -xvzf regit_MacOS_all.tar.gz

mv regit /usr/local/bin/

rm regit_MacOS_all.tar.gz
rm LICENSE
rm README.md