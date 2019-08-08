# DevEnv

> A sandbox for your different developer identities

DevEnv is an utility tool for managing different "Developer Environments".  

We all have, as developers, a set of emails, credentials, tokens, keys that are
strictly related to an "environment": be it your personal life, your company,
a specific project, a specific "identity".  

We all should, as developers, keep a baseline security posture towards those 
"environments". For example avoid loading all GitHub tokens in your shell env from 
your _.bashrc_, or adding all your keys to your single SSH agent.  
But security is hard: is easy to make errors and require a wide knowledge. Every 
step you take poses a risk (this is true for long time pro, even truer for juniors).

What if you could sandbox credentials?  
What if you could keep them safe?  
What if you could avoid *simplest* mistakes?

DevEnv aims to do this, being **simple, customizable and secure by default**.

## **TOC**
* [Why](#why)
* [Installation](#installation)
* [Why is secure?](#why-is-secure)
* [How it works](#how-it-workds)
* [Profiles](#profiles)
* [Plugins](#plugins)
* [Credits](#credits)

## Why

The problem: I started working as a freelance, and with consultancies comes clients.
For sane security, I don't want ( or is not possible ) to reuse the same credentials
for every client: ssh key, AWS credentials, GitHub/GitLab/Bitbucket tokens,
are the most common things changing between different clients.

As I hadn't find any way to properly manage this, I started building a series
of helper bash scripts to load/unload the shell environment in order to provide
the specific tools I was needing when workign for a particular client.

I think I created it in 2015, after tons of use `bash` went short and I decided to 
rewrite it, in Go. mainly I wanted to add possibility for everyone to contribute
plugins, and bash does not help you with this.

## Installation

### Go Get

If you fancy golang and have it setup, you can `go get` this repository:

```
go get -u https://github.com/endorama/devenv
```

### Homebrew/Linuxbrew

```
brew install devenv
```

### Manual

Go to the [Release](https://github.com/endorama/devenv/releases) page and grab latest version.

## Old version

The old bash version is in branch `v0`, and latest releases is 0.2.1.

## License

Copyright 2019 - Edoardo Tenani <[@endorama][github]> <[@edoardotenani][twitter]>

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

