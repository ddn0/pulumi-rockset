# Rockset Resource Provider

The Rockset Resource Provider lets you manage [Rockset](https://rockset.com) resources.

## Installing Pulumi Resource Plugin

You'll need to install this plugin manually to override the default Pulumi
repository location.

```sh
pulumi plugin install resource rockset --server github://api.github.com/ddn0/
```

## Installing SDKs

This package is available for several languages/platforms:

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-rockset
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/ddn0/pulumi-rockset/sdk/go/...
```

## Configuration

The following configuration points are available for the `rockset` provider:

- `rockset:apiKey` (environment: `ROCKSET_APIKEY`) - the API key for Rockset
- `rockset:apiServer` (environment: `ROCKSET_APISERVER`) - the API endpoint for the region in which to deploy resources
