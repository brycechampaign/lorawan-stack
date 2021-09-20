// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	time "time"
)

func (dst *User) SetFields(src *User, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &UserIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "deleted_at":
			if len(subs) > 0 {
				return fmt.Errorf("'deleted_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DeletedAt = src.DeletedAt
			} else {
				dst.DeletedAt = nil
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "attributes":
			if len(subs) > 0 {
				return fmt.Errorf("'attributes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Attributes = src.Attributes
			} else {
				dst.Attributes = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "primary_email_address":
			if len(subs) > 0 {
				return fmt.Errorf("'primary_email_address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PrimaryEmailAddress = src.PrimaryEmailAddress
			} else {
				var zero string
				dst.PrimaryEmailAddress = zero
			}
		case "primary_email_address_validated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'primary_email_address_validated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PrimaryEmailAddressValidatedAt = src.PrimaryEmailAddressValidatedAt
			} else {
				dst.PrimaryEmailAddressValidatedAt = nil
			}
		case "password":
			if len(subs) > 0 {
				return fmt.Errorf("'password' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Password = src.Password
			} else {
				var zero string
				dst.Password = zero
			}
		case "password_updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'password_updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PasswordUpdatedAt = src.PasswordUpdatedAt
			} else {
				dst.PasswordUpdatedAt = nil
			}
		case "require_password_update":
			if len(subs) > 0 {
				return fmt.Errorf("'require_password_update' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RequirePasswordUpdate = src.RequirePasswordUpdate
			} else {
				var zero bool
				dst.RequirePasswordUpdate = zero
			}
		case "state":
			if len(subs) > 0 {
				return fmt.Errorf("'state' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.State = src.State
			} else {
				var zero State
				dst.State = zero
			}
		case "state_description":
			if len(subs) > 0 {
				return fmt.Errorf("'state_description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.StateDescription = src.StateDescription
			} else {
				var zero string
				dst.StateDescription = zero
			}
		case "admin":
			if len(subs) > 0 {
				return fmt.Errorf("'admin' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Admin = src.Admin
			} else {
				var zero bool
				dst.Admin = zero
			}
		case "temporary_password":
			if len(subs) > 0 {
				return fmt.Errorf("'temporary_password' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemporaryPassword = src.TemporaryPassword
			} else {
				var zero string
				dst.TemporaryPassword = zero
			}
		case "temporary_password_created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'temporary_password_created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemporaryPasswordCreatedAt = src.TemporaryPasswordCreatedAt
			} else {
				dst.TemporaryPasswordCreatedAt = nil
			}
		case "temporary_password_expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'temporary_password_expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemporaryPasswordExpiresAt = src.TemporaryPasswordExpiresAt
			} else {
				dst.TemporaryPasswordExpiresAt = nil
			}
		case "profile_picture":
			if len(subs) > 0 {
				var newDst, newSrc *Picture
				if (src == nil || src.ProfilePicture == nil) && dst.ProfilePicture == nil {
					continue
				}
				if src != nil {
					newSrc = src.ProfilePicture
				}
				if dst.ProfilePicture != nil {
					newDst = dst.ProfilePicture
				} else {
					newDst = &Picture{}
					dst.ProfilePicture = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ProfilePicture = src.ProfilePicture
				} else {
					dst.ProfilePicture = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Users) SetFields(src *Users, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "users":
			if len(subs) > 0 {
				return fmt.Errorf("'users' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Users = src.Users
			} else {
				dst.Users = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetUserRequest) SetFields(src *GetUserRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListUsersRequest) SetFields(src *ListUsersRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}
		case "deleted":
			if len(subs) > 0 {
				return fmt.Errorf("'deleted' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Deleted = src.Deleted
			} else {
				var zero bool
				dst.Deleted = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateUserRequest) SetFields(src *CreateUserRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user":
			if len(subs) > 0 {
				var newDst, newSrc *User
				if src != nil {
					newSrc = &src.User
				}
				newDst = &dst.User
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.User = src.User
				} else {
					var zero User
					dst.User = zero
				}
			}
		case "invitation_token":
			if len(subs) > 0 {
				return fmt.Errorf("'invitation_token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.InvitationToken = src.InvitationToken
			} else {
				var zero string
				dst.InvitationToken = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateUserRequest) SetFields(src *UpdateUserRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user":
			if len(subs) > 0 {
				var newDst, newSrc *User
				if src != nil {
					newSrc = &src.User
				}
				newDst = &dst.User
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.User = src.User
				} else {
					var zero User
					dst.User = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateTemporaryPasswordRequest) SetFields(src *CreateTemporaryPasswordRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateUserPasswordRequest) SetFields(src *UpdateUserPasswordRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "new":
			if len(subs) > 0 {
				return fmt.Errorf("'new' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.New = src.New
			} else {
				var zero string
				dst.New = zero
			}
		case "old":
			if len(subs) > 0 {
				return fmt.Errorf("'old' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Old = src.Old
			} else {
				var zero string
				dst.Old = zero
			}
		case "revoke_all_access":
			if len(subs) > 0 {
				return fmt.Errorf("'revoke_all_access' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RevokeAllAccess = src.RevokeAllAccess
			} else {
				var zero bool
				dst.RevokeAllAccess = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListUserAPIKeysRequest) SetFields(src *ListUserAPIKeysRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetUserAPIKeyRequest) SetFields(src *GetUserAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "key_id":
			if len(subs) > 0 {
				return fmt.Errorf("'key_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.KeyId = src.KeyId
			} else {
				var zero string
				dst.KeyId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateUserAPIKeyRequest) SetFields(src *CreateUserAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "rights":
			if len(subs) > 0 {
				return fmt.Errorf("'rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rights = src.Rights
			} else {
				dst.Rights = nil
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				dst.ExpiresAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateUserAPIKeyRequest) SetFields(src *UpdateUserAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "api_key":
			if len(subs) > 0 {
				var newDst, newSrc *APIKey
				if src != nil {
					newSrc = &src.APIKey
				}
				newDst = &dst.APIKey
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.APIKey = src.APIKey
				} else {
					var zero APIKey
					dst.APIKey = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Invitation) SetFields(src *Invitation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}
		case "token":
			if len(subs) > 0 {
				return fmt.Errorf("'token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Token = src.Token
			} else {
				var zero string
				dst.Token = zero
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				var zero time.Time
				dst.ExpiresAt = zero
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "accepted_at":
			if len(subs) > 0 {
				return fmt.Errorf("'accepted_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AcceptedAt = src.AcceptedAt
			} else {
				dst.AcceptedAt = nil
			}
		case "accepted_by":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.AcceptedBy == nil) && dst.AcceptedBy == nil {
					continue
				}
				if src != nil {
					newSrc = src.AcceptedBy
				}
				if dst.AcceptedBy != nil {
					newDst = dst.AcceptedBy
				} else {
					newDst = &UserIdentifiers{}
					dst.AcceptedBy = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AcceptedBy = src.AcceptedBy
				} else {
					dst.AcceptedBy = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListInvitationsRequest) SetFields(src *ListInvitationsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Invitations) SetFields(src *Invitations, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "invitations":
			if len(subs) > 0 {
				return fmt.Errorf("'invitations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Invitations = src.Invitations
			} else {
				dst.Invitations = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SendInvitationRequest) SetFields(src *SendInvitationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteInvitationRequest) SetFields(src *DeleteInvitationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserSessionIdentifiers) SetFields(src *UserSessionIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "session_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionId = src.SessionId
			} else {
				var zero string
				dst.SessionId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserSession) SetFields(src *UserSession, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "session_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionId = src.SessionId
			} else {
				var zero string
				dst.SessionId = zero
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				dst.ExpiresAt = nil
			}
		case "session_secret":
			if len(subs) > 0 {
				return fmt.Errorf("'session_secret' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionSecret = src.SessionSecret
			} else {
				var zero string
				dst.SessionSecret = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserSessions) SetFields(src *UserSessions, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "sessions":
			if len(subs) > 0 {
				return fmt.Errorf("'sessions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Sessions = src.Sessions
			} else {
				dst.Sessions = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListUserSessionsRequest) SetFields(src *ListUserSessionsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *LoginToken) SetFields(src *LoginToken, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if src != nil {
					newSrc = &src.UserIds
				}
				newDst = &dst.UserIds
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					var zero UserIdentifiers
					dst.UserIds = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				var zero time.Time
				dst.ExpiresAt = zero
			}
		case "token":
			if len(subs) > 0 {
				return fmt.Errorf("'token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Token = src.Token
			} else {
				var zero string
				dst.Token = zero
			}
		case "used":
			if len(subs) > 0 {
				return fmt.Errorf("'used' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Used = src.Used
			} else {
				var zero bool
				dst.Used = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateLoginTokenRequest) SetFields(src *CreateLoginTokenRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.UserIds == nil) && dst.UserIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserIds
				}
				if dst.UserIds != nil {
					newDst = dst.UserIds
				} else {
					newDst = &UserIdentifiers{}
					dst.UserIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserIds = src.UserIds
				} else {
					dst.UserIds = nil
				}
			}
		case "skip_email":
			if len(subs) > 0 {
				return fmt.Errorf("'skip_email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SkipEmail = src.SkipEmail
			} else {
				var zero bool
				dst.SkipEmail = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateLoginTokenResponse) SetFields(src *CreateLoginTokenResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "token":
			if len(subs) > 0 {
				return fmt.Errorf("'token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Token = src.Token
			} else {
				var zero string
				dst.Token = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
