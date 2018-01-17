
__devenv_plugin__bin__setup() {
  local profile_folder
  profile_folder=$1
  [ -d "$profile_folder/bin" ] || mkdir -p "$profile_folder/bin"
}

__devenv_plugin__bin__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  echo "export PATH=$profile_folder/bin:\$PATH"
}
