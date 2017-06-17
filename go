#!/bin/bash
############################################
### mentalmuso starter and output script ###
############################################

VERS=1.23
FIR=vs99510b

BUILD=~/kernel/weta_v20_vs99510b/build/arch/arm64/boot/Image.lz4-dtb
IMG=~/kernel/weta_v20_vs99510b/imgout/boot
IMGOUT=~/kernel/weta_v20_vs99510b/imgout
ZIP=~/kernel/weta_v20_vs99510b/imgout/zip
ZIPOUT=~/kernel/weta_v20_vs99510b/zipout
WETAIMG=~/kernel/weta_v20_vs99510b/imgout/zip/kernel/weta_boot.img

export VER=$VERS
export FIRM=$FIR

############################################
### ---> clean up and build

make mrproper
./build.sh vs995

############################################
### ---> make flashable zip 

mv $BUILD $IMG/boot.img-kernel
cd $IMG
./mkbootimg --kernel boot.img-kernel --ramdisk boot.img-ramdisk.cpio.gz -o $WETAIMG
cd $ZIP
rm -f $ZIPOUT/*
zip -r $ZIPOUT/WETA_Kernel_$FIR_$VERS.zip *




