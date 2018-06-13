#!/bin/bash

profiles="$HOME/"

[ -e "$_DEVENV_CONFIG/devenv.cfg" ] && source "$_DEVENV_CONFIG/devenv.cfg"

__devenv_config_get() {
  local var
  var="$1"
  echo "${!var}"
}
