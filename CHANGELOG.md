# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- added 

### Changed

- changed

### Removed

- removed

### Fixed

- fixed

## [1.0.0-alpha]

### Added

- New subcommands `tx verify`, `tx broadcast`, `tx spend`, `account sign`

### Changed

- `account spend`, `account sign` support the `--password` flag
- default tx fee is now 20000, in line with other SDKs

### Removed


### Fixed

- keystore.json reader was not reading kdfparams properly
- rlp from go-ethereum was encoding 0 values differently from Python/Erlang implementations


## [0.25.0-0.1.0-alpha]

### Added

- Add compatibility with epoch v0.25.0

### Removed

- Remove compatibility with epoch < v0.25.0

### Changed 

- Change keystore implementation with xsalsa-poly1305/argon2id


## [0.22.0-0.1.0-alpha]

### Added

- First alpha release of the appe-sdk-go
- Add compatibility with epoch v0.22.0
- Add command line interface supporting inspection of the chain 
- Add command line interface supporting spend transaction 
