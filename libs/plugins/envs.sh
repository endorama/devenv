
__devenv_plugin__envs__setup() {
  local profile_folder
  profile_folder=$1
  [ -e "$profile_folder/envs" ] || touch "$profile_folder/envs"
}

__devenv_plugin__envs__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  local file
  file="$profile_folder/envs"
  # local env

  if [ -e "$file" ]; then
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
