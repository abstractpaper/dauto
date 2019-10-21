# Introduction

Dauto is an automation tool that can hook itself to git repositories. The default hook now is a `pre-commit` hook that gets executed first once you `git commit` a repository.

Command Line
--
- `hook` write a pre-commit git hook into a repository to enable dauto integration.
- `run` run a job.

Installation
--
You need go compiler to be installed and `GOPATH` env variable to be set. If it is not set:

```bash
export GOPATH=$HOME/go  # or anywhere you prefer
```
Once both are ready:

```bash
go get https://github.com/abstractpaper/dauto
```

How to Use
--
Dauto can be used to execute arbitrary commands for any git repository. It is language independent.

Here is an example hooking a go project to dauto:

```bash
dauto hook $GOPATH/foo/bar
echo '["go test -v foo/bar/./...", "echo win!"]' >> $GOPATH/foo/bar/dauto.json
```

That's it! Now try to add and commit your changes and see what happens. All commands in your `dauto.json` file have to execute successfully otherwise your commit will halt.

`dauto.json`
--
The configuration file for describing dauto jobs. Commands are executed in order; for example:

```json
[
    "go test -v github.com/abstractpaper/dauto/./...",
    "echo bang!"
]
```

Features:
* Shell commands.
* Email sending. (TODO)
* URL call. (TODO)
* Docker integration. (TODO)

API (TODO)
--
Add REST APIs that invoke command line commands.