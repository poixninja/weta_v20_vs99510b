#!/bin/bash

IMG1=~/kernel/weta_v20_vs99510b/build/arch/arm64/boot/Image.lz4-dtb
IMG2=~/kernel/superr/super_vs995/boot.img-zImage

make mrproper
./build.sh vs995

sudo mv $IMG1 $IMG2

