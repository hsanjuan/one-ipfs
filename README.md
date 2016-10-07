# one-ipfs


[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
TODO: Put more badges here.

> IPFS datastore and transfer drivers for OpenNebula

TODO: Fill out this long description.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

```
```

## Usage

In `oned.conf`

Add `ipfs` to the list of drivers in the `DATASTORE_MAD` configuration, which could look like:

```
DATASTORE_MAD = [
    EXECUTABLE = "one_datastore",
    ARGUMENTS  = "-t 15 -d dummy,fs,lvm,ceph,dev,iscsi_libvirt,vcenter,ipfs -s shared,ssh,ceph,fs_lvm,qcow2"
]
```

And do the same with the `TM_MAD` configuration:

```
TM_MAD = [
    EXECUTABLE = "one_tm",
    ARGUMENTS = "-t 15 -d dummy,lvm,shared,fs_lvm,qcow2,ssh,ceph,dev,vcenter,iscsi_libvirt,ipfs"
]
```

Then, add the following `DS_MAD_CONF`:

```
DS_MAD_CONF = [
    NAME = "ipfs", REQUIRED_ATTRS = "SOURCE", PERSISTENT_ONLY = "NO"
]
```

And the following `TM_MAD_CONF`:

```
TM_MAD_CONF = [
    NAME = "ipfs", LN_TARGET = "SYSTEM", CLONE_TARGET = "SYSTEM", SHARED = "YES",
    DS_MIGRATE = "YES"
]
```

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Hector Sanjuan
