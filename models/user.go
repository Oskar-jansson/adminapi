package models

type User struct {
	Id                    *uint32       `json:"id" validate:"gt=0" cmp:"skip"`
	Firstname             *string       `json:"firstname" validate:"required"`
	Lastname              *string       `json:"lastname" validate:"required"`
	Address               *string       `json:"address"`
	City                  *string       `json:"city"`
	Postalcode            *string       `json:"postalcode"`
	Phone1                *string       `json:"phone1"`
	Phone2                *string       `json:"phone2"`
	Phone3                *string       `json:"phone3"`
	Phoneconnectionnumber *int          `json:"phoneconnectionnumber"`
	Phoneconnectiontext   *string       `json:"phoneconnectiontext"`
	Email                 *string       `json:"email"`
	Extra                 *string       `json:"extra"`
	Employmentnumber      *uint32       `json:"employmentnumber"`
	Employmenttext        *string       `json:"employmenttext"`
	Customfield1          *string       `json:"customfield1"`
	Customfield2          *string       `json:"customfield2"`
	Customfield3          *string       `json:"customfield3"`
	Customfield4          *string       `json:"customfield4"`
	Customfield5          *string       `json:"customfield5"`
	Webpassword           *string       `json:"webpassword" cmp:"skip"` // non-readable
	Type                  *uint32       `json:"type"`
	Fkdepartment          *uint32       `json:"fkdepartment"`
	Fkusergroup           *uint32       `json:"fkusergroup"`
	Rastamp               *string       `json:"rastamp" validate:"required" cmp:"skip"`
	Fkfloor               *uint32       `json:"fkfloor"`
	Status                *uint32       `json:"status" validate:"required"`
	Startdate             *string       `json:"startdate" validate:"required"`
	Enddate               *string       `json:"enddate" validate:"required"`
	Refguid               *string       `json:"refguid" validate:"required"`
	Customfield6          *string       `json:"customfield6"`
	Customfield7          *string       `json:"customfield7"`
	Customfield8          *string       `json:"customfield8"`
	Customfield9          *string       `json:"customfield9"`
	Customfield10         *string       `json:"customfield10"`
	Apireference          *string       `json:"apireference"`
	Changedby             *string       `json:"changedby" validate:"required" cmp:"skip"`
	Changeddate           *string       `json:"changeddate" validate:"required" cmp:"skip"`
	Createdby             *string       `json:"createdby" validate:"required" cmp:"skip"`
	Createddate           *string       `json:"createddate" validate:"required" cmp:"skip"`
	Showregister          *bool         `json:"showregister" validate:"required"`
	Cardgroupname         *string       `json:"cardgroupname"`
	Message               *string       `json:"message"`
	Cardgroupstamp        *string       `json:"cardgroupstamp"`
	Cardtype              *int          `json:"cardtype"`
	Readinfomsg           *uint32       `json:"readinfomsg"`
	Remindpass            *int          `json:"remindpass"`
	Remindmachine         *int          `json:"remindmachine"`
	Balance               *float64      `json:"balance"`
	Lastcalc              *string       `json:"lastcalc"`
	Showmeasure           *int          `json:"showmeasure"`
	Showbooked            *int          `json:"showbooked"`
	Language              *int          `json:"language"`
	Cards                 []Card        `json:"cards"`
	Department            *Department   `json:"department"`
	Usergroup             *UserGroup    `json:"usergroup"`
	Floor                 *Floor        `json:"floor"`
	Accessgroups          []AccessGroup `json:"accessgroups"`
}

// Writeable fields in user
// Uses omitempty to allow incomplete structs to exclude said fields.
type UserInput struct {
	Firstname             *string `json:"firstname,omitempty"`
	Lastname              *string `json:"lastname,omitempty"`
	Address               *string `json:"address,omitempty"`
	City                  *string `json:"city,omitempty"`
	Postalcode            *string `json:"postalcode,omitempty"`
	Phone1                *string `json:"phone1,omitempty"`
	Phone2                *string `json:"phone2,omitempty"`
	Phone3                *string `json:"phone3,omitempty"`
	Phoneconnectionnumber *int    `json:"phoneconnectionnumber,omitempty"`
	Phoneconnectiontext   *string `json:"phoneconnectiontext,omitempty"`
	Email                 *string `json:"email,omitempty"`
	Extra                 *string `json:"extra,omitempty"`
	Employmentnumber      *uint32 `json:"employmentnumber,omitempty"`
	Employmenttext        *string `json:"employmenttext,omitempty"`
	Customfield1          *string `json:"customfield1,omitempty"`
	Customfield2          *string `json:"customfield2,omitempty"`
	Customfield3          *string `json:"customfield3,omitempty"`
	Customfield4          *string `json:"customfield4,omitempty"`
	Customfield5          *string `json:"customfield5,omitempty"`
	Webpassword           *string `json:"webpassword,omitempty"`

	// non writeable. But needed for valid PATCH requests.
	// Should always be current Rastamp of said user.
	Rastamp *string `json:"rastamp,omitempty"`

	Type          *uint32 `json:"type,omitempty"`
	Fkdepartment  *uint32 `json:"fkdepartment,omitempty"`
	Fkusergroup   *uint32 `json:"fkusergroup,omitempty"`
	Fkfloor       *uint32 `json:"fkfloor,omitempty"`
	Status        *uint32 `json:"status,omitempty"`
	Startdate     *string `json:"startdate,omitempty"`
	Enddate       *string `json:"enddate,omitempty"`
	Customfield6  *string `json:"customfield6,omitempty"`
	Customfield7  *string `json:"customfield7,omitempty"`
	Customfield8  *string `json:"customfield8,omitempty"`
	Customfield9  *string `json:"customfield9,omitempty"`
	Customfield10 *string `json:"customfield10,omitempty"`
	Apireference  *string `json:"apireference,omitempty"`
	Showregister  *bool   `json:"showregister,omitempty"`
	Cardgroupname *string `json:"cardgroupname,omitempty"`
	Message       *string `json:"message,omitempty"`
	Cardtype      *int    `json:"cardtype,omitempty"`

	//Cards        *[]Card
	//Department   *Department
	//Usergroup    *Usergroup
	//Floor        *Floor
	//Accessgroups *[]Accessgroup
}

