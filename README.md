# HkUp
> Your CLI tool with benefits built by [iton0](https://github.com/iton0) in [Go](https://go.dev/)!

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/iton0/hkup-cli)](https://github.com/iton0/hkup-cli/releases/latest)
[![godoc](https://godoc.org/github.com/iton0/hkup-cli?status.svg)](http://godoc.org/github.com/iton0/hkup-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/iton0/hkup-cli)](https://goreportcard.com/report/github.com/iton0/hkup-cli)

![GitHub watchers](https://img.shields.io/github/watchers/iton0/hkup-cli?style=social)
![GitHub stars](https://img.shields.io/github/stars/iton0/hkup-cli?style=social)

## Introduction
Git hooks automate processes in your workflow, enhancing code quality and consistency. Common use cases include:
- Commit Message Validation
- Environment Configuration
- Formatting
- Linting
- Testing

Despite their benefits, many developers avoid Git hooks due to a lack of awareness and perceived setup complexity.

HkUp simplifies Git hook management, allowing you to **focus on the logic and functionality of your hooks**. Plus, with HkUp, **your hooks are version-controlled**.

## Installation
### External Dependencies:
- `git`
- `curl`
- `grep`
> [!NOTE]
> Windows users need to install [Cygwin](https://cygwin.com/install.html) and add External Dependencies via the Cygwin setup program.

### Install Script
Run the script below:
```sh
curl -sSL https://raw.githubusercontent.com/iton0/hkup-cli/main/scripts/install | sh
```
> [!Tip]
> To update HkUp, simply rerun the script above. It will automatically replace your current version with the latest release.

### Uninstalling HkUp

```sh
# Locates and deletes the HkUp binary
sh -c 'rm "$(command -v 'hkup')"'
```

## Usage Quickstart
This section provides basic information about core usage. For detailed usage information run `hkup --help`.

### Initializing hkup
HkUp provides two ways of initializing:

#### 1. Creating a New Repository
After you create a Git repository, run the following command in your Git repository to initialize HkUp:

```sh
hkup init
```

This creates a **.hkup** directory and sets the local **core.hooksPath** variable. If the directory already exists, it will simply update the path variable. The path is relative, ensuring that moving your repository won’t affect hook sourcing.

#### 2. Cloning a Repository
The above command also works after cloning a repository but HkUp also provides means to wrap Git-related cloning commands for an easier process. By prepending your Git cloning command with `hkup --` you can clone your repository and initialize/reinitialize HkUp in a single command.

```sh
# Cloning the HkUp Github repository
# and initializing the .hkup folder
hkup -- git clone git@github.com:iton0/hkup-cli.git
```
> [!NOTE]
> This works even for bare Git repositories!

### Adding & Removing hooks
Add or remove hooks easily with:

```sh
hkup add <hook-name>

hkup remove <hook-name>
```

### Templates
A **template** is a pre-configured, reusable Git hook that simplifies and automates the process of setting up hooks in a Git repository. With **HkUp**, you can create, copy, edit, or remove templates, allowing for consistent and easy application of hooks without needing to rewrite scripts each time.

The templates are stored in the HkUp config templates directory that can either be found at **$XDG_CONFIG_HOME/hkup/templates** or **$HOME/.config/hkup/templates** depending on your system.

#### Naming Convention
Template files follow the naming convention `<template-name>#<hook-name>` where:
- `<template-name>` is the name of the template.
- `#` is the separator between template and hook name.
- `<hook-name>` is the specific Git hook (e.g., pre-commit, post-merge).

**Create a template**:

```sh
hkup template create
# OR
hkup template create <hook-name>
```

**Copy a template** into current working directory:

```sh
hkup template copy <template-name>
```

**Edit a template**:

```sh
hkup template edit <template-name>
```
> [!CAUTION]
> Editing a template will not update its copies.

**Remove a template**:

```sh
hkup template remove <template-name>
```

## Roadmap to v1.0.0
1. ~~windows support~~
2. ~~wrapper for git & gh clone~~
3. HkUp logo (Expected by end of January)
4. ~~Finalize HkUp configuration settings and make config subcommand visible~~
