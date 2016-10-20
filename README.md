# one-ipfs


[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
TODO: Put more badges here.

> IPFS datastore and transfer drivers for OpenNebula

The `one-ipfs` drivers allow to deploy OpenNebula VMs with images stored in [IPFS](https://ipfs.io), a distributed/decentralized filesystem. `one-ipfs` consists of a datastore (`ds_mad`) driver, which allows adding IPFS-backed images to OpenNebula, and a transfer driver (`tm_mad`) which allows deploying those images to OpenNebula nodes.

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

### Configuring IPFS

You will need an IPFS daemon running on every OpenNebula host (including the Frontend):

  - Learn how to install IPFS: https://ipfs.io/docs/install/
  - Learn how to run IPFS: https://ipfs.io/docs/getting-started/


## Usage

### Creating the IPFS datastore

```
$> cat ipfs-datastore.template
NAME=IPFS
DS_MAD=ipfs
TM_MAD=ipfs
$> onedatastore create -c 0 ipfs-datastore.template
```

### Adding an image

IPFS supported by the driver are in the form:

  - `/ipfs/<hash>` - Supported by `PATH` and `SOURCE` attributes
  - `/ipns/<id>` - Supported by `PATH` and `SOURCE` attributes
  - `fs:/ipfs/<hash>` - Supported by `PATH` attribute only
  - `fs:/ipns/<id>` - Supported by `PATH` attribute only

The OpenNebula CLI won't allow to add IPFS images by `PATH` due to a strict safeguard check (any non http(s) paths get interpreted as filesystems paths). Therefore they need to be created with `--source` providing size manually. This limitation is not present in Sunstone, where `PATH` is mandatory instead, an it is not possible to add images by indicating only `SOURCE`. Adding images by `PATH` via Sunstone has the advantange that the IPFS IDs are checked for correctness and the image size is automatically computed.


#### CLI

Using an IPFS hash:

```
oneimage create -d IPFS --name "Slux Linux" --type OS --source /ipfs/QmeVJdKvn5wPNBZGzPSjcc8WZQjWCnCADdnqauS1AKhAcw --size 225
```

Using an IPNS address:

```
oneimage create -d IPFS --name "Slux Linux" --type OS --source /ipns/QmXZrtE5jQwXNqCJMfHUTQkvhQ4ZAnqMnmzFMJfLewur2n --size 225
```

#### Sunstone

[Sunstone screenshot](https://ipfs.io/ipfs/QmRkekd6KAR7wXwZL9ewp5t4JvS53anaTU9Qi2ANApsS9G)


### Launching Virtual Machines

You can use the registered images and attach them to VMs like any other image. Upon deployment, they are copied to the system datastore by the IPFS daemon running in the hypervisor and run from there.

Currently, only **non-persistent** images are supported, and operations like migrations or snapshots are not tested/implemented.


## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Hector Sanjuan
