Junos OS
========

Basic commands for interacting with juniper networking gear.

## Enter configuration mode

To enter the cli and the configuration mode run the following two commands.

```
cli
configure
```

Once you have finished making changes you will need to run the following so that
they actually take effect.

```
commit
```

## Interface naming information

[This page][1] gives you a good overview of all the different interface names and what they mean.

## Helpful commands

If you need to change your hostname you can run the following.

```
# enter configuration mode
set system host-name yourhostname
```

Get a quick look at all the networking interfaces

```
show interfaces terse
```

If you are in configuration mode and want to run a command such as show interfaces you
can do this by appending `run` to it.

```
run show interfaces terse
```

## Deleting configuration

Deleting configurations is easy and following the following syntax.

```
delete thing to remove
```

For instance if we have an interface that we no longer called ge-0/0/0 you
would run the following command.

```
delete interfaces ge-0/0/0
```

## Configure a new interface

```
set interfaces xe-0/0/0 unit 0 family inet address 192.168.2.1/24
```

If you run into an error about the address not being compatible with DHCP you will
need to remove dhcp from the configuration like so.

```
delete interfaces xe-0/0/0 unit 0 family inet dhcp
```

[1]: https://www.juniper.net/documentation/en_US/junos/topics/concept/interfaces-interface-naming-overview.html
