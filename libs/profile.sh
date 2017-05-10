#!/bin/bash

set -euo pipefail

PROFILE_PATHS="${_DEVENV_ROOT}/profiles"

function get_active_profile() {
  DEVENV_ACTIVE_PROFILE=${DEVENV_ACTIVE_PROFILE:-}
  [[ $DEVENV_ACTIVE_PROFILE ]] && echo $DEVENV_ACTIVE_PROFILE
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

profile_prepare_bin_folder() {
  local profile_folder
  profile_folder=$1
  mkdir -p $profile_folder/bin
}

profile_load_envs() {
  local profile_folder
  profile_folder=$1
  local file
  file="$profile_folder/envs"
  # local env

  if [ -e $file ]; then
    while read -r line; do  
      # discard empty lines or spaces only lines and lines starting with #
      if [[ "$line" =~ [^[:space:]] && ! "$line" =~ \#.* ]]; then
        # get env from line
        # env=$(echo $line | cut -d "=" -f 1)
        # export if not already present
        # if [[ $(printenv ${env}) == "" ]]; then
          echo "export ${line}"
        # fi
      fi
    done < "$file"
  fi
  return 0
}

profile_load_ssh() {
  local profile_folder
  profile_folder=$1
  local folder
  folder="$profile_folder/ssh"

  if [ -d $folder ]; then
    if [ $SSH_AGENT_PID != "" ]; then
      [ -e "$folder/id_rsa" ] && echo "ssh-add $folder/id_rsa"
    fi

    echo -n "/usr/bin/ssh " > $profile_folder/bin/ssh
    [ -e "$folder/id_rsa" ] && echo -n "-i $folder/id_rsa " >> $profile_folder/bin/ssh
    [ -e "$folder/known_hosts" ] && echo -n "-o UserKnownHostsFile=$folder/known_hosts " >> $profile_folder/bin/ssh
    [ -e "$folder/config" ] && echo -n "-F $folder/config " >> $profile_folder/bin/ssh
    echo "\$@" >> $profile_folder/bin/ssh
  fi

  

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


  echo "#!/bin/bash" 
  echo "#" 
  echo "" 
  echo "export DEVENV_ACTIVE_PROFILE=$profile" 

  profile_prepare_bin_folder $profile_folder
  profile_load_envs $profile_folder
  profile_load_ssh $profile_folder
  profile_export_path $profile_folder

  echo "$SHELL -l" 
}