type UserList struct {
	Users []User `json:"users"`
	Count *int   `json:"count,omitempty"` // only included with ?count='true'
}

// Converts UserInput into User struct
func (u *User) ConvToUserInput() *UserInput {
	return &UserInput{
		Firstname:             u.Firstname,
		Lastname:              u.Lastname,
		Address:               u.Address,
		City:                  u.City,
		Postalcode:            u.Postalcode,
		Phone1:                u.Phone1,
		Phone2:                u.Phone2,
		Phone3:                u.Phone3,
		Phoneconnectionnumber: u.Phoneconnectionnumber,
		Phoneconnectiontext:   u.Phoneconnectiontext,
		Email:                 u.Email,
		Extra:                 u.Extra,
		Employmentnumber:      u.Employmentnumber,
		Employmenttext:        u.Employmenttext,
		Customfield1:          u.Customfield1,
		Customfield2:          u.Customfield2,
		Customfield3:          u.Customfield3,
		Customfield4:          u.Customfield4,
		Customfield5:          u.Customfield5,
		Webpassword:           u.Webpassword,
		Rastamp:               u.Rastamp,
		Type:                  u.Type,
		Fkdepartment:          u.Fkdepartment,
		Fkusergroup:           u.Fkusergroup,
		Fkfloor:               u.Fkfloor,
		Status:                u.Status,
		Startdate:             u.Startdate,
		Enddate:               u.Enddate,
		Customfield6:          u.Customfield6,
		Customfield7:          u.Customfield7,
		Customfield8:          u.Customfield8,
		Customfield9:          u.Customfield9,
		Customfield10:         u.Customfield10,
		Apireference:          u.Apireference,
		Showregister:          u.Showregister,
		Cardgroupname:         u.Cardgroupname,
		Message:               u.Message,
		Cardtype:              u.Cardtype,
	}
}

// Converts UserInput into User struct
func (u *UserInput) ConvToUser() *User {
	return &User{
		Firstname:             u.Firstname,
		Lastname:              u.Lastname,
		Address:               u.Address,
		City:                  u.City,
		Postalcode:            u.Postalcode,
		Phone1:                u.Phone1,
		Phone2:                u.Phone2,
		Phone3:                u.Phone3,
		Phoneconnectionnumber: u.Phoneconnectionnumber,
		Phoneconnectiontext:   u.Phoneconnectiontext,
		Email:                 u.Email,
		Extra:                 u.Extra,
		Employmentnumber:      u.Employmentnumber,
		Employmenttext:        u.Employmenttext,
		Customfield1:          u.Customfield1,
		Customfield2:          u.Customfield2,
		Customfield3:          u.Customfield3,
		Customfield4:          u.Customfield4,
		Customfield5:          u.Customfield5,
		Webpassword:           u.Webpassword,
		Rastamp:               u.Rastamp,
		Type:                  u.Type,
		Fkdepartment:          u.Fkdepartment,
		Fkusergroup:           u.Fkusergroup,
		Fkfloor:               u.Fkfloor,
		Status:                u.Status,
		Startdate:             u.Startdate,
		Enddate:               u.Enddate,
		Customfield6:          u.Customfield6,
		Customfield7:          u.Customfield7,
		Customfield8:          u.Customfield8,
		Customfield9:          u.Customfield9,
		Customfield10:         u.Customfield10,
		Apireference:          u.Apireference,
		Showregister:          u.Showregister,
		Cardgroupname:         u.Cardgroupname,
		Message:               u.Message,
		Cardtype:              u.Cardtype,
	}
}

