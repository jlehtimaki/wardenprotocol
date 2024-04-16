<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry is required to include a tag and
the GitHub issue reference in the following format:

* (<tag>) \#<issue-number> message

The tag should consist of where the change is being made ex. (x/staking), (store)
The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Bug Fixes" for any bug fixes.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"Consensus Breaking CHANGES" for any changes that result in a different AppState given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.1.0/
-->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Consensus Breaking Changes

* (x/intent) [#139](https://github.com/warden-protocol/wardenprotocol/pull/139) Ability for modules to dynamically resolve variables on Action creation
    * x/warden can now resolve `warden.space.owners` in Intent definitions into the list of owners of the space
* (x/intent) [#151](https://github.com/warden-protocol/wardenprotocol/pull/151) Store Intents' AST, instead of the raw string
* (x/warden) [#152](https://github.com/warden-protocol/wardenprotocol/pull/152) Prevent adding invalid addresses as Keychain parties
* (cosmwasm) [#156](https://github.com/warden-protocol/wardenprotocol/pull/156) Add support for smart contracts by integrating CosmWasm
* (x/warden) [#173](https://github.com/warden-protocol/wardenprotocol/pull/173) Remove `SignTransactionRequests` in favor of just `SignatureRequests`

### Features

* (shield) [#148](https://github.com/warden-protocol/wardenprotocol/pull/148) Use protobufs for defining the AST
* (x/intent) [#155](https://github.com/warden-protocol/wardenprotocol/pull/155) Add MsgUpdateIntent, creators of an Intent can use it to change name and definition of their Intents.
* (x/warden) [#159](https://github.com/warden-protocol/wardenprotocol/pull/159) Resolve `warden.space.owners` in Intent definitions for MsgNewSignatureRequest and MsgNewSignTransactionRequest
* (x/warden) [#160](https://github.com/warden-protocol/wardenprotocol/pull/160) Add Osmosis support
    * Derive Osmosis addresses from ECDSA_SECP256K1 keys
    * Extract DataForSigning for Osmosis Amino JSON transactions
* (cosmwasm) [#171](https://github.com/warden-protocol/wardenprotocol/pull/171) Add support for calling warden module from contracts

### Bug Fixes

### Misc

* (docs) [#127](https://github.com/warden-protocol/wardenprotocol/pull/127) Add CHANGELOG.md
* (ci) [#137](https://github.com/warden-protocol/wardenprotocol/pull/137) Add CodeRabbit configuration file, copied from Cosmos SDK's repo
* (perf) [#138](https://github.com/warden-protocol/wardenprotocol/pull/138) Add benchmarks for most hit queries in `x/warden` and `x/intent` (ActionsByAddress, AllKeys, KeysBySpaceId)
* (chore) [#180](https://github.com/warden-protocol/wardenprotocol/pull/180) Update to use pnpm v9.0.0

## [v0.2.0](https://github.com/warden-protocol/wardenprotocol/releases/tag/v0.2.0) - 2024-03-26

### Consensus Breaking Changes

* (x/warden) [#77](https://github.com/warden-protocol/wardenprotocol/pull/77) Use uint64 IDs for Space and Keychain, instead of string bech32 addresses
* (x/intent) [#55](https://github.com/warden-protocol/wardenprotocol/pull/55) Initial release of the `shield` intent engine, replacing the old `boolparser` and cleaning up the API
* (x/intent) [#112](https://github.com/warden-protocol/wardenprotocol/pull/112) Store a list of referenced addresses when a new intent is created. This is useful for querying all intents that reference a given address.
* Refactor of `x/intent`'s and `x/warden`'s store design to improve queries performance:
    * (x/intent) [#117](https://github.com/warden-protocol/wardenprotocol/pull/117) Pin the intent definition when an Action is created instead of just referencing it. This allows faster queries "by address" by storing the list of addresses directly inside the action, instead of having to re-evaluate the linked intent for each action.
    * (x/warden) [#121](https://github.com/warden-protocol/wardenprotocol/pull/121) Separate Keys queries into `AllKeys`, `KeysBySpaceId`, and `KeyById`. This allowed some improvements, especially in the `KeysBySpaceId` that can benefit from a new "space id -> key id" index.
    * (x/warden) [#122](https://github.com/warden-protocol/wardenprotocol/pull/122) Improve `SpacesByOwner` query by adding a new "owner -> space id" index.

### Features

* (client) [#116](https://github.com/warden-protocol/wardenprotocol/pull/116) Add `warden q wait-tx` command to wait for a transaction to be included in a block
    * The same command was submitted to Cosmos SDK: https://github.com/cosmos/cosmos-sdk/pull/19870
* (clichain) [#113](https://github.com/warden-protocol/wardenprotocol/pull/113) Initial release of `clichain`, a CLI for acting as a Keychain from the command line
* (x/warden) [#109](https://github.com/warden-protocol/wardenprotocol/pull/109) Add ability to assign an intent to a key, overriding the SignIntent of the Space
    * [#110](https://github.com/warden-protocol/wardenprotocol/pull/110) Add MsgUpdateKey for updating a key's intent
* (keychain-sdk) [#92](https://github.com/warden-protocol/wardenprotocol/pull/92) Increase Keychain SDK's throughput by batching requests into a single transaction
* (snap) [#83](https://github.com/warden-protocol/wardenprotocol/pull/83) Initial release of Warden Protocol's MetaMask snap
* (build) [#69](https://github.com/warden-protocol/wardenprotocol/pull/69) Add Makefile for building `wardend`
* (faucet) [#99](https://github.com/warden-protocol/wardenprotocol/pull/99) Rate limit faucet requests by IP address, not only by account
* (faucet) [#66](https://github.com/warden-protocol/wardenprotocol/pull/66) Batch multiple faucet request in a single transaction

### Bug Fixes

* (x/warden) [#120](https://github.com/warden-protocol/wardenprotocol/pull/120) Don't stop Keys query if one key is invalid (i.e. fails to derive a layer 1 address)


## [v0.1.0](https://github.com/warden-protocol/wardenprotocol/releases/tag/v0.1.0) - 2024-03-06

Initial release of `wardend` and launch of the Alfama testnet.


