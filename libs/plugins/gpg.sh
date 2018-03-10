#/ Configure gpg cli

__devenv_plugin__gpg__setup() {
  local profile_folder
  profile_folder=$1
  mkdir -p "$profile_folder/gpg"

  chmod 700 "$profile_folder/gpg"
}

__devenv_plugin__gpg__configure() {
  local profile_config_folder
  profile_config_folder=$1
  __devenv_ui_arrow "configure gpg executable"

  echo "Choose gpg executable:"
  select yn in "gpg" "gpg2"; do
    case $yn in
      gpg) gpg_exec="$(which gpg)"; break;;
      gpg2) gpg_exec="$(which gpg2)"; break;;
    esac
  done

  echo $gpg_exec > $profile_config_folder/__devenv_plugin__gpg__executable
  __devenv_ui_ok
}

__devenv_plugin__gpg__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2

  local folder
  folder="$profile_folder/gpg"

  gpg_exec=$(cat $profile_folder/.config/__devenv_plugin__gpg__executable)
  echo $gpg_exec
  echo "export GNUPGHOME=${folder}" > "$profile_folder/bin/gpg"
  echo -n "${gpg_exec} " >> "$profile_folder/bin/gpg"
  echo "\$@" >> "$profile_folder/bin/gpg"
  chmod +x "$profile_folder/bin/gpg"

  chmod 600 -R "${folder}/*" 2>/dev/null

  return 0
}
