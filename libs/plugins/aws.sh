
profile_prepare_aws_folder() {
  local profile_folder
  profile_folder=$1
  mkdir -p "$profile_folder/aws"
  [ -e "$profile_folder/aws/config" ] || touch "$profile_folder/aws/config"
  [ -e "$profile_folder/aws/credentials" ] || touch "$profile_folder/aws/credentials"
}


setup_aws() {
  __devenv_ui_arrow "setup aws"
  profile_prepare_aws_folder "$profile_folder"

  __devenv_ui_ok
}

profile_load_aws() {
  local profile_folder
  profile_folder=$1
  echo "export AWS_CONFIG_FILE=$profile_folder/aws/config"
  echo "export AWS_SHARED_CREDENTIALS_FILE=$profile_folder/aws/credentials"
}
