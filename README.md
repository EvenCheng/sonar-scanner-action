# Sonar-Scanner Github Action

An action that runs sonar-scanner in a docker-container and retrieves the
quality gate status if needed. In the latter case if the quality gate fails
the action fails as well.

## Usage

A workflow example:

```yaml
# ...
jobs:
  sonar-scanner-job:
    runs-on: ubuntu-latest
    steps:
    # ...
    - name: Run sonar-scanner
      uses: LowCostCustoms/sonar-scanner-action@v0.0.2
      with:
        image: sonarsource/sonar-scanner-cli:4.4
        wait-for-quality-gate: 'true'
        quality-gate-wait-timeout: 2m
        sonar-host-url: https://sonar-host.local.domain
        sonar-host-cert: ${{ secrets.sonar-host-public-cert }}
        project-file-location: sonar-project.properties
        sources-mount-point: '/app'
        log-level: 'info'
        tls-skip-verify: 'false'
        sources-location: ${{ github.workspace }}
    # ...
```

Minimal configuration:

```yaml
# ...
jobs:
  sonar-scanner-job:
    runs-on: ubuntu-latest
    steps:
    # ...
    - name: Run sonar-scanner
      uses: LowCostCustoms/sonar-scanner-action@v0.0.2
      with:
        project-file-location: sonar-project.properties
    # ...
```

## Action inputs

### image

**Default value**: "sonarsource/sonar-scanner-cli:latest"

The name and tag of the docker image containing sonar-scanner-cli tool.

### wait-for-quality-gate

**Default value**: "true"

If set to the "true" the quality gate staus will be polled after analysis is
finished. If the corresponding analysis task doesn't finish within a time
interval specified by the `quality-gate-wait-timeout` input or finishes with
a failure the action run is considered failed. If you don't need that behavior
set this input to "false".

### quality-gate-wait-timeout

**Default value**: "2m"

The maximum amount of time after which a non-finished analysis task is
considered failed. The value must be a positive integer followed by one of the
prefixes "s", "m" or "h" (meaning seconds, minutes and hours respectively), for
example "20s" or "1h".

### sonar-host-url

**Default value**: ""

The url where the SonarQube server is located. If value is empty or not set, the
`project-file-location` variable must be set!

### sonar-host-cert

**Default value**: ""

The PEM-encoded sonar-host certificate, if any.

### project-file-location

**Default value**: ""

The path to the sonar-scanner project file, relative to the `sources-location`.
Should be a relative path.

### sources-mount-point

**Default value**: "/app"

The mountpoint where the application sources specified by the `sources-location`
are mounted in the sonar-scanner docker container.

### sources-location

**Default value**: the current github workspace path

The place where the project sources are located. Should be the absolute path.

### tls-skip-verify

**Default value**: "false"

If set to the "true", sonar host certificate validation will be skipped. It's
not recommended to use this option, however it's still here for some reasons...

## Caveats

The file specified by the `project-file-location`, if any, should be located
within the `sources-location` directory.

