#!/bin/bash
set -e

# See: https://github.com/scop/bash-completion/blob/master/README.md

USER_DIR="$BASH_COMPLETION_USER_DIR"

if [ -z "$USER_DIR" ]; then
	DATA_HOME="$XDG_DATA_HOME"
	if [ -z "$DATA_HOME" ]; then
		DATA_HOME=~/.local/share
	fi 
	USER_DIR="$DATA_HOME/bash-completion"
fi

mkdir --parents "$USER_DIR/completions"
puccini-tosca --bash-completion "$USER_DIR/completions/puccini-tosca"
puccini-js --bash-completion "$USER_DIR/completions/puccini-js"
