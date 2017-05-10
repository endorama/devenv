#!/usr/bin/env bash

# Logging stuff.
e_header()   { echo -e "\n\033[1m$*\033[0m"; }
e_success()  { echo -e "  \033[1;32m✔\033[0m  $*"; }
e_error()    { echo -e "  \033[1;31m✖\033[0m  $*"; }
e_warn()     { echo -e "  \033[1;33m  $*\033[0m"; }
e_arrow()    { echo -e "  \033[1;33m➜\033[0m  $*"; }
e_message()  { echo -e "$@"; }

e_ok()       { e_success "ok"; }
# $1 => error message, $2 => exit code
e_abort()    { e_error $1; exit $2; }
# $1 => warn message
e_return()    { e_warn $1; }
