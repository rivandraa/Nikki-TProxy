include $(TOPDIR)/rules.mk

PKG_NAME:=nikki
PKG_RELEASE:=1

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://github.com/MetaCubeX/mihomo.git
#PKG_SOURCE_DATE:=
PKG_SOURCE_VERSION:=fd959feff28581d212b2dc7588be986f50834e28
PKG_MIRROR_HASH:=6c3a38ff47c3983e83ab6592d7330e837c43142dac8227685a08f1269972fe08

PKG_LICENSE:=GPL3.0+
PKG_MAINTAINER:=***

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_BUILD_FLAGS:=no-mips16

PKG_BUILD_VERSION:=v1.19.9
PKG_BUILD_TIME:=$(shell date -u -Iseconds)
#PKG_VERSION:=$(PKG_SOURCE_DATE)~$(PKG_BUILD_VERSION)-Mod
PKG_VERSION:=$(PKG_BUILD_VERSION)-Mod

GO_PKG:=github.com/metacubex/mihomo
GO_PKG_LDFLAGS_X:= \
	$(GO_PKG)/constant.Version=$(PKG_BUILD_VERSION) \
	$(GO_PKG)/constant.BuildTime=$(PKG_BUILD_TIME)
GO_PKG_TAGS:=with_gvisor

include $(INCLUDE_DIR)/package.mk
include $(TOPDIR)/feeds/packages/lang/golang/golang-package.mk

define Package/nikki
  SECTION:=net
  CATEGORY:=Network
  TITLE:=A rule based proxy in Go.
  URL:=https://wiki.metacubex.one
  DEPENDS:=$(GO_ARCH_DEPENDS) +ca-bundle +curl +yq firewall4 +ip-full \
           +kmod-inet-diag +kmod-nft-socket +kmod-nft-tproxy +kmod-tun
  PROVIDES:=nikki mihomo
endef

define Package/nikki/description
  A rule based proxy in Go.
endef

define Package/nikki/conffiles
/etc/config/nikki
/etc/nikki/mixin.yaml
/etc/nikki/nftables/reserved_ip.nft
/etc/nikki/nftables/reserved_ip6.nft
endef

define Package/nikki/install
	$(call GoPackage/Package/Install/Bin,$(1))

	$(INSTALL_DIR) $(1)/etc/nikki \
	               $(1)/etc/nikki/ucode \
	               $(1)/etc/nikki/scripts \
	               $(1)/etc/nikki/nftables \
	               $(1)/etc/nikki/profiles \
	               $(1)/etc/nikki/subscriptions \
	               $(1)/etc/nikki/run \
	               $(1)/etc/nikki/run/providers \
	               $(1)/etc/nikki/run/providers/rule \
	               $(1)/etc/nikki/run/providers/proxy

	$(INSTALL_DATA) $(CURDIR)/files/mixin.yaml $(1)/etc/nikki/mixin.yaml

	$(INSTALL_BIN) $(CURDIR)/files/ucode/include.uc $(1)/etc/nikki/ucode/include.uc
	$(INSTALL_BIN) $(CURDIR)/files/ucode/mixin.uc $(1)/etc/nikki/ucode/mixin.uc
	$(INSTALL_BIN) $(CURDIR)/files/ucode/hijack.ut $(1)/etc/nikki/ucode/hijack.ut

	$(INSTALL_BIN) $(CURDIR)/files/scripts/include.sh $(1)/etc/nikki/scripts/include.sh
	$(INSTALL_BIN) $(CURDIR)/files/scripts/firewall_include.sh $(1)/etc/nikki/scripts/firewall_include.sh
	$(INSTALL_BIN) $(CURDIR)/files/scripts/debug.sh $(1)/etc/nikki/scripts/debug.sh

	$(INSTALL_BIN) $(CURDIR)/files/nftables/reserved_ip.nft $(1)/etc/nikki/nftables/reserved_ip.nft
	$(INSTALL_BIN) $(CURDIR)/files/nftables/reserved_ip6.nft $(1)/etc/nikki/nftables/reserved_ip6.nft

	$(INSTALL_DIR) $(1)/etc/config
	$(INSTALL_CONF) $(CURDIR)/files/nikki.conf $(1)/etc/config/nikki

	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) $(CURDIR)/files/nikki.init $(1)/etc/init.d/nikki

	$(INSTALL_DIR) $(1)/lib/upgrade/keep.d
	$(INSTALL_DATA) $(CURDIR)/files/nikki.upgrade $(1)/lib/upgrade/keep.d/nikki
endef

define Package/nikki/postrm
#!/bin/sh
# Hanya dijalankan jika bukan dalam proses image build (chroot)
if [ -z "$$IPKG_INSTROOT" ]; then
	uci -q delete firewall.nikki
	uci -q commit firewall
fi

rm -rf /etc/nikki
rm -rf /www/luci-static/resources/view/nikki
rm -rf /tmp/log/nikki
rm -f /usr/share/luci/menu.d/luci-app-nikki.json
rm -f /www/tinyfm/nikkifm.php
rm -f /etc/config/nikki
rm -f /etc/init.d/nikki
rm -f /tmp/lock/procd_nikki.lock
rm -f /tmp/run/nikki/profile.json
exit 0
endef

define Build/Prepare
	$(Build/Prepare/Default)
	$(RM) -r $(PKG_BUILD_DIR)/rules/logic_test
endef

$(eval $(call GoBinPackage,nikki))
$(eval $(call BuildPackage,nikki))
