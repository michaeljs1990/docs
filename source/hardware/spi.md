Bus Pirate
==========

If you like me are brand new to SPI and interfacing with your MOBO and chips
on it in general you will likely want to take a read through the following
articles before starting.

[Serial Communication][1]
[SPI Interface][2]
[Pulse Width Modulation][3]
[Multimeter Voltage Basics][5]
[Multimeter Current Basics][6]
[Bus Pirate v3 SPI + flashrom reading WD800JD EEPROM][4]

# Pins

This is some common pins that you might run into when looking at datasheets as well as
a walkthrough of how to hook up the bus pirate with a datasheet found online.

|Pins       |Description                                                    |
|-----------|---------------------------------------------------------------|
|GND        |Ground                                                         |
|+3V3       |3.3 Volts                                                      |
|+5V        |5.0 Volts                                                      |
|ADC        | |
|VPU        | |
|AUX        |Auxiliary                                                      |
|CLK        |Clock                                                          |
|MOSI,SI    |Master Out Slave In. Sometimes shown as SI on the IC datasheet |
|MISO,SO    |Master In Slave Out. Sometimes shown as SO on the IC datasheet |
|CS,CE      |Chip Select / Chip Enable                                      |
|POK,PWR_GD |Power-OK or Power Good. Normally hooked up to +3.3V or +5V     |
|EN         |Enable Pin                                                     |

# Using flashrom with the buspirate

Lets first start my find a motherboard. Any board should do provided you haven't pulled the bios
chip off of it yet. Depending on your board you may have to look around a bit to find it but on
most boards it should be labeled with something like U_SPI_BIOS. If you can't find it or aren't sure
look up the manual for your system and look for a volatility chart or appendix that has mention of
"System BIOS SPI Flash".

Once you have located the chip on your board you will want to find what specific chip it is so you
can look up the datasheet for it. For my chip this is [MX25L3205D][7] however on some chips this is
super hard to read so I picked up one of [these][8] to help me out. How that we have the name of the
chip we can go googling around for the datasheet. For this chip I was led to this [datasheet][9].

This has a bunch of useful information but for now all I care about is the "8-PIN SOP" listed under
chip configurations. This tells us how we need to wire the bus pirate to connect to this chip. From
the diagram for the chip we are given...

```
CS  |----| VCC
SO  |    | HOLD
WP  |    | SCLK
GND |----| SI
```
How this will look when connected to the bus pirate is shown in the following table. The slots left empty
do not need to be hooked up.

|CHIP|Bus Pirate|
|----|----------|
|CS  | CS       |
|SO  | MISO     |
|WP  |          |
|GND | GND      |
|VCC | 3.3V     |
|HOLD|          |
|SCLK| CLK      |
|SI  | MOSI     |

Once you have everything connected up you will need to connect your chip up properly. You should see a notch
in on one side of it or a tiny dot in one of the corners. After it's hooked up properly you can run the flashrom
program. The first time you run flashrom it will look something like this.

```
sudo ./flashrom --programmer buspirate_spi:dev=/dev/ttyUSB0
```

If you are lucky it will spit out a list of chips that it thinks might be hooked up. This was not the case for me
however and I had to list all the chips and enter the chip I knew it was already.

```
# List all the chips and look for the one you need
./flashrom --list-supported
# Tell flashrom the chip you want to look for
sudo ./flashrom --programmer buspirate_spi:dev=/dev/ttyUSB0 -c "MX25L3206E/MX25L3208E"
```
Now you can read from the chip to pull off default bios code.

```
sudo ./flashrom --programmer buspirate_spi:dev=/dev/ttyUSB0 -c "MX25L3206E/MX25L3208E" -r MX25L3206E.bin
```

You will want to do this twice and then run `md5sum` on both of the reads to ensure that you have gotten a good one.
If you have a bad chip the sum will be different every time and you may need to buy a new one.

# Lingo

nybble - half a byte
cmos - bettery powered dynamic ram for your bios (in most cases).

[1]: https://learn.sparkfun.com/tutorials/serial-communication
[2]: https://learn.sparkfun.com/tutorials/serial-peripheral-interface-spi
[3]: https://learn.sparkfun.com/tutorials/pulse-width-modulation
[4]: https://www.youtube.com/watch?v=DH_cIrPxn5c
[5]: https://www.youtube.com/watch?v=ZBbgiBU96mM
[6]: https://www.youtube.com/watch?v=EVFkKBFJsZg
[7]: https://datasheet.octopart.com/MX25L3205DM2I-12G-Macronix-datasheet-8325093.pdf
[8]: https://www.amazon.com/gp/product/B079BQSPDZ/ref=oh_aui_detailpage_o00_s00?ie=UTF8&psc=1
[9]: https://datasheet.octopart.com/MX25L3205DM2I-12G-Macronix-datasheet-8325093.pdf
