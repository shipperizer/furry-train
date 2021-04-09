# furry-train

Simple web app with a status endpoint

Used mainly to test the setup of a k3s cluster on raspberryPi 4 (raspbian)

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

Kaniko is not supported officially on arm32 (it is only on 64bit), when changing distro on the Pi to [`arch linux`](https://archlinuxarm.org/platforms/armv8/broadcom/raspberry-pi-4) (@mafrosis i do this for you) i might have another go at it.

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


## Future improvements

Plan is to setup [`argocd`](https://argoproj.github.io/argo-cd/) and [`argocd  image updater`](https://argocd-image-updater.readthedocs.io/en/stable/) to pull images and update the deployment files, `skaffold` will be used only for local/manual deployment
