# DevEnv

DevEnv is a collection of scripts to manage separate development environments,
relating to credentials, custom executable, environment variables ( and probably
stuff I didn't think of ).

Is heavily inspired by the beautiful cli interface of [anyenv][anyenv].

Is based on [37's Signals sub][sub] utility.

## Why

The problem: I started working as a freelance, and with consultancies comes clients.
For sane security, I don't want ( or is not possible ) to reuse the same credentials
for every client: ssh key, AWS credentials, GitHub/GitLab/Bitbucket tokens,
are the most common things changing between different clients.

As I hadn't find any way to properly manage this, I started building app a series
of helper bash scripts to load/unload the shell environment in order to provide
the specific tools I was needing on a particular client.

After more than 6 months of every day use, I think it's time for me to make this
public.

## What's inside

We are talking (mostly) about bash scripts here. As `devenv` is based upon the
extensible [sub][sub] utility, you can plug commands as you see fit, in whatever
language you prefer, you can override specific functionalities overriding the
executable ( loading your executable before in the PATH ).  
If you're not familiar with [sub][sub], check it out for more informations.

`devenv` is structured so that you can store all sensitive data out of it's folder
(avoiding risky git push). As long as the files are readable, you are ok: store
them on an encrypted partition mounted at boot/login, on a network drive, on a
separate encrypted USB, on your phone, choose yours as long as the files are
readable when you run the `devenv` command.

**Please note:** security of credential on disk/in memory is not in the scope of `devenv`.
You are responsible for securely storing these credentials, `devenv` is just a 
quick way to load a prefixed shell environment.  
Also note that `devenv` has not been designed for complex threat models: it's
assuming that your computer is not compromised (so storing cleartext credentials
is "ok" and there are no evil process trying to steal credentials from memory).
Any help in this direction would be awesome and really appreciated. 

## Installation

### Requirements

- make sure you have bash available (most of the time it is)
- make sure bash it at least at version 4.x

### Install

To install `devenv`, clone it (or download a tarball).

Then run `path/to/devenv/bin/devenv init -`.

Is generally advised to add a line to your shell load file (bash_profile, zshenv, ...)
to load the utility automatically, making it available at every shell:

```
eval "$(path/to/devenv/bin/devenv init -)"
```

## Configuration

`devenv` will search for a `devenv.cfg` file inside the `$_DEVENV_CONFIG` folder.

By default this variable has value of `$XDG_CONFIG_HOME/devenv` if available, or
`$HOME/.config/devenv` if not.

### Available configuration parameter

- `profiles`: a path to a folder containing profiles

## Profiles

`devenv` use the concept of *profiles* to separate different environments.

A profile is a folder containing all the required data to load the wanted environment:

- plugin's configurations
- `devenv` files to properly setup an environment: `load.sh` and `run.sh`

Profile name can be anything (as long as you can type it in the terminal), but
note that the name **shared** is reserved by `devenv`, so will be ignored.

## Plugins

To support different configurations/programs/whatever, `devenv` uses a modular
structure called *plugin* for brevity.

Currently available plugins are:

- `aws`: Configure AWS cli ( configs and credentials )
- `bin`: Add bin folder to PATH
- `email`: Configure EMAIL variable
- `envs`: Configure custom environment variables
- `shell-history`: Store custom shell history
- `ssh`: Configure ssh keys, config and known_hosts

### aws

AWS cli is really a tricky peace of tool. It's configurations come from a folder
( usually `$HOME/.aws` ) in which there are a the configuration and credential
files. This folder _apparently_ is not configurable.

Generic AWS cli documentation assume you are using the basic credential set, and
having complex configuration ( for different AWS accounts ) can get messy pretty
quickly. Plus AWS cli "profiles" needs to be named, and the only way to 
automatically set the profile to be used by the cli is the `AWS_DEFAULT_PROFILE`
env variable. 

Thanks to RTFC, the credential folder for the [botocore][botocore] library on 
which AWS cli is based can be configured [via env variables][botocore-envs]!

So this is the `HOME/.aws`, just in another place!

### bin

Sometimes you need to wrap a particular command to enable specific functionalities.
`devenv` will add this `bin` folder to your path automatically, you can place 
here profile specific executables.

If the profile has `ssh` credentials available and use a custom ssh config file
or ssh known hosts file, a `ssh` and `scp` wrappers will be created to customize
this configurations.

### email

An identity is generally tied to an email address. An environment variable `EMAIL`
is set to the specified email address ( and can be used in the loaded shell
environment ).

### envs

The most common scenario in different profiles are different environment variables.

Cloud provider credentials and authentication tokens for example.

Add vars to this file as you where adding them to `/etc/environment`, they will
all be loaded for the specific profile.

### shell-history

Generally if you use the cli a lot, you search in it's history a lot. If you have
different shell environments, other profiles history is a noise you want filtered
out. This plugin enable per-profile shell history file.

### ssh

Specific keys/certificates for this profile? Specific ssh config file? Or simply
you would like to separate the known hosts of this profile from the global one?

Look no further and place files inside this folder.

`ssh` and `scp` command will be wrapped to use the per-profile config and known_hosts
files.

A ssh agent *per profile* is loaded the first time you load a profile, so:

- each profile load will have a ssh agent ready and running
- your long ssh passphrase will be required only once (you have one, right?)

**Note** that a ssh agent bootstrap file is created in the `/tmp` folder, with
these permissions:

```
-rwx------  1 user user  132 ago 24 11:31 profile-ssh-agent.tmp
```


## How it works

`devenv` works in a pretty straightforward way: source a shell script (based on
profile) with all the required configurations and then run or a specified command
or a login shell.

The loader and runner scripts are static files, generated by the `devenv rehash`
command. This has multiple advantages:

1. you can edit them by hand, and all the changes are persisted (until your next
`devenv rehash`)
1. it's easy to understand what is the resulting environment
1. profiles are completely portable (the profile folder after the first `rehash`
are self-contained)
1. it's faster than performing text parsing, regexp and all fancy stuff at run
time!

## How to use it

Please read the Profile section before this!

Run `$ devenv` for a list of commands with brief description.

Use `devenv new` to create a new profile.  
Use `devenv run` to run spot command with a specified profile.  
Use `devenv shell` to load a shell with the configured environment.  
Use `devenv rehash` when changing the profile files/folders to create the runner
and loader scripts.  
Use `devenv whoami` to get the current profile.

## Development

Clone, hack and be happy!

General guidelines:

- submit PR with `feature-branch`
- use `rebase` and not `merge`
- respect indentation (the `.editorconfig` file is there for a reason)

## Credits

Edoardo Tenani <[@edoardotenani][twitter]>

[anyenv]: https://github.com/riywo/anyenv
[sub]: https://github.com/basecamp/sub
[twitter]: https://twitter.com/edoardotenani
[botocore]: https://github.com/boto/botocore
[botocore-envs]: https://github.com/boto/botocore/blob/d10dac4f8d812b7c58e3b8f8b117ec4f520aaec1/tests/functional/__init__.py#L19
