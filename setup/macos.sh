#!/bin/bash

if command -v brew &> /dev/null
then
    read -p "Regit depends on git-filter-repo, so it will be installed through brew. Are you okay with it? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]
    then
        brew install git-filter-repo
    fi
    else
        echo "Please install Homebrew first."
        exit 1
    fi
fi

latest_release=$(curl --silent "https://api.github.com/repos/woaitsAryan/regit/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

curl -LO "https://github.com/woaitsAryan/regit/releases/download/${latest_release}/regit_MacOS_all.tar.gz"

tar -xvzf regit_MacOS_all.tar.gz

mv regit /usr/local/bin/

rm regit_MacOS_all.tar.gz
rm LICENSE
rm README.md