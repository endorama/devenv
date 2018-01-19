
__devenv_plugin__email__configure() {
  local profile_config_folder
  profile_config_folder=$1
  __devenv_ui_arrow "configure email"
  read -p "Email address: " email
  echo $email > $profile_config_folder/__devenv_plugin__email__config
  __devenv_ui_ok
}

__devenv_plugin__email__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  if [ -e $profile_folder/.config/__devenv_plugin__email__config ]; then
    email=$(cat $profile_folder/.config/__devenv_plugin__email__config)
    echo "export EMAIL=$email"
  fi
}