// u2 fields take precedence over u1 fields (if u2 field is non-nil, use it).
func (u1 *User) MergeUsers(u2 *User) *User {
	return &User{
		Id:                    coalesceUint32(u2.Id, u1.Id),
		Firstname:             coalesceString(u2.Firstname, u1.Firstname),
		Lastname:              coalesceString(u2.Lastname, u1.Lastname),
		Address:               coalesceString(u2.Address, u1.Address),
		City:                  coalesceString(u2.City, u1.City),
		Postalcode:            coalesceString(u2.Postalcode, u1.Postalcode),
		Phone1:                coalesceString(u2.Phone1, u1.Phone1),
		Phone2:                coalesceString(u2.Phone2, u1.Phone2),
		Phone3:                coalesceString(u2.Phone3, u1.Phone3),
		Phoneconnectionnumber: coalesceInt(u2.Phoneconnectionnumber, u1.Phoneconnectionnumber),
		Phoneconnectiontext:   coalesceString(u2.Phoneconnectiontext, u1.Phoneconnectiontext),
		Email:                 coalesceString(u2.Email, u1.Email),
		Extra:                 coalesceString(u2.Extra, u1.Extra),
		Employmentnumber:      coalesceUint32(u2.Employmentnumber, u1.Employmentnumber),
		Employmenttext:        coalesceString(u2.Employmenttext, u1.Employmenttext),
		Customfield1:          coalesceString(u2.Customfield1, u1.Customfield1),
		Customfield2:          coalesceString(u2.Customfield2, u1.Customfield2),
		Customfield3:          coalesceString(u2.Customfield3, u1.Customfield3),
		Customfield4:          coalesceString(u2.Customfield4, u1.Customfield4),
		Customfield5:          coalesceString(u2.Customfield5, u1.Customfield5),
		Webpassword:           coalesceString(u2.Webpassword, u1.Webpassword),
		Type:                  coalesceUint32(u2.Type, u1.Type),
		Fkdepartment:          coalesceUint32(u2.Fkdepartment, u1.Fkdepartment),
		Fkusergroup:           coalesceUint32(u2.Fkusergroup, u1.Fkusergroup),
		Rastamp:               coalesce(u2.Rastamp, u1.Rastamp), //nolint:all
		Fkfloor:               coalesceUint32(u2.Fkfloor, u1.Fkfloor),
		Status:                coalesceUint32(u2.Status, u1.Status),
		Startdate:             coalesceString(u2.Startdate, u1.Startdate),
		Enddate:               coalesceString(u2.Enddate, u1.Enddate),
		Refguid:               coalesceString(u2.Refguid, u1.Refguid),
		Customfield6:          coalesceString(u2.Customfield6, u1.Customfield6),
		Customfield7:          coalesceString(u2.Customfield7, u1.Customfield7),
		Customfield8:          coalesceString(u2.Customfield8, u1.Customfield8),
		Customfield9:          coalesceString(u2.Customfield9, u1.Customfield9),
		Customfield10:         coalesceString(u2.Customfield10, u1.Customfield10),
		Apireference:          coalesceString(u2.Apireference, u1.Apireference),
		Changedby:             coalesce(u2.Changedby, u1.Changedby),     //nolint:all
		Changeddate:           coalesce(u2.Changeddate, u1.Changeddate), //nolint:all
		Createdby:             coalesce(u2.Createdby, u1.Createdby),     //nolint:all
		Createddate:           coalesce(u2.Createddate, u1.Createddate), //nolint:all
		Showregister:          coalesceBool(u2.Showregister, u1.Showregister),
		Cardgroupname:         coalesceString(u2.Cardgroupname, u1.Cardgroupname),
		Message:               coalesceString(u2.Message, u1.Message),
		Cardgroupstamp:        coalesceString(u2.Cardgroupstamp, u1.Cardgroupstamp),
		Cardtype:              coalesceInt(u2.Cardtype, u1.Cardtype),
		Readinfomsg:           coalesceUint32(u2.Readinfomsg, u1.Readinfomsg),
		Remindpass:            coalesceInt(u2.Remindpass, u1.Remindpass),
		Remindmachine:         coalesceInt(u2.Remindmachine, u1.Remindmachine),
		Balance:               coalesceFloat64(u2.Balance, u1.Balance),
		Lastcalc:              coalesceString(u2.Lastcalc, u1.Lastcalc),
		Showmeasure:           coalesceInt(u2.Showmeasure, u1.Showmeasure),
		Showbooked:            coalesceInt(u2.Showbooked, u1.Showbooked),
		Language:              coalesceInt(u2.Language, u1.Language),

		//nolint:all
		Cards: coalesceSlice(u2.Cards, u1.Cards),

		Department:   coalesce(u2.Department, u1.Department),
		Usergroup:    coalesce(u2.Usergroup, u1.Usergroup),
		Floor:        coalesce(u2.Floor, u1.Floor),
		Accessgroups: coalesceSlice(u2.Accessgroups, u1.Accessgroups),
	}
}
