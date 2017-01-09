Cisco Catalyst 3560G Config
===========================

This document covers a range of different configurations needed and some not needed for getting my switch working.
Most of this is likely wrong or has a better way of doing it i'm still learning. I will try and come back and clean
this up later so I don't leave crap info sitting around forever.

# Set Cisco Banner (motd)

This will set the banner that is shown when first loggin into your switch.

```
enable
configure terminal
banner motd [
end
copy running-config startup-config
```

You can now type in whatever you want and when you are finished type `[` and press enter.

# Enable LLDP

Enable connected deviced to retrieve attributes about the network. More about LLDP can be seen in the first paragraph
of this [link](http://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst3560/software/release/12-2_55_se/configuration/guide/3560_scg/swlldp.html).

```
enable
configure terminal
lldp run
end
copy running-config startup-config
```
