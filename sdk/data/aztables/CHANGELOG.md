# Release History

## 0.5.1 (Unreleased)

### Features Added

### Breaking Changes

### Bugs Fixed

### Other Changes

## 0.5.0 (2022-01-12)

### Other Changes
* Updates `azcore` dependency from `v0.20.0` to `v0.21.0`

## 0.4.0 (2021-11-09)

### Features Added
* Added `NextPagePartitionKey` and `NextPageRowKey` to `ListEntitiesPager` for retrieving continuation tokens.
* Added `PartitionKey` and `RowKey` to `ListEntitiesOptions` for using exposed continuation tokens.

### Bug Fixes
* Fixed a bug on transactional batches where `InsertMerge` and `InsertReplace` failed if the entity did not exist.

## 0.3.0 (2021-11-02)

### Features Added
* Added `NewClientWithNoCredential` and `NewServiceClientWithNoCredential` for authenticating the `Client` and `ServiceClient` with SAS URLs
* Added `NewClientWithSharedKey` and `NewServiceClientWithSharedKey` for authenticating the `Client` and `ServiceClient` with Shared Keys

### Breaking Changes
* `NewClient` and `NewServiceClient` is now used for authenticating the `Client` and `ServiceClient` with credentials from `azidentity` only.
* `ClientOptions` embeds `azcore.ClientOptions` and removes all named fields.

## 0.2.0 (2021-10-05)

### Features Added
* Added the `ClientOptions.PerTryPolicies` for policies that execute once per retry of an operation.

### Breaking Changes
* Changed the `ClientOptions.PerCallOptions` field name to `ClientOptions.PerCallPolicies`
* Changed the `ClientOptions.Transporter` field name to `ClientOptions.Transport`

## 0.1.0 (09-07-2021)
* This is the initial release of the `aztables` library
