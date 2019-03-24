BGP
===

This document covers some basic documentation on using BGP in
a number of different ways.

## ECMP (Equal Cost Multi-Path Routing / Anycast)

In this example I am going to cover [ECMP][1] setup using the quagga bgpd
daemon. ECMP will allow us to have multiple routes to the same CIDR block.
In most cases something like a /24 is handed out for a specific application
such as NTP. The idea of ECMP being that you could have a number of servers
sitting behind an IP such as 192.168.60.1 and as requests come in they are
sent to multiple physical hosts. If one fails it can quickly be removed and
any applications communicating with your service will be largely unaffected
as the update can happen in a matter of seconds.

Compare this to load balancing behind something like DNS where a record could
be cached in any number of locations causing consumers to continue to try
reading from a downed service. Additionally for services that are huge pains
to move such as DNS if you need to physically move your server and in doing so
give it a new IP you no longer have to care. Provided you already had ECMP in
the way I am about to show you setup you could continue broadcasting that you
own 192.168.60.1 at the new location.

With that being said lets take a look at a very simple setup using quagga and
the bgpd daemon installed on three different servers. To start with you can
run the following commands on a debian based OS.

```
apt-get install -y quagga

touch /etc/quagga/bgpd.conf
touch /etc/quagga/zebra.conf
touch /etc/quagga/vtysh.conf
```

This will create the needed config files for you to start up your daemons. After
you have done this we will now setup the two nodes that you can think of as being
nodes that will contain our DNS service. On each one run the following commands

```
ip addr add 192.168.60.1/24 dev lo
```

This will add a new CIRD block to the loopback interface so we can now ping and
run a service on 192.168.60.1. Next we will configure bgpd on the two DNS servers.

```
eth1ip=$(ip addr show eth1 | grep -Po 'inet \K[\d.]+')

cat > /etc/quagga/bgpd.conf <<EOF
!
! Zebra configuration saved from vty
!   2018/11/18 01:55:29
!
!
router bgp 200
 bgp router-id 10.10.12.102
 redistribute connected
 neighbor 10.10.12.101 remote-as 100
! no auto-summary
!
 address-family ipv6
 exit-address-family
 exit
!
line vty
!
EOF
```

This will create a bgp route for asn 200 and set the router-id which is used by
other routers to identify your client. You can use anything but your public IP
or the IP that other bgp daemons will reach you at is commonly used. The neighbor
command is specifying the IP at which we will establish a bgp session and the asn
that we expect it to have which is 100 in this case.

To dig into more of the commands and also change this to your liking check out the 
quagga [docs][2]. Now start up the bgpd and zebra daemon.

```
systemctl start zebra
systemctl start bgpd
```

After setting up the above on both DNS nodes lets now setup the node that will be doing
the actual ECMP routing. In our case it's just another server running qugga but in production
this is likely a router such as cisco or juniper.

```
cat > /etc/quagga/bgpd.conf <<EOF
!
! Zebra configuration saved from vty
!   2018/11/16 05:15:45
!
router bgp 100
 bgp router-id 10.10.12.101
 neighbor 10.10.12.102 remote-as 200
 neighbor 10.10.12.102 route-map r60 in
 neighbor 10.10.12.103 remote-as 200
 neighbor 10.10.12.103 route-map r60 in
 maximum-paths 2
!
 address-family ipv6
 exit-address-family
 exit
!
access-list 10 permit 192.168.60.0 0.0.0.255
!
route-map r60 permit 10
 match ip address 10
!
line vty
!
EOF
```

This case is very similar to the last config except we now specify the names of the two BGP
servers. We additionally add a route-map to each of the neighbors so that we only get routes
that we are expecting from them. Without this they would be allowed to send any
routes they want and I believe they could even overwrite public routes. To break down
how this is working first we define an access-list with the name of 10. This could
have anything in it's place. we permit 192.168.60.0/24 and add it to a route-map
called r60 that will permit the route to be redistributed. Additionally the route-map
is given sequence number ten which if multiple route-maps are use decides when the map
is evaluated lowest comes first. The match ip address 10 is saying that we will match
against the ip address in access-list 10. 

Additionally route-maps are a bit tricky and if you are interested in really understanding 
them [read this][7].

The maximum-paths 2 is where we are allowing the ECMP magic to happen. With out this quagga
will only allow one path to be entered into the routing table. If you want to have more paths
increase this number. 

When the BGP session is established they will broadcast what routes they know about
which in our case is the 192.168.60.0/24 which we created above. Start zebra and bgpd and you
should see an entry in your routing table now that looks like this.

```
root@debian1:~# ip route
default via 10.0.2.2 dev eth0 
10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 
10.10.12.0/24 dev eth1 proto kernel scope link src 10.10.12.101 
192.168.60.0/24 proto zebra metric 20 
        nexthop via 10.10.12.102  dev eth1 weight 1
        nexthop via 10.10.12.103  dev eth1 weight 1
```

Try listening on a port with netcat and then telnet to 192.168.60.1 swice to see what happens.
This is a very basic example of what you can do with BGP and ECMP and I also have a vagrantfile
for those who want to easily play around with this [here][3].

## Route distribution

You can pick what route you want to send to hosts in multiple ways two of the most useful are
manually setting it with network or just passing everything on with redistribute.

```
router bgp 100
  redistribute connected
```

```
router bgp 100
  network 192.168.60.0 mask 255.255.255.0
```

## Show neighbors routes

If you ever are having issues getting the routes to show up you can see what the following command
outputs by hopping on a neighbor you are connecting to and running the following. This will output
a bunch of information about the state of the routes as well as what ones it's getting.

```
show ip bgp neighbors 10.10.12.103 routes
```

## Show route table for bgp session

```
show ip bgp
```

# Random

[General BGP Intro][8]
[Getting your own ASN and IP Space][9]
[Simple BGP tutorial][10]

[1]: https://en.wikipedia.org/wiki/Equal-cost_multi-path_routing
[2]: https://www.quagga.net/docs/quagga.html#BGP
[3]: https://github.com/michaeljs1990/testlabs/tree/master/anycast
[4]: https://serverfault.com/questions/696675/multipath-routing-in-post-3-6-kernels
[5]: https://www.noction.com/blog/equal-cost-multipath-ecmp
[6]: http://highscalability.com/blog/2014/8/4/tumblr-hashing-your-way-to-handling-23000-blog-requests-per.html
[7]: https://www.cisco.com/c/en/us/td/docs/security/asa/asa84/configuration/guide/asa_84_cli_config/route_maps.pdf
[8]: http://packetfire.org/post/intro-to-bgp/
[9]: https://labs.ripe.net/Members/samir_jafferali/build-your-own-anycast-network-in-nine-steps
[10]: http://www.m0rd0r.eu/simple-bgp-peering-with-quagga-ex-zebra/
