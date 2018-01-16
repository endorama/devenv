#!/usr/bin/env bash

# Logging stuff.
__devenv_ui_header()   { echo -e "\n\033[1m$*\033[0m"; }
__devenv_ui_success()  { echo -e "  \033[1;32m✔\033[0m  $*"; }
__devenv_ui_error()    { echo -e "  \033[1;31m✖\033[0m  $*"; }
__devenv_ui_warn()     { echo -e "  \033[1;33m  $*\033[0m"; }
__devenv_ui_arrow()    { echo -e "  \033[1;33m➜\033[0m  $*"; }
__devenv_ui_message()  { echo -e "$@"; }

__devenv_ui_ok()       { __devenv_ui_success "ok"; }
# $1 => error message, $2 => exit code
__devenv_ui_abort()    { __devenv_ui_error "$1"; exit "$2"; }
# $1 => warn message
__devenv_ui_return()    { __devenv_ui_warn "$1"; }

__devenv_ui_prompt() {
  echo "$1"
  select yn in "Yes" "No"; do
    case $yn in
      Yes) return 0; break;;
      No) return 1; break;;
    esac
  done
}
