
__devenv_plugin__email__configure() {
  __devenv_ui_arrow "configure email"
  read -p "Email address: " email
  echo $email > /tmp/__devenv_plugin__email__config
  __devenv_ui_ok
}

__devenv_plugin__email__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  email=$(cat /tmp/__devenv_plugin__email__config)
  echo "export EMAIL=$email"
  rm /tmp/__devenv_plugin__email__config
}
