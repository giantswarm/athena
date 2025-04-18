# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Make Ingress annotations configurable via values (`.Values.ingress.annotations`)

### Changed

- Resolve golangci-lint v2 problems in athena by removing redundant explicit references to embedded Resolver struct fields to resolve staticcheck linter warning.

## [1.13.1] - 2025-02-03

### Removed

- Removed unused chart value `.secret.firestoreServiceAccountKey` and `.secret`.

## [1.13.0] - 2025-01-16

### Removed

- Removed analytics functionality, kept GraphQL API and optional Helm values for compatibility reasons.

### Changed

- Changed ownership to Team Shield

## [1.12.4] - 2024-08-13

### Added

- Made GraphQL introspection configurable and disabled by default

### Changed

- Change ImagePullPolicy from Always to IfNotPresent to reduce image network traffic.

## [1.12.3] - 2024-07-30

### Changed

- Bump architect-orb@5.3.1 to fix CVE-2024-24790.

## [1.12.2] - 2024-06-12

### Changed

- Default clusterIssuer to letsencrypt-giantswarm and update ingress annotation

## [1.12.1] - 2023-12-11

### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.
- Updated `github.com/nats-io/nkeys` to `v0.4.6` to address CVE-2023-46129, resolving the issue of using a hard-coded cryptographic key in the nkeys library.

### Security

- Resolved a high-severity security vulnerability (CVE-2023-46129) associated with `github.com/nats-io/nkeys` by updating to version `v0.4.6`. This update ensures enhanced security in cryptographic key handling.

### Dependencies

- Updated dependencies to ensure compatibility and security with the latest `nkeys` version.

## [1.12.0] - 2023-10-02

## Changed

- Propagate `global.podSecurityStandards.enforced` value set to `false` for PSS migration

## [1.11.0] - 2023-08-02

### Added

- Add `.Values.kubernetes.api.port` for workload clusters

## [1.10.3] - 2023-07-13

### Changed

- Make `ingressClassName` configurable

## [1.10.2] - 2023-07-04

### Changed

- Update deployment to be PSS compliant.

## [1.10.1] - 2023-06-19

### Changed

- Add template function to determine cluster codename

## [1.10.0] - 2023-06-15

### Removed

- Remove obsolete 'fixJob' that replaces certificate secrets in certain situation.
- Stop pushing to `openstack-app-collection`.
- Remove pull secret.

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

[Unreleased]: https://github.com/giantswarm/athena/compare/v1.13.1...HEAD
[1.13.1]: https://github.com/giantswarm/athena/compare/v1.13.0...v1.13.1
[1.13.0]: https://github.com/giantswarm/athena/compare/v1.12.4...v1.13.0
[1.12.4]: https://github.com/giantswarm/athena/compare/v1.12.3...v1.12.4
[1.12.3]: https://github.com/giantswarm/athena/compare/v1.12.2...v1.12.3
[1.12.2]: https://github.com/giantswarm/athena/compare/v1.12.1...v1.12.2
[1.12.1]: https://github.com/giantswarm/athena/compare/v1.12.0...v1.12.1
[1.12.0]: https://github.com/giantswarm/athena/compare/v1.11.0...v1.12.0
[1.11.0]: https://github.com/giantswarm/athena/compare/v1.10.3...v1.11.0
[1.10.3]: https://github.com/giantswarm/athena/compare/v1.10.2...v1.10.3
[1.10.2]: https://github.com/giantswarm/athena/compare/v1.10.1...v1.10.2
[1.10.1]: https://github.com/giantswarm/athena/compare/v1.10.0...v1.10.1
[1.10.0]: https://github.com/giantswarm/athena/compare/v1.9.4...v1.10.0
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
