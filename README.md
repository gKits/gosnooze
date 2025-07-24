<!-- "TODO: Create documentation" -->

<div align="center">
  <a href="https://github.com/gkits/gosnooze">
    <img src="assets/gopher.png" alt="Logo" width="180" height="180">
  </a>
  <h1 align="center">GoSnooze</h1>
  <p align="center">
    A hackable alarm clock using TinyGo.
  </p>
</div>

## About

## Getting started

### Bill of Materials

- Raspberry Pico W
- DS3231 real time clock
- HD44780 16x2 LCD display + I2C converter (like [this](https://www.az-delivery.de/en/products/lcd-display-16x2-mit-blauem-hintergrund-und-i2c-converter-bundle?variant=8192126353504&utm_source=google&utm_medium=cpc&utm_campaign=16964979024&utm_content=166733588295&utm_term=&gad_source=1&gad_campaignid=16964979024&gclid=Cj0KCQjws4fEBhD-ARIsACC3d28ngqHbnuj5FSCwEo_tkcEt6AOFJBR66nADGQOAEowSFmo4Xrp6RZEaAtgXEALw_wcB))
- ~~3D printed case~~ [not yet]

### Assembly

[not yet]

### Flashing

1. Clone this repo and move into the directory.
   ```sh
   git clone https://github.com/gkits/gosnooze
   cd gosnooze
   ```

2. Start the Pico W in boot selection mode by holding down the `BOOTSEL` button while connecting it
   to your computer via USB.

3. Mount the Pico W to your filesystem.
    1. Find the correct block device by running `lsblk`. The correct device should look like
       something like this:
       ```
       sdx           x:xx   1  128M  0 disk
       └─sdx1        x:xx   1  128M  0 part
       ```

    2. Mount the device to your filesystem:
       ```sh
       mkdir /mnt/RPI-RP2
       mount /dev/sdx1 /mnt/RPI-RP2
       ```

4. Flash the device:
   ```sh
   tinygo flash -target=pico-w
   ```

The device should now reboot and after some time the display should show the time.

---
<p align="center">Logo created with <a href=https://github.com/quasilyte/gopherkon)>gopherkon</a></p>
