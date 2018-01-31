#!/bin/bash

# The .ssh directory probably exists, but why take the risk?
mkdir -p ~/.ssh
# This command clears all SSH keys from the current session,
# preventing the GitHub repository deploy key from being used
ssh-add -D
# This adds a directive to the SSH config which specifies our
# deploy user's key is used for authenticating with all connections
# to github.com
echo -e "Host github.com\n  User my-machine-user\n  IdentityFile ${TRAVIS_BUILD_DIR}/ci/ssh/insecure-key\n" >> ~/.ssh/config
# SSH keys will be unusable unless they are unreadable by
# anyone other than the owner
chmod 600 ${TRAVIS_BUILD_DIR}/ci/ssh/*
# We need to add the SSH signature for GitHub to the SSH known_hosts
cp ${TRAVIS_BUILD_DIR}/ci/ssh/known_hosts ~/.ssh/known_hosts
# Now update the submodules we're using
git submodule update --init --recursiv