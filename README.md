> [!IMPORTANT]
> **This repository has moved to Codeberg.**
> All future development, issues, and releases are at **[codeberg.org/iton0/hkup-cli](https://codeberg.org/iton0/hkup-cli)**.
> This GitHub repo is archived and no longer maintained.

---

<img alt="HkUp Logo" align="left" height="128" src="https://codeberg.org/iton0/hkup-cli/raw/branch/main/assets/logo.svg" /> *Your CLI tool with benefits built by [iton0](https://codeberg.org/iton0) in [Go](https://go.dev/)!*

[![Codeberg release](https://img.shields.io/gitea/v/release/iton0/hkup-cli?gitea_url=https://codeberg.org)](https://codeberg.org/iton0/hkup-cli/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/codeberg.org/iton0/hkup-cli/v2.svg)](https://pkg.go.dev/codeberg.org/iton0/hkup-cli/v2)
[![Go Report Card](https://goreportcard.com/badge/codeberg.org/iton0/hkup-cli/v2)](https://goreportcard.com/report/codeberg.org/iton0/hkup-cli/v2)

#
Git hooks automate processes in your workflow, enhancing code quality and consistency. Common use cases include:
- Commit Message Validation
- Environment Configuration
- Formatting
- Linting
- Testing

Despite their benefits, many developers avoid Git hooks due to a lack of awareness and perceived setup complexity.

HkUp simplifies Git hook management, allowing you to **focus on the logic and functionality of your hooks**. Plus, with HkUp, **your hooks are version-controlled**.

## Installation

This script downloads and installs the latest release binary from Codeberg.

### External Dependencies:

The following core system utilities are required for the installation script to run successfully:

* **`curl`**: Used to fetch the latest release information and download the binary.
* **`grep`** & **`sed`**: Essential command-line text processing tools.
* **`jq`** (Recommended): Provides reliable JSON parsing for fetching the latest version from the Codeberg API. If `jq` is not found, the script will fall back to using `grep` and `sed`.
* **`install`** (Optional): A standard utility used for atomic file replacement and securely setting executable permissions. If not found, the script will fall back to using `mv` and `chmod`.

> [!NOTE]
> Windows users must execute the installation script within a Unix-like environment such as **Cygwin**, **MSYS2**, or **WSL** (Windows Subsystem for Linux) and ensure the dependencies above are installed within that environment.

### Install Script
Run the script below:
```sh
curl -sSL https://codeberg.org/iton0/hkup-cli/raw/branch/main/scripts/install | bash
```
> [!Tip]
> To update HkUp, simply rerun the script above. It will automatically replace your current version with the latest stable release.

### Uninstalling HkUp

```sh
# Locates and deletes the HkUp binary
sh -c 'rm "$(command -v 'hkup')"'
```

## Usage Quickstart
This section provides basic information about core usage. For detailed usage information run `hkup --help`.

### Initializing HkUp
After you create a Git repository, run the following command in your Git repository to initialize HkUp:

```sh
hkup init
```

This creates a **.hkup** directory and sets the local **core.hooksPath** variable. If the directory already exists, it will simply update the path variable. The path is relative, ensuring that moving your repository won't affect hook sourcing.

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

## Credits
[DeAndre Sawyer](https://deandresawyer.myportfolio.com/) for the logo.
