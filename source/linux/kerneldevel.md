Kernel Development
==================

This walks through a workflow for getting you up, running, and quickly iterating on kernel development.
In this guide I walk through how to use QEMU and build your own custom kernel for testing. As well as setting
up the OS that you will use for testing your kernel.

## Building the kernel

Before we get started we will want to build a fresh version of the kernel. I am just going to be building master
but you can always check out whatever tag or branch you wish to work on. As of writing the kernel version is 4.19-r6.

First you will want to clone the kernel. I found that github was very slow and that the official kernel github repo
was much faster. The UI for this can be found at https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git.

```
git clone git://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git
```

Now that we have the kernel we will want to setup our `.config` file. You can get some nice same defaults by running
the following commands.

```
make x86_64_defconfig
make kvmconfig
```

This will set you up with a nice sane default `.config` file additionally you might want to change the value of
`CONFIG_LOCALVERSION` to your name so that when you run `uname -a` on your test machine you can clearly see that
it's a kernel you built.

Now that we have everything configured as you like you will want to build the kernel. You can do so by running
`make -j8` where 8 for me is the number of cores on my computer. Additionally if you won't want to build all the
modules you can simply run `make bzImage` which will just produce the kernel.

Careful not to run `make install` as you likely don't want to be installing this on your laptop. A bad configuration
will cause your system to be unable to boot.

## Setting up QEMU (Debian)

You can do this a few different ways some more automated than others. I chose to do a manual install so I could have
a nice sane base environment for testing in but i'll go over a quick install a little farther down.

To start you will need to make a QEMU image. For this I used a qcow2 image that way I could layer changes on top of it
making it easy to throw away changes I make in an environment during testing.

```
qemu-img create -f qcow2 debian.qcow 2G
```

This will set you up with a raw image that you can then begin to install your OS on. For debian you can pull the needed
OS from [here][1].

```
wget http://cdimage.debian.org/cdimage/daily-builds/daily/arch-latest/i386/iso-cd/debian-testing-i386-netinst.iso
```

After you have this pulled down and your image created you can run the following to kick off the process.

```
qemu-system-x86_64 -hda debian.qcow -cdrom debian-testing-i386-netinst.iso -boot d -m 1024 -smp 2
```

To break this down `hda` is just telling the qemu file to use as a hard disk in this case the qcow2 image we created above.
`cdrom` sets the disk we want to use, `boot d` says boot right away and use the first cdrom we find. `-m` passes the amount
of memory in MB and `spm` is the number of cores to give to the machine. You may want to do more to speed the process up.

After qemu starts select expert install and do what you want. By default if you use the build in partitioning it will try to give
1G of your 2G disk away for swap. I would recommend manually setting up the partition and not using swap at all or significantly
reducing the amount you use.

After you have your system configured the way you want you can create another layer ontop of your qcow2 image. That means you can use
the base OS you made over and over and any changes will be written to another layer you can toss if you happen to mess it up.

```
qemu-img create -f qcow2 -o backing_file=debian.qcow test.qcow
```

## Setting up QEMU (Debian Quick)

If you want something very quick you can do this as well to setup debian.

```
IMG=qemu-image.img
DIR=mount-point.dir
qemu-img create $IMG 1g
sudo mkfs.ext2 $IMG
mkdir $DIR
sudo mount -o loop $IMG $DIR
sudo debootstrap --arch amd64 jessie $DIR
sudo umount $DIR
rmdir $DIR
```

This will give you a small debian install that you can use. However you do not get the niceness of image layering that you would in the
longer setup.

To test this out you can run `qemu-system-x86_64 -hda qemu-image.img`. This method is largely coppied from [here][2] which is another QEMU
kernel dev post. But I hope to provide you a little more depth here.

## Testing Your Kernel

Now that we have an image and a kernel that is built lets try things out. We are going to go the very simple route of booting with no initramfs
to start with just to make sure things are working.

```
qemu-system-x86_64 \
  -drive file=test.qcow2,if=virtio \
  -kernel ../linux/arch/x86/boot/bzImage \
  -append "console=ttyS0 root=/dev/vda1 ro" \
  -nographic \
  -m 2048 \
  -smp 4 \
  -cpu Haswell \
  --enable-kvm
```

Some things to note is that because qcow2 is not a natively supported format the kernel knows about we can't just pass it in with an `-hda` flag.
Doing so in this case will just end up with the kernel telling you it seens no drives at all. In this case we tell it to use the `test.qcow2` file
and the virtio driver which you can read more about [here][3]. Kernel is as it sounds the bzImage you produced from building your kernel. `append`
allows us to pass in all kinda of kernel params in this case I tell the kernel to output to the ttyS0 console and use the root device /dev/vda1.
ttyS0 is needed because i passed in nographic so that I can quickly test in a terminal without spawning a new window for the VM. `/dev/vda1` is used
because this is the name of the device when virtualized by virtio.

