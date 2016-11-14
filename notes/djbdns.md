djbdns (How to setup your own DNS)
======

The following will give you a good background on how dns works. At the most basic level setting up your own DNS requries two different programs. a DNS cache and a DNS server. The cache will handle lookups and will delegate to your internal DNS server when needed. For all other queries it will use of the root DNS servers.

* http://cr.yp.to/djbdns/intro-dns.html

For installing the dns server you can follow this link.

* http://askubuntu.com/questions/53352/how-do-i-setup-an-authoritative-name-server-with-tinydns

However by itself this is worthless since any computer set to use this dns will fail on any domain not in it's data file. Also note that in this setup we are going to put the dnscache and dnsserver on the same computer so you will want to use 127.0.0.1 for the ip when setting up tinydns. This is because dns cache and tiny dns run on the same port however you only need your dns server available to the dns cache since users should not be directly talking to your dns server.

For setting up dns cache it's the exact same as the link above except you should make a dnscache user and create the service using dnscache-conf. If you followed the steps above correctly all the required binaries should be available on your system already. The following link has a small section on setting up dnscache if you need further help. I also explains how to allow others on the network to query your dns server.

* http://www.fredshack.com/docs/djbdns.html

Extra dnscache reading

* http://thedjbway.b0llix.net/djbdns/dnscache.html

Both of these services are run with daemontools so starting and stopping the service looks like this in my setup.

```
# start service
sudo svc -u /etc/service/tinydns
# stop service
sudo svc -d /etc/service/tinydns
# restart
sudo svc -t /etc/service/tinydns
```

Read over `man svc` for more information. The following link also contains an in depth look.

* https://cr.yp.to/daemontools.html

When troubleshooting issues the following command came in very useful. If you are having issues first stop dns cache and the dns server. Manually start the programs with the ./run binary located in the root of tinydns or dnscache so you can see all logs in real time. If you are having issues with anything resolving you can use the following.

```
dnsqr any google.com
```

This will query the dns server set in resolve.conf which should point to the server your dns cache is setup on. If nothing is returned you may need to update your list of root DNS servers or the machine your cache is running on may not have an internet conection. If you are able to query sites such as google.com but can't get your internal dns to resolve make sure that in your dns cache root directory you have a file such as example.local with an ip of 127.0.0.1. This will ensure that when a query for something.example.local comes in your local dns server is used instead of the root DNS servers.

If you are having trouble getting tinydns to resolve you can verify that it is properly working by cd'ing into the directory that contains your data.cdb file and running `tinydns-get any your.domain`. If nothing is returned you make need to run `make` again to regenerate your db. The link below contains more on tinydns-data.

* https://cr.yp.to/djbdns/tinydns-data.html

Extras

1. http://www.mn-linux.org/meetings/pastnotes/djbdns.pdf
2. http://www.troubleshooters.com/linux/djbdns/djbdns_tshoot.htm
3. https://cr.yp.to/djbdns/run-server.html
