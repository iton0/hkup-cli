# HkUp
Your CLI tool with benefits built by [iton0](https://github.com/iton0) in [Go](https://go.dev/)!

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/iton0/hkup-cli)](https://github.com/iton0/hkup-cli/releases/latest)
[![godoc](https://godoc.org/github.com/iton0/hkup-cli?status.svg)](http://godoc.org/github.com/iton0/hkup-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/iton0/hkup-cli)](https://goreportcard.com/report/github.com/iton0/hkup-cli)

![GitHub watchers](https://img.shields.io/github/watchers/iton0/hkup-cli?style=social)
![GitHub stars](https://img.shields.io/github/stars/iton0/hkup-cli?style=social)

## Introduction
Git hooks automate and implement processes in your workflow, increasing code quality and consistency.

However, many developers avoid git hooks due to a lack of awareness and the perceived complexity of setup, discouraging them from using this feature.

**HkUp** simplifies the management of git hooks, allowing you to focus on the logic and usage of your hooks instead.

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
> To update HkUp, rerun the above script.
> It will replace the current version.

#### Uninstalling hkup

```sh
# Locates and deletes the HkUp binary
sh -c 'rm "$(command -v 'hkup')"'
```

</details>

## Usage Quickstart
This section provides basic information about core usage. For detailed options run `hkup --help`.

#### Initializing hkup
Run the following command in your git repository to initialize HkUp:
```sh
hkup init
```

This command creates a **.hkup** folder and sets the local **core.hooksPath** variable. If the folder already exists, it will simply update the path variable. The path is relative, ensuring that moving your repository won’t affect hook sourcing.

#### Adding & Removing hooks
Add or remove hooks easily with:
```sh
hkup add <hook-name>
hkup remove <hook-name>
```

#### Info & Docs
There are two commands that will help you with both HkUp and git hooks:

**`hkup list {hook|lang}`**
Outputs list of either available hooks or supported languages.

**`hkup doc <hook-name>`**
Opens your browser with Git documentation for the specified git hook, helping you understand its usage.

## Future TODOs
- [ ] add either flags or subcommand for init to specify dir and worktree; also if you want the hkup folder to be hidden or not
- [ ] functionality to save custom setups (ie gitdir and workdir are not in same location)
- [ ] make an update subcommand
- [ ] store custom git hooks as templates for future use (via add template subcmd)
    - Allow users to create, store, and share templates for common hooks. Users can fetch these templates over the network.
- [ ] branch-specific hooks
- [ ] logo maybe?

## Contributing
HkUp welcomes contributions to enhance this CLI application! Before submitting a pull request (PR) for a new feature, please follow these steps:

1. **Create an Issue**:
    If you have an idea for a new feature, please create a new issue in the repository using the **feature_request** template. Provide a clear description of the feature and its potential benefits. Please note that issues submitted without using the template may be closed without warning.

2. **Wait for Approval**:
    Once you submit your issue, I’ll review it and provide feedback. If I approve the feature request, I will let you know that you're free to proceed with your PR.

3. **Submit Your PR**:
    After receiving approval, you can create your PR. Be sure to reference the issue in your PR description.

Please note that PRs submitted without prior approval through an issue may be closed without merging. This process helps us manage feature requests effectively and ensures that contributions align with the project’s goals.
