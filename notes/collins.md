[Collins](http://tumblr.github.io/collins/)
=======

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
