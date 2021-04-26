# furry-train

Simple web app with a status endpoint

Used mainly to test the setup of a k3s cluster on raspberryPi 4 (Ubuntu server 64bit)

Checkable at `https://api.shipperizer.org/api/v0/status`

```
shipperizer@xxx 9:36:57 π  http https://api.shipperizer.org/api/v0/status
HTTP/1.1 200 OK
Content-Length: 25
Content-Type: text/plain; charset=utf-8
Date: Fri, 09 Apr 2021 08:43:37 GMT
Vary: Accept-Encoding

{
    "message": "Purple Bro"
}
```

## Build and deploy

Build setup is for multiarch support, a requirements for this is  [buildx](https://docs.docker.com/buildx/working-with-buildx/)

For `skaffold` integration i followed the suggestion [here](https://github.com/GoogleContainerTools/skaffold/tree/master/examples/custom-buildx) as ther eis no direct integration between `skaffold` and `buildx`


Images are pushed to `ghcr.io/shipperizer/furry-train-web-app`, `k3s` cluster has a secret allowing it to pull them, see the snippet below in `deployments.yaml`

```
containers:
- image: ghcr.io/shipperizer/furry-train-web-app
  name: furry-train-web-app
  envFrom:
    - configMapRef:
        name: furry-train-web-app
  name: furry-train-web-app
  ports:
  - name: http
    containerPort: 8000
imagePullSecrets:
- name: regcred-github
```

### Kaniko

For `kaniko` builds, use the `--profile kaniko` modifier on `skaffold`, for this you will need an `Opaque` secret:

```
 echo '{"auths":{"ghcr.io":{"auth":"****************"}}}' | kubectl create secret generic regcred-github-kaniko --from-file=config.json=/dev/stdin
 ```

the profile is targeted at building on an `arm64` cluster only, if you need to use a different arch change `initImage` and `image` values


## Future improvements

Plan is to setup [`argocd`](https://argoproj.github.io/argo-cd/) and [`argocd  image updater`](https://argocd-image-updater.readthedocs.io/en/stable/) to pull images and update the deployment files, `skaffold` will be used only for local/manual deployment
