# one-ipfs


[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
TODO: Put more badges here.

> IPFS datastore and transfer drivers for OpenNebula

The `one-ipfs` drivers allow to deploy OpenNebula VMs with images stored in [IPFS](https://ipfs.io). `one-ipfs` consists of a datastore (`ds_mad`) driver, which allows adding IPFS-backed images to OpenNebula, and a transfer driver (`tm_mad`) which allows deploying those images to OpenNebula nodes.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

### Building

The drivers are written in Go. Therefore Go needs to be installed and the Go environment needs to be set up. After that simply run:

```
make deps # Make sure the needed go dependencies are available and updated
make build
sudo make install
...
sudo make uninstall
```

`make install` honors `$ONE_LOCATION`, and defaults to `/var/lib/one`.

### Configuring OpenNebula

In `oned.conf`:

  - Add `ipfs` to the list of drivers in the `DATASTORE_MAD` configuration, which could look like:

```
DATASTORE_MAD = [
    EXECUTABLE = "one_datastore",
    ARGUMENTS  = "-t 15 -d dummy,fs,lvm,ceph,dev,iscsi_libvirt,vcenter,ipfs -s shared,ssh,ceph,fs_lvm,qcow2"
]
```

  - Add `ipfs` to the list of drivers in the `TM_MAD` configuration:

```
TM_MAD = [
    EXECUTABLE = "one_tm",
    ARGUMENTS = "-t 15 -d dummy,lvm,shared,fs_lvm,qcow2,ssh,ceph,dev,vcenter,iscsi_libvirt,ipfs"
]
```

  - Add the following `DS_MAD_CONF` section:

```
DS_MAD_CONF = [
    NAME = "ipfs", REQUIRED_ATTRS = "SOURCE", PERSISTENT_ONLY = "NO"
]
```

  - Add the following `TM_MAD_CONF` section:

```
TM_MAD_CONF = [
    NAME = "ipfs", LN_TARGET = "SYSTEM", CLONE_TARGET = "SYSTEM", SHARED = "YES",
    DS_MIGRATE = "YES"
]
```


## Usage

TODO.

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Hector Sanjuan
