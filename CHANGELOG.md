# Changelog

## [2.0.0](https://github.com/iton0/hkup-cli/compare/v1.3.2...v2.0.0) (2025-11-26)


### ⚠ BREAKING CHANGES

* **cmd:** The command and internal logic module structure has been completely flattened. Subcommands (like `config` and `template`) and their logic are now defined in standalone files within the top-level `cmd/` and `internal/logic/` directories (e.g., `cmd/config/get.go` is now `cmd/config-get.go`).

### Features

* bump Go module to v2, upgrade Go/dependencies, and refactor Git logic ([887ec35](https://github.com/iton0/hkup-cli/commit/887ec3554f0ad5096e2fbc982e3560276fd046a8))
* change license from MIT to GPLv3 ([13b3c85](https://github.com/iton0/hkup-cli/commit/13b3c855b452dae40cc6ab9b51e5c174aceba143))


### Code Refactoring

* **cmd:** flatten command structure and update module paths ([165dc44](https://github.com/iton0/hkup-cli/commit/165dc447ad01b895758df763188f21260f5c382d))

## [1.3.2](https://github.com/iton0/hkup-cli/compare/v1.3.1...v1.3.2) (2025-05-20)


### Bug Fixes

* **init:** update logic to fix 'hkup init' error ([749cecc](https://github.com/iton0/hkup-cli/commit/749ceccd5669212d4ef6ed07baa160f68582b22a))
* revert commit 8f61a33 ([ea5e481](https://github.com/iton0/hkup-cli/commit/ea5e48197b4b6ecc617964b6559d6698b66677d1))
* update logic to properly handle empty values ([1b9e297](https://github.com/iton0/hkup-cli/commit/1b9e297de21b48bda4d6159dccb2e4604a193b61))

## [1.3.1](https://github.com/iton0/hkup-cli/compare/v1.3.0...v1.3.1) (2025-05-20)


### Bug Fixes

* **init:** update logic to fix 'hkup init' error ([8f61a33](https://github.com/iton0/hkup-cli/commit/8f61a331a21cb7165c75928c8b09539b33a37ac8))

## [1.3.0](https://github.com/iton0/hkup-cli/compare/v1.2.0...v1.3.0) (2025-04-24)


### Features

* **end:** add command output ([a50b826](https://github.com/iton0/hkup-cli/commit/a50b826ffc7930fde6099a017dcb9567c53b3865))


### Bug Fixes

* **status:** handle unset local git core.hooksPath ([ac59915](https://github.com/iton0/hkup-cli/commit/ac5991551160be6c129ed3424c7ca187b75bd113))
* **template:** handle unset global git core.editor ([46010b6](https://github.com/iton0/hkup-cli/commit/46010b60ee11296bf92cfa0914395e4d392ea1e0))

## [1.2.0](https://github.com/iton0/hkup-cli/compare/v1.1.0...v1.2.0) (2025-04-22)


### Features

* add status subcommand ([40e2e43](https://github.com/iton0/hkup-cli/commit/40e2e43f4af64f6f4c362287f458e4b9c2c56853))


### Bug Fixes

* **logic:** handle all errors ([f92d252](https://github.com/iton0/hkup-cli/commit/f92d25260d8a772f7255cb851cff95c91eb4ecc7))

## [1.1.0](https://github.com/iton0/hkup-cli/compare/v1.0.0...v1.1.0) (2025-02-11)


### Features

* add end subcommand ([a087ec9](https://github.com/iton0/hkup-cli/commit/a087ec97f23f7a4088df071fc21edabf09ac9b62))


### Bug Fixes

* **internal/logic/init:** update conditional logic ([db3d70f](https://github.com/iton0/hkup-cli/commit/db3d70fa477208b1baf05abe0e2d9cf60ba4da99))

## [0.5.0](https://github.com/iton0/hkup-cli/compare/v0.4.2...v0.5.0) (2025-01-15)


### Features

* add config subcommand ([1fdf8d5](https://github.com/iton0/hkup-cli/commit/1fdf8d5151615b08f03c93f90004ab03ea65e251))
* **cmd/list:** add config arg ([35ce6d5](https://github.com/iton0/hkup-cli/commit/35ce6d5728ac0b21b4e1d702913f56e83fd2714f))


### Bug Fixes

* improve HkUp config directory creation ([b9423f2](https://github.com/iton0/hkup-cli/commit/b9423f26209286c23d7da24d8048077556382b2d))
* **util:** improve robustness of config functions ([79c8af9](https://github.com/iton0/hkup-cli/commit/79c8af98888e4668080770b5f7eb1d893c894ed3))

## [0.4.2](https://github.com/iton0/hkup-cli/compare/v0.4.1...v0.4.2) (2024-12-27)


### Performance Improvements

* **internal/logic/template:** update template creation logic ([0f460cb](https://github.com/iton0/hkup-cli/commit/0f460cb18215e514199bf4ec926cc2a0836218f5))

## [0.4.1](https://github.com/iton0/hkup-cli/compare/v0.4.0...v0.4.1) (2024-12-21)


### Bug Fixes

* **internal/logic/root:** resolve gh command issue ([bb210d1](https://github.com/iton0/hkup-cli/commit/bb210d1941dab8171772eb8b82b607a24a8a6446))

## [0.4.0](https://github.com/iton0/hkup-cli/compare/v0.3.1...v0.4.0) (2024-12-18)


### Features

* **add:** initialize .hkup directory ([974762f](https://github.com/iton0/hkup-cli/commit/974762f53f504e9495ae4f8500009f7548a1cef8))
* improve list subcommand ([f2cdd16](https://github.com/iton0/hkup-cli/commit/f2cdd16438dd51f50df06fb837ad1dcdfc0e6238))
* **init:** add '--force' flag ([0a8e499](https://github.com/iton0/hkup-cli/commit/0a8e499a5296df2215a582d3752332afba02fa8c))


### Bug Fixes

* **internal/root:** make Root function more robust ([f5d36b3](https://github.com/iton0/hkup-cli/commit/f5d36b378e3450d3af79897d8beb08d21a53dd99))
* **root:** simplify internal root logic ([5b1a7d0](https://github.com/iton0/hkup-cli/commit/5b1a7d0c1f52c65dc18e2e085b9ac9cf9362edef))

## [0.3.1](https://github.com/iton0/hkup-cli/compare/v0.3.0...v0.3.1) (2024-12-05)


### Bug Fixes

* "hkup -v" produces correct version ([0671c0e](https://github.com/iton0/hkup-cli/commit/0671c0e04d9e05d8e81fc82e464248784dd2aa9d))

## [0.3.0](https://github.com/iton0/hkup-cli/compare/v0.2.1...v0.3.0) (2024-12-05)


### Features

* add config subcommand ([66a2bd8](https://github.com/iton0/hkup-cli/commit/66a2bd8d3fdae6f8baba2637b23de8d241146dec))
* add root git-wrapper functionality ([8313203](https://github.com/iton0/hkup-cli/commit/8313203b06b8b538ff5b5c14cc8b06795e1f762d))
* add template subcommand ([103b4f1](https://github.com/iton0/hkup-cli/commit/103b4f14f7af4de7ba4eb31e1c8c4a2ca0caff67))
* add windows support ([3c55f2b](https://github.com/iton0/hkup-cli/commit/3c55f2b64fbba373fbc01a98be50a6b93041c7b1))
* **util:** add file ops, prompts, and config functions ([0cb9605](https://github.com/iton0/hkup-cli/commit/0cb96050a2c69db118f3a700eaed3c81a0c0b9c1))

## [0.2.1](https://github.com/iton0/hkup-cli/compare/v0.2.0...v0.2.1) (2024-11-01)


### Bug Fixes

* correct version to 0.2.1 ([f0bd325](https://github.com/iton0/hkup-cli/commit/f0bd3251b349c60c5312a7ebb626bf37d775c9d6))

## [0.2.0](https://github.com/iton0/hkup-cli/compare/v0.1.0...v0.2.0) (2024-11-01)


### Features

* **init:** add flag group for git directory and working tree ([#1](https://github.com/iton0/hkup-cli/issues/1)) ([ed3721f](https://github.com/iton0/hkup-cli/commit/ed3721f83354dab268236e2e6ca1a033dcdc1427))
