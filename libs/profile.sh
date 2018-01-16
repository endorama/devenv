#!/bin/bash

set -euo pipefail

PROFILE_PATHS="${_DEVENV_ROOT}/profiles"

source "$_DEVENV_ROOT/libs/plugins/bin.sh"
source "$_DEVENV_ROOT/libs/plugins/env.sh"
source "$_DEVENV_ROOT/libs/plugins/ssh.sh"
source "$_DEVENV_ROOT/libs/plugins/aws.sh"

function get_active_profile() {
  DEVENV_ACTIVE_PROFILE=${DEVENV_ACTIVE_PROFILE:-}
  [[ "$DEVENV_ACTIVE_PROFILE" ]] && echo "$DEVENV_ACTIVE_PROFILE"
  echo ""
}

profile_exists() {
  local profile_name
  profile_name=$1
  if [ ! -d "$PROFILE_PATHS/$profile_name" ]; then
    return 1
  else
    return 0
  fi
}

create_profile() {
  local profile_name
  profile_name=$1
  local profile_folder
  profile_folder=$2
  profile_exists "$profile_name" || mkdir "$profile_folder"
  return 0
}

profile_export_path() {
  local profile_folder
  profile_folder=$1
  echo "export PATH=$profile_folder/bin:\$PATH"
}

profile_generate_loader() {
  local profile
  profile=$1
  local profile_folder
  profile_folder="$(get_config "profiles")/$profile"

  echo "#!$SHELL"
  echo "#"
  echo ""
  echo "export DEVENV_ACTIVE_PROFILE='$profile'"
  echo "export DEVENV_ACTIVE_PROFILE_PATH='$profile_folder'"
  profile_prepare_bin_folder "$profile_folder"
  echo "export HISTFILE='$profile_folder/zsh-history'"
  profile_load_envs "$profile_folder"
  profile_load_ssh "$profile_folder" "$profile"
  profile_load_aws "$profile_folder" "$profile"
  profile_export_path "$profile_folder"
}
