[![wercker status](https://app.wercker.com/status/432be585777f05f8f7ce56eda6def734/m "wercker status")](https://app.wercker.com/project/bykey/432be585777f05f8f7ce56eda6def734)

# Secure Service
A simple example of a microservice that authenticates secure API requests via API key. In a real-world scenario, you would compare the key and secret
against some kind of backing store of valid clients. For the purposes of this sample, we inject the valid API key via an environment variable (see the `env` file in `local_config` as well as the `buildlocal` and `runlocal` scripts for how we do this).
