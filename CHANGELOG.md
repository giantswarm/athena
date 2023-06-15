# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Add template function to determine cluster codename

### Removed

- Remove obsolete 'fixJob' that replaces certificate secrets in certain situation.
- Stop pushing to `openstack-app-collection`.

## [1.9.4] - 2023-04-27

### Changed

- Changed pod disruption budget to percentage.
- Remove shared app collection from circle CI

## [1.9.3] - 2023-04-04

### Changed

- Push to `kvm` and `capz` app collections.

## [1.9.2] - 2023-03-02

- Add additional annotations on all `ingress` objects to support DNS record creation via `external-dns`
- Added the Runtime Default seccompprofile

## [1.9.1] - 2022-12-21

### Fixed

- Fixed labels in certs-secret Helm template

### Removed

- Remove CA pem example from readme as it is no longer required

## [1.9.0] - 2022-11-29

### Added

- Add option to configure custom clusterIssuer for certificate generation.

### Changed

- Marked `caPem` in schema as deprecated.

## [1.8.2] - 2022-11-24

### Added

- Added support for polymorphic values of `managementCluster` - it can be defined either as `string` or as `object`.

## [1.8.1] - 2022-11-14

### Fixed

- Fix typo in `CiliumNetworkPolicy` endpoint selector.

## [1.8.0] - 2022-11-14

### Added

- Added `CiliumNetworkPolicy` for the CRD install job.

### Changed

- The helm job that installs CRDs is not removed if the job fails.

## [1.7.6] - 2022-10-25

### Added

- Add CI job to push to `openstack-app-collection`.
- Add CI job to push to `gcp-app-collection`.
- Configured read header timeout in the http server

### Fixed

- Dependency updates for security.

## [1.7.5] - 2022-09-15

### Fixed

- Updated CA certificate parsing to sanitize leading and trailing white spaces and newlines

## [1.7.4] - 2022-09-12

- Do not display `nil` value for CA in athena configmap.

## [1.7.3] - 2022-09-12

### Added

- Add environment variable to override the cluster CA cert.

## [1.7.2] - 2022-08-09

### Added

- Add CI job to push to `capa-app-collection`.

## [1.7.1] - 2022-07-06

### Added

- Replace `|-` with `|` in CA templating.

## [1.7.0] - 2022-07-05

### Added

- Enforce the management cluster name in workload cluster info and automatically fill in the CA.

### Fixed

- Fix broken relative link in README

## [1.6.2] - 2022-06-09

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


[Unreleased]: https://github.com/giantswarm/athena/compare/v1.9.4...HEAD
[1.9.4]: https://github.com/giantswarm/athena/compare/v1.9.3...v1.9.4
[1.9.3]: https://github.com/giantswarm/athena/compare/v1.9.2...v1.9.3
[1.9.2]: https://github.com/giantswarm/athena/compare/v1.9.1...v1.9.2
[1.9.1]: https://github.com/giantswarm/athena/compare/v1.9.0...v1.9.1
[1.9.0]: https://github.com/giantswarm/athena/compare/v1.8.2...v1.9.0
[1.8.2]: https://github.com/giantswarm/athena/compare/v1.8.1...v1.8.2
[1.8.1]: https://github.com/giantswarm/athena/compare/v1.8.0...v1.8.1
[1.8.0]: https://github.com/giantswarm/athena/compare/v1.7.6...v1.8.0
[1.7.6]: https://github.com/giantswarm/athena/compare/v1.7.5...v1.7.6
[1.7.5]: https://github.com/giantswarm/athena/compare/v1.7.4...v1.7.5
[1.7.4]: https://github.com/giantswarm/athena/compare/v1.7.3...v1.7.4
[1.7.3]: https://github.com/giantswarm/athena/compare/v1.7.2...v1.7.3
[1.7.2]: https://github.com/giantswarm/athena/compare/v1.7.1...v1.7.2
[1.7.1]: https://github.com/giantswarm/athena/compare/v1.7.0...v1.7.1
[1.7.0]: https://github.com/giantswarm/athena/compare/v1.6.2...v1.7.0
[1.6.2]: https://github.com/giantswarm/athena/compare/v1.6.1...v1.6.2
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
