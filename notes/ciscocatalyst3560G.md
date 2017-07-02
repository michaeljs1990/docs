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

# Set Hostname & DNS

```
enable
configure terminal
hostname <hostname>
ip domain-name <domain-root>
ip name-server <ip>
end
copy running-config startup-config
```

Hostname in my case would be something in the convention of `<switch model>.<rack name>.<pod number>.<datacenter>` this could
look like `c3650g.tor-a1.pod-01.pit-shd-1` which stands for...

```
c3650g    - Cisco 3650G switch model not required you may have a better way to dermine this in your environment
tor-a1    - This is a top of rack switch called a1 so it's easy to locate inside the datacenter
pod-01    - This is the first pod in the datacenter normally logically divided by OOB networks or IP space.
pit-shd-1 - This is a datacenter some place in pittsburgh
```

The domain-name is something like website.com if you happened to own that domain and used it for all of the computers located
inside this facility. You will likely have a dedicated DNS server that serves this domain name in your datacenter.

The name-server is just the ip, hopefully you can enter more than one although i'm not sure how that configuration would look
at this time since I only have one setup for testing.

# Setup SSH

Before you start make sure you have done all configuration in the "Set Hostname & DNS" section.

```
enable
configure terminal
crypto key generate rsa
end

configure terminal
line vty 0 4
transport input ssh
login local
exit

configure terminal
username <username> password <password>

copy running-config startup-config
```

You should now be able to log in with the username and password you have set. You can set multiple accounts as well. I imagine 
you can obtain the private key as well and use that but haven't looked into it yet. To see what your current settings are you
can run the following two commands.

```
show ip ssh
show ssh
```

# List interfaces

Not much going on here outside of remembering that if the node connected to the interface is offline it will report as a-100
instead of whatever the true interface speed is.

```
show interfaces status
```

If you want to filter for connected only interfaces on devices with a large number of ports..

```
show interfaces status | include connected
```

# VLANs

General overview of working with VLAN for noobs like me.

To get an overview of the current/default VLANs on your switch and the ports connected to them run 

```
show vlan
```

To create a new VLAN run the following commands on your switch. VLAN 2 is used in this example
but you can use any VLAN that you wish. Some things to note is that `no shutdown` is used to
bring the interface up since by default it is set to down. The default-gateway is not needed 
however I explicitly state it here in case you want to use a different gateway for the specific
VLAN than the one used for all interfaces on the router.

```
enable
configure terminal

vlan 2
name ToR
exit

interface Vlan 2
ip address 10.0.0.1 255.255.255.0
no shutdown
exit

copy running-config startup-config
```

The above will create the VLAN interface which you can now start assigning to ports. You can use
the show interfaces status to get a list of ports that are currently active.

```
enable
configure terminal

interface GigabitEthernet 0/33
switchport access vlan 2
exit

copy running-config startup-config
```

If you did everything correct the Vlan you configured should say `Vlan2 is up, line protocol is up`
when you check the output of `show interfaces`. `show interfaces brief` should list all ports using
the new VLAN as well.

To test you can use the cisco ping command `ping 192.168.1.1 source 10.0.0.1` to ensure that your
routing is setup properly.
