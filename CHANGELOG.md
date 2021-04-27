# Changelog

## [1.2.0](https://www.github.com/shipperizer/furry-train/compare/v1.1.0...v1.2.0) (2021-04-27)


### Features

* add argocd manifest ([af2b543](https://www.github.com/shipperizer/furry-train/commit/af2b543f473eee3c285b781be641cd46bbaa6e48))


### Bug Fixes

* adjust certificate manifest so argocd dosn't complain ([f283784](https://www.github.com/shipperizer/furry-train/commit/f283784b65cbf90a559d15d7d7c76b3a5bf9f06a))
* adjust image-updated annotations ([1c405ec](https://www.github.com/shipperizer/furry-train/commit/1c405ec297d6b14a447f0a183e1456e4ca1e475a))

## [1.1.0](https://www.github.com/shipperizer/furry-train/compare/v1.0.3...v1.1.0) (2021-04-26)


### Features

* allow arm64 build ([51739c5](https://www.github.com/shipperizer/furry-train/commit/51739c5a0976f3a547409d4d9dda074dd85380bb))
* kaniko setup for arm64 ([3f2a3bd](https://www.github.com/shipperizer/furry-train/commit/3f2a3bd3a7e5aefea1839cf13dc7804879e8b93b))
* update ingress to networking/v1 ([d726651](https://www.github.com/shipperizer/furry-train/commit/d726651f9604ef7c5f3c311ebb3471b7f47bb3f0))


### Bug Fixes

* use contour ingress controller ([de7d860](https://www.github.com/shipperizer/furry-train/commit/de7d860b5d9da931c7c3541086f35813b9936859))

### [1.0.3](https://www.github.com/shipperizer/furry-train/compare/v1.0.2...v1.0.3) (2021-04-23)


### Bug Fixes

* drop skaffold bin from root of repo ([5aa2978](https://www.github.com/shipperizer/furry-train/commit/5aa29782403227119176e8d76d1ac8a0e7a994b2))

### [1.0.2](https://www.github.com/shipperizer/furry-train/compare/v1.0.1...v1.0.2) (2021-04-23)


### Bug Fixes

* use separate but dependent jobs to get tag ([161a5e9](https://www.github.com/shipperizer/furry-train/commit/161a5e98a16ff675aee9baff54767116882c0541))

### [1.0.1](https://www.github.com/shipperizer/furry-train/compare/v1.0.0...v1.0.1) (2021-04-23)


### Bug Fixes

* fetch all tags on release ([9e8a819](https://www.github.com/shipperizer/furry-train/commit/9e8a8191575954616c1c36e1964f68947c774390))

## 1.0.0 (2021-04-23)


### ⚠ BREAKING CHANGES

* docker build and release
* use release-please action for release management

### Features

* add docker build for main branch ([c027e49](https://www.github.com/shipperizer/furry-train/commit/c027e492a90c3880f355225542b1d01e49930383))
* docker build and release ([04ac63e](https://www.github.com/shipperizer/furry-train/commit/04ac63ee2158bb00a2582598a410ed5924c84daf))
* use release-please action for release management ([12b012e](https://www.github.com/shipperizer/furry-train/commit/12b012e590d27ccd36725c96d8f7a4a38efed9ed))


### Bug Fixes

* add code owners ([b15e5ac](https://www.github.com/shipperizer/furry-train/commit/b15e5ac9a95a9f4d8c4663d8d42622dcc9a43d1d))
* add stale GH action ([04a3cf6](https://www.github.com/shipperizer/furry-train/commit/04a3cf60ef2e4a3afbe5b4bac4fd1ef7026c08cd))
* missing checkout on release action ([be18695](https://www.github.com/shipperizer/furry-train/commit/be18695e5b489fd5f87202c645f581827770c6f9))
* push docker images ([c0e6f02](https://www.github.com/shipperizer/furry-train/commit/c0e6f02a1712afbddc4ceea107015e3146f2aef6))
