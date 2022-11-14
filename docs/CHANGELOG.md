## [1.1.1](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.0...v1.1.1) (2022-11-14)


### Bug Fixes

* **release:** publish docker tag on tags ([eb189e1](https://github.com/karlderkaefer/argocd-ecr-updater/commit/eb189e14fc553d38b48068c7c27715c59a928f2c))

# [1.1.0](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.0.1...v1.1.0) (2022-11-07)


### Bug Fixes

* fix container port ([7dc3e56](https://github.com/karlderkaefer/argocd-ecr-updater/commit/7dc3e56349a437d33c8c50681f7b2260ef636c12))
* fix rbac permission and remove service by default ([a383505](https://github.com/karlderkaefer/argocd-ecr-updater/commit/a38350587455538e73fc101d4c5aee146f2b3d4b))
* **security:** fix violating kyverno pod security policies ([14ee558](https://github.com/karlderkaefer/argocd-ecr-updater/commit/14ee558f990294d4397daf0c8bf10d37ec4ec7c2))


### Features

* add initial helm chart ([2ad7f03](https://github.com/karlderkaefer/argocd-ecr-updater/commit/2ad7f03c277ee9c84cc472fbb2655c6956d3aadf))
* add initial helm chart ([9dbcac6](https://github.com/karlderkaefer/argocd-ecr-updater/commit/9dbcac65a83569af480a1c2212f204b5a6a9c28d))
* add initial helm chart ([8cc8ece](https://github.com/karlderkaefer/argocd-ecr-updater/commit/8cc8ece71d942a8c82229becc1600704772a1a16))
* add initial helm chart ([8436bee](https://github.com/karlderkaefer/argocd-ecr-updater/commit/8436bee4ceadd8c1427a1f2e30f62b307a345541))
* add initial helm chart ([abfaffe](https://github.com/karlderkaefer/argocd-ecr-updater/commit/abfaffe4442bca72948b8542e8bf600d58d564dc))
* **helm:** set default namespace to argocd ([7e269bb](https://github.com/karlderkaefer/argocd-ecr-updater/commit/7e269bba12677913979ae951d1b43ef28031e135))
* **helm:** set default resource limits ([00e0829](https://github.com/karlderkaefer/argocd-ecr-updater/commit/00e0829a70c50f75f19fe38758a59bf3bd6579c4))
* rename prefix from ECR to ARGOCD_ECR_UPDATER ([019e148](https://github.com/karlderkaefer/argocd-ecr-updater/commit/019e1487427132a366c4676eae9e8517694111d5))
* set default interval for refresh to 6h ([726a12f](https://github.com/karlderkaefer/argocd-ecr-updater/commit/726a12fd634eea2766f245e6ba908312297e64fb))

## [1.0.1](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.0.0...v1.0.1) (2022-11-07)


### Bug Fixes

* **deps:** update module github.com/spf13/viper to v1.14.0 ([#7](https://github.com/karlderkaefer/argocd-ecr-updater/issues/7)) ([2fa4da8](https://github.com/karlderkaefer/argocd-ecr-updater/commit/2fa4da8817b6208d5e2816f9f2877869b6afec96))

# 1.0.0 (2022-11-06)


### Features

* add docker file and ci workflow ([e77134a](https://github.com/karlderkaefer/argocd-ecr-updater/commit/e77134aa09eece90d12cbf6fedd93bc89b102b24))
* init first version ([0664404](https://github.com/karlderkaefer/argocd-ecr-updater/commit/0664404075b662c06b0ae353b928e941aff2db0a))
