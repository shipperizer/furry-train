# Changelog

## [3.3.0](https://www.github.com/shipperizer/furry-train/compare/v3.2.1...v3.3.0) (2021-11-29)


### Features

* add grpc proxy code ([c9012c5](https://www.github.com/shipperizer/furry-train/commit/c9012c59bc3073fef91348bb296db744d480f467))
* add status handler ([3920bc5](https://www.github.com/shipperizer/furry-train/commit/3920bc5801f922096c511d7390b4a1b43dc682b0))

### [3.2.1](https://www.github.com/shipperizer/furry-train/compare/v3.2.0...v3.2.1) (2021-11-29)


### Bug Fixes

* use APP_NAME in docker build process ([23c7347](https://www.github.com/shipperizer/furry-train/commit/23c73474c78300a5fd869b3afd03ed2bdeb03131))

## [3.2.0](https://www.github.com/shipperizer/furry-train/compare/v3.1.1...v3.2.0) (2021-11-29)


### Features

* add gRPCs check and health ([7631d9e](https://www.github.com/shipperizer/furry-train/commit/7631d9ec206007992e47a2225b99cdb8c4fb3bb8))
* grpc app ([2d7c837](https://www.github.com/shipperizer/furry-train/commit/2d7c837d633cdd1776341ff2323d2986cf3b1ba8))


### Bug Fixes

* set json content type ([8012a58](https://www.github.com/shipperizer/furry-train/commit/8012a58995b482d643c76498da06b6b1ecbf1c76))

### [3.1.1](https://www.github.com/shipperizer/furry-train/compare/v3.1.0...v3.1.1) (2021-11-29)


### Bug Fixes

* broken api metric name ([a674def](https://www.github.com/shipperizer/furry-train/commit/a674deff122784bb29647f29b9d3c658f26093aa))

## [3.1.0](https://www.github.com/shipperizer/furry-train/compare/v3.0.0...v3.1.0) (2021-11-29)


### Features

* add monitoring pkg for prometheus ([b1bf653](https://www.github.com/shipperizer/furry-train/commit/b1bf653e46b3e733433824973ace11b33bb1a3a8))
* create a check endpoint ([7fd0a90](https://www.github.com/shipperizer/furry-train/commit/7fd0a905dd3233aeca4c527d71b602e03b2dac36))
* update web app to use miniature-monkey lib ([1b1ff62](https://www.github.com/shipperizer/furry-train/commit/1b1ff62af2a72c08ef8d3e80409c56fd621b92d3))


### Bug Fixes

* update deps ([de11177](https://www.github.com/shipperizer/furry-train/commit/de1117777483dfd43f150f142d8acb40a0ee2e14))

## [3.0.0](https://www.github.com/shipperizer/furry-train/compare/v2.0.1...v3.0.0) (2021-06-25)


### ⚠ BREAKING CHANGES

* bad path for furry-train proxy
* use kubectl to deploy argocd manifest
* dump GO_PRIVATE wrong value

### Features

* new hello message on status ([85bfdd6](https://www.github.com/shipperizer/furry-train/commit/85bfdd6e869a51e1d2f66c95a999a03eb056cc59))
* swap ingress with contour httpProxy ([b3be48c](https://www.github.com/shipperizer/furry-train/commit/b3be48ce8806552a67e88832a8daf997eb13adf9))
* use contour authorization server ([8f41226](https://www.github.com/shipperizer/furry-train/commit/8f412268c9e740662dbdfef91dbd9da682635b83))


### Bug Fixes

* adjust labels, keep rootProxy global for now ([5761210](https://www.github.com/shipperizer/furry-train/commit/576121097e5ebb10db724093b0a2d60c31ff8cee))
* adjust log level based on env var ([191d3e6](https://www.github.com/shipperizer/furry-train/commit/191d3e66b45fe6731c1dd297a9254616f6e142c5))
* bad path for furry-train proxy ([5b40ee5](https://www.github.com/shipperizer/furry-train/commit/5b40ee50c97a434daf3cb989c80e5d4399905c3a))
* drop strategic merge for rootProxy ([5260aea](https://www.github.com/shipperizer/furry-train/commit/5260aeacb104a020e45c1c24b48033c87afaf56e))
* dump GO_PRIVATE wrong value ([2c32309](https://www.github.com/shipperizer/furry-train/commit/2c32309889d43e9a23d0719902692047f0c86cbe))
* missing rootProxy from file list ([87ac58c](https://www.github.com/shipperizer/furry-train/commit/87ac58c8705e19a0af10e8e061c7d487d7c2034c))
* narrow down ingress prefix to only status endpoint ([bd6f20a](https://www.github.com/shipperizer/furry-train/commit/bd6f20a35e39a89418a38a1a642f3c99f546151d))
* patch routes, not includes ([cc67d69](https://www.github.com/shipperizer/furry-train/commit/cc67d69f3f41c5f9ea556dd2667f560d4d77fd14))
* refactor httpProxy setup ([be13451](https://www.github.com/shipperizer/furry-train/commit/be134514ca5697f99a6e065b9811997cf05a1085))
* set log level to ERROR ([78ab49f](https://www.github.com/shipperizer/furry-train/commit/78ab49f71640c1b24a2ecb0170ca1d3bd8a8ff5b))
* skaffold needs extra params ([fdb6589](https://www.github.com/shipperizer/furry-train/commit/fdb6589dc063fd6860902b1606c69c716716d7ac))
* update argocd secret ([25eac5e](https://www.github.com/shipperizer/furry-train/commit/25eac5ee027da68ac9eaee60f6dda7531959dcd3))
* update grpc health check version ([252172f](https://www.github.com/shipperizer/furry-train/commit/252172fadaeb521dc3a926c6d16fadd1431c6139))
* update kustomize ([79bd116](https://www.github.com/shipperizer/furry-train/commit/79bd116ef8695159485bc051bd553d267e14495e))
* use different patch method for kustomize ([9738a2e](https://www.github.com/shipperizer/furry-train/commit/9738a2e8e2accb317307cc5bcbf8cc45e9133d11))
* use kubectl to deploy argocd manifest ([f964ac0](https://www.github.com/shipperizer/furry-train/commit/f964ac01d0785c81bc8bcfb5dcfd7696eec4abc1))

### [2.0.1](https://www.github.com/shipperizer/furry-train/compare/v2.0.0...v2.0.1) (2021-04-27)


### Bug Fixes

* add semver regexp for image updater annotation ([e2898ed](https://www.github.com/shipperizer/furry-train/commit/e2898ed6eeb7fb3c1902bd5d8371c7e70c6c9905))
* update skaffold setup for argo ([773990d](https://www.github.com/shipperizer/furry-train/commit/773990d67991c0b8bddecba999986cd0614c9bb2))

## [2.0.0](https://www.github.com/shipperizer/furry-train/compare/v1.2.0...v2.0.0) (2021-04-27)


### ⚠ BREAKING CHANGES

* specify main branch on git
* drop kustomize type on argocd

### Bug Fixes

* add image to kustomize file ([fb18fd0](https://www.github.com/shipperizer/furry-train/commit/fb18fd0e0638685d98ae8b1f8e24d976dee8d299))
* drop kustomize type on argocd ([b8814b2](https://www.github.com/shipperizer/furry-train/commit/b8814b29e53983198797a8a25b0d1dd7289c3b88))
* specify main branch on git ([587f5ae](https://www.github.com/shipperizer/furry-train/commit/587f5ae24e8aa993277116c5e456713c637d05f0))

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
