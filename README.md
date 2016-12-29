# gemfire-plugin
a plugin to deploy p-gemfire

This is an enaml plugin meant to generate a manifest for the p-gemfire-c0
release found: (https://github.com/c0-ops/cf-gemfire-release)


## how to run tests
### pre-reqs
  - docker
  - wercker cli 
  - (see: http://devcenter.wercker.com/docs/cli/installation)

```bash



# first run: this will always pull a fresh container image and not use anything
in your local cache

$> ./testrunner init

# or to use the cache (faster after first run)

$> ./testrunner

```


## how to sync with a new bosh release version
 **be sure to substitute your desired versions in the below examples**
### pre-reqs
  - enaml cli
```bash
# to install the enaml cli
#linux
$> wget -O /usr/local/bin/enaml https://github.com/enaml-ops/enaml/releases/download/v0.0.17/enaml-linux && chmod +x /usr/local/bin/enaml

#osx
$> wget -O /usr/local/bin/enaml https://github.com/enaml-ops/enaml/releases/download/v0.0.17/enaml-osx && chmod +x /usr/local/bin/enaml
```

### generate golang structs to be used in the plugin
```bash
# use enaml cli to auto generate objects from the target BOSH release
$> enaml generate GemFire-v1.0.6.tgz
completed generating release job structs for  GemFire-v1.0.6.tgz

# notice the creation of an enaml-gen directory
$> ls
GemFire-v1.0.6.tgz enaml-gen

# enaml-gen directory contains golang packages matching the targeted BOSH
Release's jobs and job properties
$> ls enaml-gen
locator             server              test_cluster_health test_service_health
```

