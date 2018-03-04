#/ Configure ssh keys, config and known_hosts

__devenv_plugin__ssh__setup() {
  local profile_folder
  profile_folder=$1
  mkdir -p "$profile_folder/ssh"
  [ -e "$profile_folder/ssh/config" ] || touch "$profile_folder/ssh/config"
  [ -e "$profile_folder/ssh/known_hosts" ] || touch "$profile_folder/ssh/known_hosts"
}

__devenv_plugin__ssh__configure() {
  __devenv_ui_arrow "setup ssh"
  if [ ! -f "$profile_folder/ssh/id_rsa" ]; then
    if __devenv_ui_prompt "Generate new ssh keypair?"; then
      read -p "Add comment: " comment
      if [[ $comment != "" ]]; then
        comment="$email-$profile_name-$comment"
      else
        comment="$email-$profile_name"
      fi
      ssh-keygen -b 4096 -t rsa -C "$comment" -f "$profile_folder/ssh/id_rsa"
    fi
  fi
  __devenv_ui_ok
}

__devenv_plugin__ssh__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  local folder
  folder="$profile_folder/ssh"
  local ssh_agent_cache
  ssh_agent_cache="/tmp/$profile-ssh-agent.tmp"

  echo "if [ ! -f $ssh_agent_cache ]; then"
    echo "echo \$(ssh-agent -s | sed \"s/echo/# echo/\") > $ssh_agent_cache"
    echo "chown $USER:$USER $ssh_agent_cache; chmod 700 $ssh_agent_cache"
  echo "fi"
  echo "source $ssh_agent_cache"

  if [ -d "$folder" ]; then
    for file in $folder/*.pub; do
      if [ -e "$file" ]; then
        echo "ssh-add -l 2> /dev/null | grep ${file%.*} > /dev/null"
        echo "[ \$? -gt 0 ] && ssh-add ${file%.*} > /dev/null"
      fi
    done
    for file in $folder/*.pem; do
      if [ -e "$file" ]; then
        echo "ssh-add -l 2> /dev/null | grep $file > /dev/null"
        echo "[ \$? -gt 0 ] && ssh-add $file > /dev/null"
      fi
    done

    echo -n "/usr/bin/ssh " > "$profile_folder/bin/ssh"
    [ -e "$folder/known_hosts" ] && echo -n "-o UserKnownHostsFile=$folder/known_hosts " >> "$profile_folder/bin/ssh"
    [ -e "$folder/config" ] && echo -n "-F $folder/config " >> "$profile_folder/bin/ssh"
    echo "\$@" >> "$profile_folder/bin/ssh"
    chmod +x "$profile_folder/bin/ssh"

    echo -n "/usr/bin/scp " > "$profile_folder/bin/scp"
    [ -e "$folder/config" ] && echo -n "-F $folder/config " >> "$profile_folder/bin/scp"
    echo "\$@" >> "$profile_folder/bin/scp"
    chmod +x "$profile_folder/bin/scp"
  fi

  return 0
}
