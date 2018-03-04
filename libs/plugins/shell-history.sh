
__devenv_plugin__shellhistory__generate_loader() {
  local profile_folder
  profile_folder=$1
  local profile
  profile=$2
  echo "export HISTFILE='$profile_folder/$(basename $SHELL)-history'"
}
