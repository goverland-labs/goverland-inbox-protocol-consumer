# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Update platform events library to collect nats metrics

## [0.0.6] - 2023-07-15

### Changed
- Refactored the base logic - currently we send to queue messages as is with base validation
- Pass ID from the webhook to the internal queue

## [0.0.5] - 2023-07-15

### Changed
- Updated platforms-events dependency to v0.0.20

## [0.0.4] - 2023-07-14

### Fixed
- Updated platform-events dependency to v0.0.15

## [0.0.3] - 2023-07-12

### Fixed
- Updated platform-events dependency to v0.0.13

## [0.0.2] - 2023-07-11

### Fixed
- Updated platform-events dependency to v0.0.11

## [0.0.1] - 2023-07-11

### Added
- Added skeleton app
- Added feed handlers
- Added Dockerfile
