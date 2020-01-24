[Collins](http://tumblr.github.io/collins/)
=======

Collins was designed from the beginning to represent assets in the simplest way possible. This simplicity makes for 
an efficient data-model and allows for a large degree of flexibility.

# Creating a node classifier

Although not obvious collins has a classification system built in that is easy to use and flexible. You can do so by
creating a configuration asset in the drop down. To create a configuration asset that will let you find all servers
that has 1Gb NICs, 32GB of RAM, and 138GB hard drive you can follow these steps.

1. Create an configuration and name it something like "compute_node". Do not generate IPMI info and you can leave status blank as well.
2. Add another tag called `IS_NODECLASS` with collins-cli this looks like `collins set -t compute_node -a IS_NODECLASS:true`.
3. Add the attributes associated with the nodeclass for example `collins set -t compute_node -a DISK_STORAGE_TOTAL:146163105792 -a MEMORY_SIZE_TOTAL:34359738368 -a NIC_SPEED:1000000000`.

If you now go to the compute_node tag in collins you will see a link that lets you search for Unallocated or All nodes 
that match the above criteria. If you don't know what the stats of your node are you can always use the collins-shell
to get the attributes associated with a node you know has the stats you want like this `collins-shell asset get M000001`.

The main documentation although not super in depth is available [here](http://tumblr.github.io/collins/configuration.html#node%20classifier).

Here is an example print out of all the different things you can set on your node classifier.

```
+----+--------------------------+----------+--------------------------+----------------------------------------------------------+------------+
| id | name                     | priority | label                    | description                                              | value_type |
+----+--------------------------+----------+--------------------------+----------------------------------------------------------+------------+
|  1 | SERVICE_TAG              |        1 | Service Tag              | Vendor supplied service tag                              |          1 |
|  2 | CHASSIS_TAG              |        1 | Chassis Tag              | Tag for asset chassis                                    |          1 |
|  3 | RACK_POSITION            |        1 | Rack Position            | Position of asset in rack                                |          1 |
|  4 | POWER_PORT               |        2 | Power Port               | Power port of asset                                      |          1 |
|  5 | SWITCH_PORT              |        2 | Switch Port              | Switch port that asset is connected to                   |          1 |
|  6 | CPU_COUNT                |       -1 | CPU Count                | Number of physical CPUs in asset                         |          1 |
|  7 | CPU_CORES                |       -1 | CPU Cores                | Number of cores per physical CPU                         |          1 |
|  8 | CPU_THREADS              |       -1 | CPU Threads              | Number of threads per CPU core                           |          1 |
|  9 | CPU_SPEED_GHZ            |        3 | CPU Speed                | CPU Speed in GHz                                         |          1 |
| 10 | CPU_DESCRIPTION          |       -1 | CPU Description          | CPU description, vendor labels                           |          1 |
| 11 | MEMORY_SIZE_BYTES        |       -1 | Memory                   | Size of Memory Stick                                     |          1 |
| 12 | MEMORY_DESCRIPTION       |       -1 | Memory Description       | Memory description, vendor label                         |          1 |
| 13 | MEMORY_SIZE_TOTAL        |        4 | Memory Total             | Total amount of available memory in bytes                |          1 |
| 14 | MEMORY_BANKS_TOTAL       |       -1 | Memory Banks             | Total number of memory banks                             |          1 |
| 15 | NIC_SPEED                |        5 | NIC Speed                | Speed of nic, stored as bits per second                  |          1 |
| 16 | MAC_ADDRESS              |        2 | MAC Address              | MAC Address of NIC                                       |          1 |
| 17 | NIC_DESCRIPTION          |       -1 | NIC Description          | Vendor labels for NIC                                    |          1 |
| 18 | DISK_SIZE_BYTES          |       -1 | Disk Size                | Disk size in bytes                                       |          1 |
| 19 | DISK_TYPE                |        6 | Inferred disk type       | Inferred disk type: SCSI, IDE or FLASH                   |          1 |
| 20 | DISK_DESCRIPTION         |       -1 | Disk Description         | Vendor labels for disk                                   |          1 |
| 21 | DISK_STORAGE_TOTAL       |        7 | Total disk storage       | Total amount of available storage                        |          1 |
| 22 | LLDP_INTERFACE_NAME      |       -1 | LLDP Interface Name      | Interface name reported by lldpctl                       |          1 |
| 23 | LLDP_CHASSIS_NAME        |       -1 | LLDP Chassis Name        | Chassis name reported by lldpctl                         |          1 |
| 24 | LLDP_CHASSIS_ID_TYPE     |       -1 | LLDP Chassis ID Type     | Chassis ID Type reported by lldpctl                      |          1 |
| 25 | LLDP_CHASSIS_ID_VALUE    |       -1 | LLDP Chassis ID Value    | Chassis ID Value reported by lldpctl                     |          1 |
| 26 | LLDP_CHASSIS_DESCRIPTION |       -1 | LLDP Chassis Description | Chassis Description reported by lldpctl                  |          1 |
| 27 | LLDP_PORT_ID_TYPE        |       -1 | LLDP Port ID Type        | Port ID Type reported by lldpctl                         |          1 |
| 28 | LLDP_PORT_ID_VALUE       |       -1 | LLDP Port ID Value       | Port ID Value reported by lldpctl                        |          1 |
| 29 | LLDP_PORT_DESCRIPTION    |       -1 | LLDP Port Description    | Port Description reported by lldpctl                     |          1 |
| 30 | LLDP_VLAN_ID             |       -1 | LLDP VLAN ID             | VLAN ID reported by lldpctl                              |          1 |
| 31 | LLDP_VLAN_NAME           |       -1 | LLDP VLANE Name          | VLAN name reported by lldpctl                            |          1 |
| 32 | INTERFACE_NAME           |       -1 | Interface Name           | Name of physical interface, e.g. eth0                    |          1 |
| 33 | INTERFACE_ADDRESS        |       -1 | IP Address               | Address on interface, e.g. 10.0.0.1                      |          1 |
| 34 | BASE_PRODUCT             |        1 | Base Product             | Formal product model designation                         |          1 |
| 35 | BASE_VENDOR              |        1 | Base Vendor              | Who made your spiffy computer?                           |          1 |
| 36 | BASE_DESCRIPTION         |       -1 | Base Description         | How does your computer introduce itself?                 |          1 |
| 37 | BASE_SERIAL              |       -1 | Base Serial              | How does your computer identify itself?                  |          1 |
| 38 | HOSTNAME                 |        0 | Hostname                 | Hostname of asset                                        |          1 |
| 39 | NODECLASS                |        0 | Nodeclass                | Nodeclass of asset                                       |          1 |
| 40 | POOL                     |        0 | Pool                     | Groups related assets spanning multiple functional roles |          1 |
| 41 | PRIMARY_ROLE             |        0 | Primary Role             | Primary functional role of asset                         |          1 |
| 42 | SECONDARY_ROLE           |        0 | Secondary Role           | Secondary functional role of asset                       |          1 |
| 43 | POWER_OUTLET             |       -1 | Power_outlet             | POWER_OUTLET                                             |          1 |
| 44 | IS_NODECLASS             |       -1 | Is_nodeclass             | IS_NODECLASS                                             |          1 |
| 45 | SUFFIX                   |       -1 | Suffix                   | SUFFIX                                                   |          1 |
| 46 | BUILD_CONTACT            |       -1 | Build_contact            | BUILD_CONTACT                                            |          1 |
| 47 | CONTACT                  |       -1 | Contact                  | CONTACT                                                  |          1 |
| 48 | CONTACT_NOTES            |       -1 | Contact_notes            | CONTACT_NOTES                                            |          1 |
| 49 | PASSWORD                 |       -1 | Password                 | password                                                 |          1 |
+----+--------------------------+----------+--------------------------+----------------------------------------------------------+------------+
```
