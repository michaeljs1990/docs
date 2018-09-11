ioctl in Ruby
=============

I'll add a more detailed breakdown of this code but for now here it is.

```
#!/usr/bin/env ruby

require 'socket'


SIOCETHTOOL = 0x8946
ETHTOOL_GRINGPARAM = 0x00000010
ETHTOOL_SRINGPARAM = 0x00000011

sock = UDPSocket.new

# https://github.com/torvalds/linux/blob/9a76aba02a37718242d7cdc294f0a3901928aa57/include/uapi/linux/ethtool.h#L487
ringparam_struct = [ETHTOOL_GRINGPARAM].pack "Vx32"
ifreq = ["enp0s31f6", ringparam_struct].pack "a16P36"
sock.ioctl(SIOCETHTOOL, ifreq)

=begin
  struct ethtool_ringparam - RX/TX ring parameters
  @cmd: Command number = %ETHTOOL_GRINGPARAM or %ETHTOOL_SRINGPARAM
  @rx_max_pending: Maximum supported number of pending entries per
       RX ring.  Read-only.
  @rx_mini_max_pending: Maximum supported number of pending entries
       per RX mini ring.  Read-only.
  @rx_jumbo_max_pending: Maximum supported number of pending entries
       per RX jumbo ring.  Read-only.
  @tx_max_pending: Maximum supported number of pending entries per
       TX ring.  Read-only.
  @rx_pending: Current maximum number of pending entries per RX ring
  @rx_mini_pending: Current maximum number of pending entries per RX
       mini ring
  @rx_jumbo_pending: Current maximum number of pending entries per RX
       jumbo ring
  @tx_pending: Current maximum supported number of pending entries
       per TX ring
=end
rv =  ringparam_struct.unpack "V9"

# pop off the GRINGPARAM cmd from the struct
_, *tail = rv

# To see what is all changable by the user you can read the source code here
# https://github.com/torvalds/linux/blob/9a76aba02a37718242d7cdc294f0a3901928aa57/drivers/net/ethernet/freescale/gianfar_ethtool.c#L419
ifreq = tail.insert(0, ETHTOOL_SRINGPARAM)
ifreq[5] = 512
ifreq[8] = 512

# Now lets set the data
ringparam_struct = ifreq.pack "V9"
ifreq = ["enp0s31f6", ringparam_struct].pack "a16P36"
sock.ioctl(SIOCETHTOOL, ifreq)
```
