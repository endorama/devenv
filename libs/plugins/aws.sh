
__devenv_plugin__aws__setup() {
  local profile_folder
  profile_folder=$1
  mkdir -p "$profile_folder/aws"
  [ -e "$profile_folder/aws/config" ] || touch "$profile_folder/aws/config"
  [ -e "$profile_folder/aws/credentials" ] || touch "$profile_folder/aws/credentials"
}


__devenv_plugin__aws__configure() {
  __devenv_ui_arrow "setup aws"
  __devenv_ui_ok
}

__devenv_plugin__aws__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  echo "export AWS_CONFIG_FILE=$profile_folder/aws/config"
  echo "export AWS_SHARED_CREDENTIALS_FILE=$profile_folder/aws/credentials"
}
