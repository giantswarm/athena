# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [1.6.1] - 2022-06-09

### Fixed

- Fix templating of `provider` value when app is installed from the `giantswarm` catalog.

## [1.6.0] - 2022-06-08

### Added

- Use SVG icon
- Add instructions on how to install `athena` in a workload cluster.
- Push `athena` to the `giantswarm` catalog
- Add labels and annotations to the `athena` chart.
- Use `app-test-suite` to execute smoke tests.
- Add schema validation for `values.yaml`.
- Use `app-build-suite` to generate `application.giantswarm.io/metadata`.


### Changed

- Allow missing `firestoreServiceAccountKey`.
- Make changes to helm chart for deployment on workload clusters.

## [1.5.4] - 2022-03-31

- Use Go 1.17, dependency updates
- Add PodDisruptionBudget

## [1.5.3] - 2022-03-03

### Fixed

- Fix Firestore collection name used for writing RUM data.

## [1.5.2] - 2022-02-14

### Fixed

- Fix RBAC rule for fix secret job.

## [1.5.1] - 2022-02-14

- No changes compared to v1.5.0.

## [1.5.0] - 2022-02-14

### Changed

- Updated Ingress resources in helm chart.
- Add workaround for Chart upgrade not working when not using lets encrypt due to changed secret type.

## [1.4.0] - 2021-09-29

### Added

- Add support for submitting real user monitoring (RUM) events.

## [1.3.0] - 2021-08-26


### Changed

- Update architect-orb to v4.1.0.

## [1.2.1] - 2021-08-20

### Changed

- Fix a bug in certs secret.


## [1.2.0] - 2021-07-13

### Changed

- Prepare helm values to configuration management.
- Update architect-orb to v3.0.0.

### Fixed

- Fix jwt-go failing dependency.

## [1.1.0] - 2021-05-06

### Added

- Add support for CIDR whitelisting in the ingress template.

## [1.0.1] - 2021-04-12

### Added

- Add network policy template.

## [1.0.0] - 2021-04-12

- First release.


[Unreleased]: https://github.com/giantswarm/athena/compare/v1.6.1...HEAD
[1.6.1]: https://github.com/giantswarm/athena/compare/v1.6.0...v1.6.1
[1.6.0]: https://github.com/giantswarm/athena/compare/v1.5.4...v1.6.0
[1.5.4]: https://github.com/giantswarm/athena/compare/v1.5.3...v1.5.4
[1.5.3]: https://github.com/giantswarm/athena/compare/v1.5.2...v1.5.3
[1.5.2]: https://github.com/giantswarm/athena/compare/v1.5.1...v1.5.2
[1.5.1]: https://github.com/giantswarm/athena/compare/v1.5.0...v1.5.1
[1.5.0]: https://github.com/giantswarm/athena/compare/v1.4.0...v1.5.0
[1.4.0]: https://github.com/giantswarm/athena/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/athena/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/giantswarm/athena/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/giantswarm/athena/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/giantswarm/athena/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/giantswarm/athena/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/athena/releases/tag/v1.0.0
