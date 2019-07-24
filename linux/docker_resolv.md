## Taking a look at dockers resolv.conf setup

Recently I ran into an issue where I had multiple long running docker containers on a
host that had started up pointing to my internal DNS server. When I was in the process
of changing this over I updated resolv.conf on my system and switched the IP. Everything
went fine except my docker containers were now unable to resolve any hosts.

### What happened?

At the most basic level even though docker is using my /etc/resolv.conf it is not directly
mounting it into the container but instead making a copy of the file before it does.
To see what this looks like we can start up a small container and poke around.

```
$ docker run -it alpine sh
/ # df -h | grep resolv
/dev/root               233.6G    124.0G     97.7G  56% /etc/resolv.conf
/ # cat /proc/self/mountinfo | grep resolv
131 110 259:3 /var/lib/docker/containers/e696563e1ffb2c6443cc9505cd52850b0686729ff6f829826495d56f5bc455a4/resolv.conf /etc/resolv.conf rw,noatime - ext4 /dev/root rw,discard
```

Here we start up an alpine container look at the mounted file systems and find resolv.conf.
Since it's mounting in a specific file we know it's a bind mount and can find out more info
from `/proc/self/mountinfo`. This shows up the path to resolv.conf on the host system.

```
$ ls -lah /var/lib/docker/containers/e696563e1ffb2c6443cc9505cd52850b0686729ff6f829826495d56f5bc455a4/resolv.conf
-rw-r--r-- 1 root root 118 Jul 24 00:35 /var/lib/docker/containers/e696563e1ffb2c6443cc9505cd52850b0686729ff6f829826495d56f5bc455a4/resolv.conf
```

Checking that file it's not pointing at the host system resolv.conf but is instead a copy of it.
Additionally since it shows 1 after the file permissions we know that it is not a hard link.

Taking a look at the [documentation][1] also gives a little insight to this as well.

> Note: The file change notifier relies on the Linux kernel’s inotify feature.
> Because this feature is currently incompatible with the overlay filesystem driver,
> a Docker daemon using “overlay” will not be able to take advantage of the
> /etc/resolv.conf auto-update feature.

I am not sure why the docker daemon can't support this when using the overlay driver because
unless `/etc/resolv.conf` on the host system is running on an overlay this shouldn't matter at
all. I think this documentation is partly wrong. In addition to get the update into the container
you will have to bounce the container since no action will be taken on a [running][2] ones resolv.conf.
This is a sane default but it would be a nice flag to add to the daemon to say I don't care
just change it in the running containers as well.

### So whats the fix?

The real fix would be to modify this [patch][2] with the ability to configure when resolv.conf
gets updated and convince the docker project it's the correct change. 

The easy fix if to just do a find for resolv.conf under your docker containers directory. Likely
`/var/lib/docker/containers` and copy your resolv.conf to them with something like.

```
$ find /var/lib/docker/containers -name resolv.conf -exec cp /etc/resolv.conf {} +
```

1: https://docs.docker.com/v17.09/engine/userguide/networking/default_network/configure-dns/
2: https://github.com/moby/moby/pull/9648/files#diff-6ce9e79ddb91a3f06352abe1f2c72ecbR1033