If this is your first time looking through qemu I highly encourage you to pull up the man page for the parts that aren't clear here and take a look.
They do a fairly good job of explaining all the flags and this is only touching the surface of what you can do with qemu.

## Setup networking

You can go about this several ways and this is mostly just here for myself but this was been the best way I have found so far to be able to interact
with my Guest VM from the host machine.

```
IFACE=wlp2s0
modprobe tun

ip tuntap add dev tap0 mode tap group kvm
ip link set dev tap0 up promisc on
ip addr add 0.0.0.0 dev tap0
ip link add br0 type bridge
ip link set br0 up
ip link set tap0 master br0
echo 0 > /sys/class/net/br0/bridge/stp_state
ip addr add 10.0.100.1/24 dev br0
sysctl net.ipv4.conf.tap0.proxy_arp=1
sysctl net.ipv4.conf.$IFACE.proxy_arp=1
sysctl net.ipv4.ip_forward=1
iptables -t nat -A POSTROUTING -o $IFACE -j MASQUERADE
iptables -A FORWARD -m state --state RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -i br0 -o $IFACE -j ACCEPT
```

This will setup two interface that you can then use to attach to your VMs when you start them up. You can do so with the addition of the following
to your qemu command `-nic tap,ifname=tap0,script=no,downscript=no`. If you use the same CIDR block that I do you would then have to manually set
your VM to use an IP in the 10.0.100.1/24 subnet. More on how this all work can be seen [here][4] and how to configure you interface can be seen
[here][5] depending on your OS. Additionally more about the `-nic` flag can be read [here][6].

## Building an initramfs (dracut)

Now that we have a simple working kernel lets see how we can build an initramfs which could be desired if you wanted to encrypt your drive for instance.

To start with we will need to produce all modules and install them to a new location being careful not to install them onto our actual system. The following
will take the modules we built above and install them to a folder I created called modules.

```
make modules_install INSTALL_MOD_PATH=/home/eatingthenight/code/kernel-devel/modules
```

Now that you have your modules installed you can use `dracut` to produce an initramfs image for you. You may need to alter the path slightly if you are
running a different version of the kernel.
```
sudo dracut \
  --kmoddir /home/eatingthenight/code/kernel-devel/modules/lib/modules/4.19.0-rc6+/ \
  initramfs
```

Now you can start your new OS up with the following commands.

```
sudo qemu-system-x86_64 \
  -drive file=test.qcow2,if=virtio \
  -kernel ../linux/arch/x86/boot/bzImage \
  -initrd ../linux/initramfs \
  -append "root=/dev/vda1 ro console=ttyS0" \
  -nographic \
  -m 2048 \
  -cpu Haswell \
  -smp 4
```

You will see the following line or one very similar fly by which confirms that you are using an initramfs now when booting.

```
[    0.373010] Unpacking initramfs...
```

Additionally here is some other helpful sources I found when debugging my initramfs.

1. If your initramfs is failing early on try reading though [here][7] or [here][8].
2. If you want to take a look at what is all installed in your initramfs try using [this program][10].
3. For dracut specific debugging steps check [here][11].
4. Look around in your initramfs with `gunzip -c initramfs | cpio -cvit` this will extract it to disk.
5. Pass `rdinit=/bin/bash` to the kernel during boot to drop to a shell inside your initramfs.
6. If you run into any very cryptic errors such as something like `trap invalid opcode` you likely need to use a different
CPU arch. I used Haswell to get around this error however if you have a different chip set you may need to change it.
7. IMPORTANT: if you set the kernel param `rdinit` to anything outside of `/init` all of the `rd.*` commands from 
`man dracut.cmdline` will no longer work since that init script looks for them. For instance if you switch your 
rdinit to `/sbin/init` it will happily ignore all of your rd commands. More can be seen [here][9]


[1]: http://cdimage.debian.org/cdimage/daily-builds/daily/arch-latest/i386/iso-cd/
[2]: https://www.collabora.com/news-and-blog/blog/2017/01/16/setting-up-qemu-kvm-for-kernel-development/
[3]: https://wiki.libvirt.org/page/Virtio
[4]: https://wiki.gentoo.org/wiki/QEMU#Networking
[5]: https://wiki.debian.org/NetworkConfiguration
[6]: https://www.qemu.org/2018/05/31/nic-parameter/
[7]: https://wiki.debian.org/InitramfsDebug
[8]: https://www.askapache.com/linux/linux-debugging/
[9]: https://unix.stackexchange.com/questions/30414/what-can-make-passing-init-path-to-program-to-the-kernel-not-start-program-as-i
[10]: http://manpages.ubuntu.com/manpages/xenial/man8/lsinitramfs.8.html
[11]: https://fedoraproject.org/wiki/How_to_debug_Dracut_problems

[20]: https://www.youtube.com/watch?v=PBY9l97-lto
[21]: https://www.kernel.org/doc/html/v4.14/admin-guide/kernel-parameters.html
[22]: https://www.berrange.com/posts/2018/06/29/cpu-model-configuration-for-qemu-kvm-on-x86-hosts/
