# `sandbox/api/User.go`

| Constant | Value | Description |
| --- | --- | --- |
| `UserRoleAdmin` | `iota` | UserRoleAdmin is the admin role |
| `UserRoleViewer` |  | UserRoleViewer is the viewer role |

## `UserRole`

UserRole is the role of the user

`type UserRole int`

## `User`

User is the user that is logged in

| Field | Type |
| --- | --- |
| `Id` | `string` |
| `Email` | `string` |
| `Username` | `string` |
| `Role` | `UserRole` |

[every contract](doc.md)
