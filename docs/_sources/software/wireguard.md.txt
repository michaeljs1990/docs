# WireGuard

WireGuard is an awesome new in kernel based VPN. To learn more about is checkout out
the (website)[1] which has good information to get you started.

Here I am just going to go over a basic setup that will tunnel traffic from a server
at my house (which I will refer to as the client) to a server run by some cloud provider
(which I will refer to as the server).

## Install

First you will need to install WireGuard on your OS. They have a good list here that
should include how to install WireGuard for most common OS https://www.wireguard.com/install/.

## Server Setup

First me need to create the keys that will be used to communiate with our WireGuard install.
You can generate your keys using the `wg` command.

```
wg genkey > private_key
cat private_key | wg pubkey > pub_key
```

Next we will create a config file that is used to configure our new wg0 interface that all
out traffic will pass through. This can also be done with `ip` and `wg` commands but I won't
conver that here.

```
[Interface]
Address = 172.22.22.1
SaveConfig = true
PrivateKey = <server_private_key>
ListenPort = 51820

[Peer]
PublicKey = <client_public_key>
AllowedIPs = 172.22.22.0/24
```

To test out this config run  `wg-quick up wg0` and watch the output.

To break down what is going on we are creating an interface and giving it an IP of 172.22.22.1
giving it the private key that clients will need the public key of and setting a specific Port
for listening on. You can read about all the options available (here)[2].

The Peer section is setting what clients are allowed to connect to us. If you don't care where
the connections come from you can set AllowedIPs to 0.0.0.0/0.

Now that the servers can talk to each other you may want to do some additional setup so that you
can route additional traffic through this server and use it as a VPN.

First we enable IP forwarding If you want to make it persist reboots you need to modify `/etc/sysctl.conf`
and add the following line `net.ipv4.ip_forward=1`.

```
echo 1 > /proc/sys/net/ipv4/ip_forward
```


## Client Setup

We will need to first generate a public key and private key for our client. This is something
you will need to do a lot since all clients and servers need to do this step.

```
wg genkey > private_key
cat private_key | wg pubkey > pub_key
```

Next we create the client configuration file. Some of this information you can leave blank
and fill in after the server is setup. One thing to note that AllowedIPs when using the
`wg-quick` command will use that CIDR to add a route to your table. This means that if you
are connecting to a remote server and forward everything to 0.0.0.0/0 you may have a bad
day. It could be what you want but i'll leave that up to you.

```
[Interface]
Address = 172.22.22.2
PrivateKey = <client_private_key>

[Peer]
PublicKey = <server_public_key>
Endpoint = <server_ip>:51820
AllowedIPs = 172.22.22.0/24
PersistentKeepalive = 10
```

To test out this config run  `wg-quick up wg0` and watch the output.

To make sense of this we need to look at the server setup. We are saying use the address of
172.22.22.2 because we need an IP in the 172.22.22.0/24 range and 172.22.22.1 is taken by the
server config. Additionally we pull the port from the server config as well and the IP just
needs to be a publicly reachable IP or any IP reachable to your server if you are doing this
on an internal network. PersistentKeepalive works nice when you have a client that is behind
a firewall but want the connection to remain open even if the client hasn't sent us anything
in a while.

To test that everything works you can `ping 172.22.22.1` and then on the server run `wg` you
should see something like "latest handshake" that happened in the past few seconds. If you see
that you are able to talk with the server.

## Common Operations

Removing a peer:

If you have a peer that you would like to remove you can run the following command. You can find
the public key needed for this command by running `wg`.

```
wg set wg0 peer <public_key> remove
```

Remove Interface:

This is just the same as any other interface so you can remove it using the following.

```
ip link set wg0 down
ip link delete wg0
```


[1]: https://www.wireguard.com/quickstart
[2]: https://git.zx2c4.com/WireGuard/about/src/tools/man/wg-quick.8
