![GitHub License](https://img.shields.io/github/license/rivandraa/Nikki-TProxy?style=for-the-badge&logo=github) ![GitHub Tag](https://img.shields.io/github/v/release/rivandraa/Nikki-TProxy?style=for-the-badge&logo=github) ![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/rivandraa/Nikki-TProxy/total?style=for-the-badge&logo=github) ![GitHub Repo stars](https://img.shields.io/github/stars/rivandraa/Nikki-TProxy?style=for-the-badge&logo=github)

# Nikki TProxy

Transparent Proxy with Mihomo on OpenWrt.

## Prerequisites

- OpenWrt >= 23.05
- Linux Kernel >= 5.13
- firewall4

## Feature

- Transparent Proxy (Redirect/TPROXY/TUN, IPv4 and/or IPv6)
- Access Control
- Profile Mixin
- Profile Editor
- Scheduled Restart

The package files will be found under `bin/packages/your_architecture/nikki`.

## Dependencies

- ca-bundle
- curl
- yq
- firewall4
- ip-full
- kmod-inet-diag
- kmod-nft-socket
- kmod-nft-tproxy
- kmod-tun
