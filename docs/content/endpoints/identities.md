---
title: Identities
---

## Introduction

### Concept

**Identity** is the core entity used for authorization and authentication at Misakey.

They are auto-generated: the system consider a new wild identity appears when an
unknown identifier is entered during a login flow.

They refer then to an identifier (email, phone number...) and they are optionally attached to [an account](../accounts/).

End-users can claim an identity by creating an account on it or link an existing account.

They can be considered both as profiles and concrete identities. Users have many identities:
- as citizens: having both Korean and Russian nationalities...
- as internet fellows: having a Travian account and a Mastodon account...

### Resource Access

Resources access is based on the identity the end-user is logged in.

The end-user can select the identity they want to be logged in during the auth flow,
or in the interface at any moment.

There is no need to re-authenticate switching identities unless the security level
of an identity is higher than the current one the end-user is logged in.

## Create an account on an identity

Described [in accounts section](../accounts/#create-an-account-on-an-identity)

## Get an identity

This route allows the retrieval of the information related to an identity.

### Request

```bash
  GET https://api.misakey.com.local/identities/:id
```
_Headers:_
- `Authorization` (opaque token) (ACR >= 1): `subject` claim as the identity id.

_Path Parameters:_
- `id` (uuid string): the identity unique id.

### Success Response

_Code:_
```bash
  HTTP 200 CREATED
```

_JSON Body:_
```json
  {
    "id": "89a27dec-c0bb-40ed-bfc8-dc74a1b99dc9",
    "account_id": null,
    "identifier_id": "abc9dd47-0c4e-4ef3-ae27-0c21ea6d7450",
    "is_authable": true,
    "display_name": "non@dpo.com",
    "notifications": "minimal",
    "avatar_url": null,
    "confirmed": true
  }
```

- `id` (uuid string): the unique identity id.
- `account_id` (uuid string) (nullable): the linked account unique id.
- `identifier_id` (uuid string): the linked identifier unique id.
- `is_authable` (uuid string): either the identity can be used in a login flow.
- `display_name` (uuid string): the name to display to represent the identity.
- `notifications` (uuid string): the frequency of notifications for this identity.
- `avatar_url` (uuid string) (nullable): the web-address of the avatar's file content.
- `confirmed` (uuid string): either the identity has been guaranted to be owned by the end-user.