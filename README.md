[![GoDoc](https://godoc.org/github.com/endorama/devenv?status.svg)](https://godoc.org/github.com/endorama/devenv)

# DevEnv

> A sandbox for your different profiles

What if you could *sandbox* credentials?  
What if you could avoid *simple* mistakes?  
What if you could backup them *safely*?  

For every professional working on the web, credential compromise is a ever present risk.

We all have, as developers or operators, a set of emails, credentials, tokens, keys 
that are strictly related to an "environment": be it your personal life, your company,
a specific project, a specific "identity".

We all should keep a baseline security posture towards those "environments".  
But security is hard: making errors is easy and require deep knowledge. Every 
step you take poses a risk (this is true for long time pro, even truer for juniors).

DevEnv aims to help you keeping those environment protected, by being **simple, customizable and 
secure by default**.

*NB*: No tool can prevent security breaches, and this one make no exception. But
tools can make more usable some common patterns that would be a burden to do manually
and error prone.

## **TOC**

* [Why](#why)
* [Installation](#installation)
* [How it works](#how-it-workds)
* [Why is secure?](#why-is-secure)
* [Profiles](#profiles)
* [Plugins](#plugins)
* [Credits](#credits)

## Why

The problem: I started working as a freelance, and with consultancies comes clients.
For sane security, I didn't want (or is not possible) to reuse the same credentials
for every client: ssh key, AWS credentials, GitHub/GitLab/Bitbucket tokens,
are the most common things changing between different clients.

As I hadn't find any way to properly manage this, I started building a series
of helper bash scripts to load/unload the shell environment in order to provide
the specific tools I was needing when workign for a particular client.

I think I created it in 2015, after tons of use `bash` went short and I decided to 
rewrite it, in Go. mainly I wanted to add possibility for everyone to contribute
plugins, and bash does not help you with this.

## Installation

### From Source

Clone this repository, lookup current Golang development version in `.tool-versions`
and compile it.

Optionally, checkout the release tag.

### Go Get

If you fancy golang and have it setup, you can `go get` this repository:

```
go get -u https://github.com/endorama/devenv
```

### Homebrew/Linuxbrew

You can install latest stable release using Homebrew (both macOS and Linux):

```
brew install devenv
```

### Manual

Go to the [Release](https://github.com/endorama/devenv/releases) page and grab latest version.

## Old version (Bash)

The old bash version is in branch `v0`, and latest releases is `0.2.1`.

## How it works

`devenv` takes a set of files inside a folder and via a plugin system produce 2 output files:
- a shell loader: this file contains profile setup instruction and then `exec` the speficied
  shell; useful for long running interactive sessions
- a shell runner: this files contains profile setup instruction and the `exec` the specified
  command; useful for one off commands

This 2 files will be used by `shell` and `run` commands, but may be used as is, as both are valid
`bash` files. This is a major advantage of `devenv`: even if the cli breaks, you're dealing only with
BASH files, easy enough to be modified manually. An hidden advantage are easier backups.

## Why is secure?

`devenv` aims to reduce security risks by improving your security posture. It uses sane and secure 
defaults, reducing toil in manual operations.

How do you backup your SSH key? How do 

Currently `devenv` is not complete, and is not the best solution. You are still required to know and
perform some steps knowing how to posture yourself securely, but with the growth is expected to 
encompass more and more best practices, removing the burden from the user.

## Key Concepts

### Evironments

Even if the tool is called `devenv`, no reference to `environments` are present in the documentation.  
Environments has been widely used by a multitide of softwares, so we will use `profiles` instead.

### Profiles

Profiles are set of credentials you want to isolate between each other. Each profile may contain
its credentials allowing you to load them on demand, otherwise hidden from your system.

### Plugins

`devenv` is build around a plugin system to enhance its capabilities. Each plugin is responsible
to manage configurations, creating files for a profiles and producing the shell code to be 
integrated in the shell loader and runner files.

Plugins lives in this codebase (for simplicity, may be split at later time) and are availble under
[`internal/plugins`](https://github.com/endorama/devenv/tree/master/internal/plugins).  
Each plugin is a separate go package implementing the `Pluggable` interface.

Each plugin may implement one or more of the optional interfaces for extending its capabilities:

- `Configurable`
- `Generator`
- `Setuppable`

Additional details about pluing interfaces may be found in the docs.

*Available plugins*:

-	[aws](https://github.com/endorama/devenv/blob/master/internal/plugins/aws.go)
- [bin](https://github.com/endorama/devenv/blob/master/internal/plugins/bin.go)
- [email](https://github.com/endorama/devenv/blob/master/internal/plugins/email.go)
- [envs](https://github.com/endorama/devenv/blob/master/internal/plugins/envs.go)
- [shell-history](https://github.com/endorama/devenv/blob/master/internal/plugins/shell-history.go)
- [ssh](https://github.com/endorama/devenv/blob/master/internal/plugins/ssh.go)

## License

Copyright 2019-2020 - Edoardo Tenani ([@endorama][github])

Licensed under Mozilla Public License 2.0.

A copy is available in [LICENSE](./LICENSE).  
An [online copy](https://choosealicense.com/licenses/mpl-2.0/) is available.

## Credits

DevEnv CLI interface is heavily inspired by the beautiful cli interface of [rbenv][rbenv].

It's `0.x` version was based on the awesome [37's Signals sub][sub] utility.

I had some mighty beta testers, that tried, debugged and broke my bash scripts:
- [@matjack1](https://github.com/matjack1)

To whom goes my THANK YOU! ^^


[rbenv]: https://github.com/rbenv/rbenv
[sub]: https://github.com/basecamp/sub
[github]:  https://github.com/endorama

