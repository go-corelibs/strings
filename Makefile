#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += STRINGS
STRINGS_MK_SUMMARY := go-corelibs/strings
STRINGS_MK_VERSION := v1.0.0

include CoreLibs.mk
