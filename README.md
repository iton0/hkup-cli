# HkUp
> Your CLI tool with benefits built by [iton0](https://github.com/iton0) in [Go](https://go.dev/)!

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/iton0/hkup-cli)](https://github.com/iton0/hkup-cli/releases/latest)
[![godoc](https://godoc.org/github.com/iton0/hkup-cli?status.svg)](http://godoc.org/github.com/iton0/hkup-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/iton0/hkup-cli)](https://goreportcard.com/report/github.com/iton0/hkup-cli)

![GitHub watchers](https://img.shields.io/github/watchers/iton0/hkup-cli?style=social)
![GitHub stars](https://img.shields.io/github/stars/iton0/hkup-cli?style=social)

## Introduction
Git hooks automate and implement processes in your workflow, increasing code quality and consistency.

Common use cases include:
- Commit Message Validation
- Environment Configuration
- Formatting
- Linting
- Testing

However, many developers avoid Git hooks due to a lack of awareness and the perceived complexity of setup, discouraging them from using this feature.

**HkUp** simplifies the management of Git hooks, allowing you to focus on the logic and usage of your hooks instead.

## Installation
External Dependencies:
- `git`
- `curl`
- `grep`

Run the script below (supports Linux and macOS):

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

</details>

## Usage Quickstart
This section provides basic information about core usage. For detailed usage information run `hkup --help`.

### Initializing hkup
Run the following command in your git repository to initialize HkUp:
```sh
hkup init
```

This creates a **.hkup** directory and sets the local **core.hooksPath** variable. If the directory already exists, it will simply update the path variable. The path is relative, ensuring that moving your repository won’t affect hook sourcing.

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
Template files follow the naming convention:
`<template-name>#<hook-name>`
Where:
- `<template-name>` is the name of the template.
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
>[!CAUTION]
> Editing a template will not update its copies.

**Remove a template**:
```sh
hkup template remove <template-name>
```

## Roadmap to v1.0.0
1.  windows support
2.  wrapper for git init & clone and gh repo create & clone
3.  HkUp logo (may or may not keep this one)
4.  better test coverage
5.  Allow users to create, store, and share templates. Users can fetch these templates over internet (may need to make another repo for this).
