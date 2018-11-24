#!/bin/sh

set -e

repo=$1
dest=$2
cred=$3
# remote name
remote=origin
# branch name
branch=master

cd ${dest}
if [[ ! -d ".git" ]]; then 
    git init
    git remote add ${remote} "https://${cred}@github.com/${repo}"
fi
git pull ${remote} ${branch}

