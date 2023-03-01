## [1.2.4](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.2.3...v1.2.4) (2023-03-01)


### Bug Fixes

* **deps:** update golang dependencies ([#44](https://github.com/karlderkaefer/argocd-ecr-updater/issues/44)) ([9426ad0](https://github.com/karlderkaefer/argocd-ecr-updater/commit/9426ad0165e383a6bb0567e9e9c910920087c147))

## [1.2.3](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.2.2...v1.2.3) (2023-02-02)


### Bug Fixes

* **deps:** update golang dependencies ([#38](https://github.com/karlderkaefer/argocd-ecr-updater/issues/38)) ([5a90fdc](https://github.com/karlderkaefer/argocd-ecr-updater/commit/5a90fdc1fb650e7c3c36909f27316abee3d6d353))

## [1.2.2](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.2.1...v1.2.2) (2023-01-06)


### Bug Fixes

* **deps:** update golang dependencies ([#31](https://github.com/karlderkaefer/argocd-ecr-updater/issues/31)) ([e53581a](https://github.com/karlderkaefer/argocd-ecr-updater/commit/e53581ab1a61ddc0e6be040c7524711c52ae6e9d))

## [1.2.1](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.2.0...v1.2.1) (2022-12-03)


### Bug Fixes

* check error after authenticate to ecr repository ([42bf536](https://github.com/karlderkaefer/argocd-ecr-updater/commit/42bf53667420af4d76dbf9376b46895d8c35b671))

# [1.2.0](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.7...v1.2.0) (2022-12-02)


### Bug Fixes

* use static distroless image ([740d12b](https://github.com/karlderkaefer/argocd-ecr-updater/commit/740d12b09697758becdfe119afe5e1c4138a9e5d))


### Features

* read region from secret and use in aws client ([f9f50fb](https://github.com/karlderkaefer/argocd-ecr-updater/commit/f9f50fb865c15e88871d5a0868b95489d6204071)), closes [#13](https://github.com/karlderkaefer/argocd-ecr-updater/issues/13)

## [1.1.7](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.6...v1.1.7) (2022-12-01)


### Bug Fixes

* **deps:** update golang dependencies ([#23](https://github.com/karlderkaefer/argocd-ecr-updater/issues/23)) ([2d92563](https://github.com/karlderkaefer/argocd-ecr-updater/commit/2d925631ab6f14a1377c9ea4714ed5d8393aeec5))

## [1.1.6](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.5...v1.1.6) (2022-11-16)


### Bug Fixes

* decode base64 token and extract password for registry ([7513b06](https://github.com/karlderkaefer/argocd-ecr-updater/commit/7513b06d4929c3601aa0dac0eef14258c70b200c))
* set default namespace to argocd ([aa1b6d4](https://github.com/karlderkaefer/argocd-ecr-updater/commit/aa1b6d447de981c4d038daac011b138e59ed8868))

## [1.1.5](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.4...v1.1.5) (2022-11-14)


### Bug Fixes

* **release:** bump helm chart on new version ([f3331db](https://github.com/karlderkaefer/argocd-ecr-updater/commit/f3331dbacb762a1b9a83086bfaf802eb3a3cdb03))

## [1.1.4](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.3...v1.1.4) (2022-11-14)


### Bug Fixes

* **release:** publish docker tag on tags ([727ce00](https://github.com/karlderkaefer/argocd-ecr-updater/commit/727ce00eef5cb5baea39322d3cfcc99267a422e7))

## [1.1.3](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.2...v1.1.3) (2022-11-14)


### Bug Fixes

* **release:** publish docker tag on tags ([d0ba51f](https://github.com/karlderkaefer/argocd-ecr-updater/commit/d0ba51f6f7ce35f92e5f578e629a28f6367c97fe))
* **release:** publish docker tag on tags ([c73cedd](https://github.com/karlderkaefer/argocd-ecr-updater/commit/c73ceddcbcfbc25c85e43b2f4a37fc03c6325253))

## [1.1.2](https://github.com/karlderkaefer/argocd-ecr-updater/compare/v1.1.1...v1.1.2) (2022-11-14)


### Bug Fixes

* **release:** publish docker tag on tags ([3433910](https://github.com/karlderkaefer/argocd-ecr-updater/commit/3433910ad52a69bdf3528e4e4da5f89e41af6cc3))

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
