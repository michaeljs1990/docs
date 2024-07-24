Cisco Catalyst 3560G Config
===========================

This document covers a range of different configurations needed and some not needed for getting my switch working.
Most of this is likely wrong or has a better way of doing it i'm still learning. I will try and come back and clean
this up later so I don't leave crap info sitting around forever.

## Set Cisco Banner (motd)

This will set the banner that is shown when first loggin into your switch.

```
enable
configure terminal
banner motd [
end
copy running-config startup-config
```

You can now type in whatever you want and when you are finished type `[` and press enter.

## Enable LLDP

Enable connected deviced to retrieve attributes about the network. More about LLDP can be seen in the first paragraph
of this [link](http://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst3560/software/release/12-2_55_se/configuration/guide/3560_scg/swlldp.html).

```
enable
configure terminal
lldp run
end
copy running-config startup-config
```

## Set Hostname & DNS

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

## Setup SSH

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

## List interfaces

Not much going on here outside of remembering that if the node connected to the interface is offline it will report as a-100
instead of whatever the true interface speed is.

```
show interfaces status
```

If you want to filter for connected only interfaces on devices with a large number of ports..

```
show interfaces status | include connected
```

## VLANs

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

Note you will still have to setup a static route so that in this case the router handling
`192.168.1.1` which in my case is a static route 10.0.0.1 -> 192.168.1.8 where 192.168.1.8
is the ip for VLAN 1 on my switch.

## Configure DHCP 

Very similiar to the steps for settingup up vlans for the most part you can configure your switch
port to use DHCP to get it's IP address.

```
enable
configure terminal

interface GigabitEthernet 0/1
ip address dhcp
exit

copy running-config startup-config
```

This will allow you to let an upstream router acting as the DHCP set your IP instead of having to change it by hand
when you need to make changes to your router.

## DHCP Forwarding 

If you have a DHCP server on another broadcast domain you can set the DHCP server the ToR switch
should forward to using the following. In this example 192.168.1.5 is my DHCP server.
```
enable
configure terminal

interface Vlan 2
ip helper-address 192.168.1.5
exit

copy running-config startup-config
```

## Random

To stop Cisco from trying to do a dns lookup when you type in an invalid command
you can issue the following.

```
no ip domain lookup
```

If you would like to not be blasted with logging while you are typing..

```
no logging console
```

Get the IP of connected devices that are in your ARP table. This can be useful to check if a connected
device has a static IP set that is not routable on the network.

```
sh ip arp
```

## Why doesn't my SFP work?

This is for a number of reasons. Maybe you didn't read the full spec for the hardware you bought. Maybe
the cable is bad. Whatever the reason this might help you figure it out.

First take a look at

```
enable
show interfaces
```

The output might look like this or maybe you have an entirely different problem which I can't help with.
If your output does look like this you have Gi0/51 and Gi0/52 which are not plugged in at all or the
SFP is bad and not registering.

For Gi0/49 and Gi0/50 it is showing err-disabled which means it connected and then cisco instantly disabled
the interface and will not send anything over it.

```
Gi0/49                       err-disabled 1            auto   auto unknown
Gi0/50                       err-disabled 1            auto   auto unknown
Gi0/51                       notconnect   1            auto   auto Not Present
Gi0/52                       notconnect   1            auto   auto Not Present
```

To dive a little deeper into what is going on you can run the following command.

```
show log
```

For me this shows the following. Which which is a CRC error when validating the hardware.
The following is from ciscos website.

> Note: Cisco-approved Small Form-factor Pluggable (SFP) modules have a serial EEPROM that
> contains the module serial number, the vendor name and ID, a unique security code, and CRC.
> When an SFP module is inserted in the switch, the switch software reads the EEPROM to verify
> the serial number, vendor name and vendor ID, and recompute the security code and CRC. If the
> serial number, the vendor name or vendor ID, the security code, or CRC is invalid, the software
> generates the security error message and possibly places the interface in an error-disabled state.

```
Jul 23 19:51:36.585: %GBIC_SECURITY_CRYPT-4-VN_DATA_CRC_ERROR: GBIC in port Gi0/52 has bad crc
Jul 23 19:51:36.585: %PM-4-ERR_DISABLE: gbic-invalid error detected on Gi0/52, putting Gi0/52 in err-disable state
Jul 23 19:51:44.336: %GBIC_SECURITY_CRYPT-4-VN_DATA_CRC_ERROR: GBIC in port Gi0/51 has bad crc
Jul 23 19:51:44.336: %PM-4-ERR_DISABLE: gbic-invalid error detected on Gi0/51, putting Gi0/51 in err-disable state
Jul 23 19:59:46.466: %GBIC_SECURITY_CRYPT-4-VN_DATA_CRC_ERROR: GBIC in port Gi0/52 has bad crc
Jul 23 20:00:12.975: %GBIC_SECURITY_CRYPT-4-VN_DATA_CRC_ERROR: GBIC in port Gi0/50 has bad crc
Jul 23 20:00:35.440: %GBIC_SECURITY_CRYPT-4-VN_DATA_CRC_ERROR: GBIC in port Gi0/51 has bad crc
```

To remedy the error cisco recommends the following.

> In order to resolve this, verify that the software release that runs on the system supports the
> GBIC module. If the GBIC module is newer, a system software upgrade is possibly required.

Looking at the spec sheet for my hardware I can see that it only support a SFP (1GBi) and not a 
SFP+ (10GBi). So it's likely I need to buy a much older SFP cable that is supported by this catalyst.
However because it's only a SFP that supports 1GBi it's not even worth it since ethernet supports
the same thing. I'll likely upgrade from this system that was originally released in `10-FEB-2004`
but for now this 20 year old hardware works pretty well.

## What kind of system do I have

If you aren't sure what kind of hardware you are running on you can run the following command to get an
identifier.

```
show inventory
```

## Get System Temp Info

If you are looking for info on if your system is overheating you can use

```
show env all
```

To return information about the temp, system fan, and temp thresholds.

[1]: https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst3560/software/release/12-2_55_se/configuration/guide/3560_scg/swiprout.html
