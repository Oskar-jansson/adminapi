package models

type UserGroup struct {
	Id        *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Groupname *string `json:"groupname" validate:"required"`
	Rastamp   *string `json:"rastamp" validate:"required" cmp:"skip"`
	Groupdata *string `json:"groupdata"`
	Users     []User  `json:"users"`
}

type UserGroupInput struct {
	Groupname *string `json:"groupname,omitempty"`
	Rastamp   *string `json:"rastamp,omitempty"`
	Groupdata *string `json:"groupdata,omitempty"`
}

type UserGroupList struct {
	Usergroups []UserGroup `json:"usergroups"`
	Count      *int        `json:"count,omitempty"` // only included with ?count='true'
}

func (u *UserGroup) ConvToUsergroupInput() *UserGroupInput {
	return &UserGroupInput{
		Groupname: u.Groupname,
		Rastamp:   u.Rastamp,
		Groupdata: u.Groupdata,
	}
}

func (u *UserGroupInput) ConvToUsergroup() *UserGroup {
	return &UserGroup{
		Groupname: u.Groupname,
		Rastamp:   u.Rastamp,
		Groupdata: u.Groupdata,
	}
}

// u2 fields take precedence over u1 fields (if u2 field is non-nil, use it).
func (u1 *UserGroup) MergeUsergroups(u2 *UserGroup) *UserGroup {
	return &UserGroup{
		Id:        coalesceUint32(u2.Id, u1.Id),
		Groupname: coalesceString(u2.Groupname, u1.Groupname),
		Rastamp:   coalesce(u2.Rastamp, u1.Rastamp), //nolint:all
		Groupdata: coalesceString(u2.Groupdata, u1.Groupdata),
		Users:     coalesceSlice(u2.Users, u1.Users),
	}
}
