
[ -e $_DEVENV_CONFIG/devenv.cfg ] && source $_DEVENV_CONFIG/devenv.cfg

function get_config() {
  local var
  var="$1"
  echo "${!var}"
}