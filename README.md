# Network Cloud Manifest Helper

This utility helps parse network cloud site manifests to retrieve useful information.

## Building

```
go build
```

## Commands Implemented

### getip

The `getip` command allows you to retrieve the IP address of nodes.

```
$ ./nc-manifest-helper getip -h
Usage of getip:
  -host string
    	The hostname of the host or genesis
  -site string
    	The site name (e.g. mtn52a)
  -sites-dir string
    	Path to aic-clcp-sites-manifests (default "/data/alan/Workbench/aic-clcp-site-manifests")

```

You can retrieve the IP of a specific host or pass in the string `genesis`

```
$ ./nc-manifest-helper getip -site mtn52a -host=genesis
1.2.3.4
```


