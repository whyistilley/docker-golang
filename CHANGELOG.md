# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/); Given a change, note it below with one or more of the following keywords:

**Guidelines**

- There should be an entry for every single version
- The same types of changes should be grouped
- Versions and sections should be linkable
- The latest version comes first
- The release date of each version is displayed

Regarding types of changes, the following is a list of acceptable change types:

- `Added` — for new features
- `Deprecated` — for soon-to-be removed features
- `Fixed` — for fixes made to existing features
- `Updated` — for updates to existing features
- `Security` — for vulnerability remediation and / or mitigations

The `Unreleased` section is to be kept above released versions and serves two purposes:

- Anyone can see the changes being made for upcoming releases
- Allows a way to seamlessly track and move logged changes into a new release version section, without having to backtrack

As well, this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html); Given a version number MAJOR.MINOR.PATCH, increment the:

1. MAJOR version when you make incompatible API changes
2. MINOR version when you add functionality in a backward compatible manner
3. PATCH version when you make backward compatible bug fixes

Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.

New releases will follow the following formats when updating the `CHANGELOG.md`: 

- New release version subsection title: `## [MAJOR.MINOR.PATCH] - YYYY-MM-DD`
- Latest unreleased to latest release version comparison: `[unreleased]: https://github.com/whyistilley/docker-golang/compare/vMAJOR.MINOR.PATCH...HEAD`
- Latest release version link: `[MAJOR.MINOR.PATCH]: https://github.com/whyistilley/docker-golang/releases/tag/vMAJOR.MINOR.PATCH`

## [Unreleased]

### Added

- Golang / PostgreSQL setup
- Support for env vars in containers and local builds via Makefile
- SQL migrations to seed the PostgreSQL container
- Support for local development using docker compose

## [0.0.0] - 2024-10-26

### Added

- Initial project files

[unreleased]: https://github.com/whyistilley/docker-golang/compare/v0.0.0...HEAD
[0.0.0]: https://github.com/whyistilley/docker-golang/releases/tag/v0.0.0

