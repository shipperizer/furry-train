CERT_MANAGER_RELEASE_URL?=https://github.com/jetstack/cert-manager/releases/download/v1.2.0/cert-manager.yaml
HTTP?=http
KUBECTL?=kubectl

update:
	$(HTTP) --output cert-manager.yaml --follow $(CERT_MANAGER_RELEASE_URL)

verify:
	$(KUBECTL) apply -f test-self-cert-issuer.yaml
