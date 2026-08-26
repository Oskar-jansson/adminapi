package models

type Card struct {
	Id                *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Fkusernumber      *uint32 `json:"fkusernumber" validate:"required"`
	Name              *string `json:"name" validate:"required"`
	Type              *uint32 `json:"type" validate:"required"`
	Cardidentity      *string `json:"cardidentity"`
	Regnumber         *string `json:"regnumber"`
	Pincode           *string `json:"pincode"`
	Startdate         *string `json:"startdate"`
	Enddate           *string `json:"enddate"`
	Alarmoff          *string `json:"alarmoff"`
	Alarmon           *string `json:"alarmon"`
	Timecode          *uint32 `json:"timecode"`
	Timecodetype      *uint32 `json:"timecodetype"`
	Timebookings      *string `json:"timebookings"`
	Booktype          *string `json:"booktype"`
	Rastamp           *string `json:"rastamp" validate:"required" cmp:"skip"`
	Phoneshort        *string `json:"phoneshort"`
	Phonetele         *string `json:"phonetele"`
	Phoneteleall      *string `json:"phoneteleall" validate:"required"`
	Phoneuseexternal  *uint32 `json:"phoneuseexternal"`
	Disability        *uint32 `json:"disability"`
	Fieldex1          *string `json:"fieldex1"`
	Fieldex2          *string `json:"fieldex2"`
	Cardidentityraw   *string `json:"cardidentityraw" validate:"required"`
	Refguid           *string `json:"refguid" validate:"required"`
	Selectpindatetime *string `json:"selectpindatetime"`
	Apireference      *string `json:"apireference"`
	Changedby         *string `json:"changedby" validate:"required" cmp:"skip"`
	Changeddate       *string `json:"changeddate" validate:"required" cmp:"skip"`
	Createdby         *string `json:"createdby" validate:"required" cmp:"skip"`
	Createddate       *string `json:"createddate" validate:"required" cmp:"skip"`
	Asciicard         *bool   `json:"asciicard" validate:"required"`
	Showregister      *bool   `json:"showregister" validate:"required"`
	Pinblocked        *bool   `json:"pinblocked" validate:"required"`
	Isblocked         *bool   `json:"isblocked" validate:"required"`
	Inherituseraccess *bool   `json:"inherituseraccess" validate:"required"`
	Expired           *bool   `json:"expired" validate:"required"`
	User              *User   `json:"user"`
	Parentcard        *uint32 `json:"parentcard"` // undocumented
	Vacation          *uint32 `json:"vacation"`   // undocumented

	// exposed when using ?include=accessalarms
	Accessalarms *[]struct {
		Fkaccessgroup uint32 `json:"fkaccessgroup"`
		Alarmon       uint32 `json:"alarmon"`
		Alarmoff      uint32 `json:"alarmoff"`
	} `json:"accessalarms"`

	Accessgroups []AccessGroup  `json:"accessgroups"`
	Readeraccess []ReaderAccess `json:"readeraccess"`
}

type CardInput struct {
	Fkusernumber      *uint32 `json:"fkusernumber,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *uint32 `json:"type,omitempty"`
	Cardidentity      *string `json:"cardidentity,omitempty"`
	Regnumber         *string `json:"regnumber,omitempty"`
	Pincode           *string `json:"pincode,omitempty"`
	Startdate         *string `json:"startdate,omitempty"`
	Enddate           *string `json:"enddate,omitempty"`
	Alarmoff          *string `json:"alarmoff,omitempty"`
	Alarmon           *string `json:"alarmon,omitempty"`
	Timecode          *uint32 `json:"timecode,omitempty"`
	Timecodetype      *uint32 `json:"timecodetype,omitempty"`
	Timebookings      *string `json:"timebookings,omitempty"`
	Booktype          *string `json:"booktype,omitempty"`
	Rastamp           *string `json:"rastamp,omitempty"`
	Phoneshort        *string `json:"phoneshort,omitempty"`
	Phonetele         *string `json:"phonetele,omitempty"`
	Phoneteleall      *string `json:"phoneteleall,omitempty"`
	Phoneuseexternal  *uint32 `json:"phoneuseexternal,omitempty"`
	Disability        *uint32 `json:"disability,omitempty"`
	Fieldex1          *string `json:"fieldex1,omitempty"`
	Fieldex2          *string `json:"fieldex2,omitempty"`
	Cardidentityraw   *string `json:"cardidentityraw,omitempty"`
	Refguid           *string `json:"refguid,omitempty"`
	Selectpindatetime *string `json:"selectpindatetime,omitempty"`
	Apireference      *string `json:"apireference,omitempty"`
	Asciicard         *bool   `json:"asciicard,omitempty"`
	Showregister      *bool   `json:"showregister,omitempty"`
	Pinblocked        *bool   `json:"pinblocked,omitempty"`
	Isblocked         *bool   `json:"isblocked,omitempty"`
	Inherituseraccess *bool   `json:"inherituseraccess,omitempty"`

	//User              User    `json:",omitempty"`
	// Accessgroups      []Accessgroup
}

type CardList struct {
	Cards []Card `json:"cards"`
	Count *int   `json:"count,omitempty"` // only included with ?count='true'
}

func (c *Card) ConvToCardInput() *CardInput {
	return &CardInput{
		Fkusernumber:      c.Fkusernumber,
		Name:              c.Name,
		Type:              c.Type,
		Cardidentity:      c.Cardidentity,
		Regnumber:         c.Regnumber,
		Pincode:           c.Pincode,
		Startdate:         c.Startdate,
		Enddate:           c.Enddate,
		Alarmoff:          c.Alarmoff,
		Alarmon:           c.Alarmon,
		Timecode:          c.Timecode,
		Timecodetype:      c.Timecodetype,
		Timebookings:      c.Timebookings,
		Booktype:          c.Booktype,
		Rastamp:           c.Rastamp,
		Phoneshort:        c.Phoneshort,
		Phonetele:         c.Phonetele,
		Phoneteleall:      c.Phoneteleall,
		Phoneuseexternal:  c.Phoneuseexternal,
		Disability:        c.Disability,
		Fieldex1:          c.Fieldex1,
		Fieldex2:          c.Fieldex2,
		Cardidentityraw:   c.Cardidentityraw,
		Refguid:           c.Refguid,
		Selectpindatetime: c.Selectpindatetime,
		Apireference:      c.Apireference,
		Asciicard:         c.Asciicard,
		Showregister:      c.Showregister,
		Pinblocked:        c.Pinblocked,
		Isblocked:         c.Isblocked,
		Inherituseraccess: c.Inherituseraccess,
	}
}

func (c *CardInput) ConvToCard() *Card {
	return &Card{
		Fkusernumber:      c.Fkusernumber,
		Name:              c.Name,
		Type:              c.Type,
		Cardidentity:      c.Cardidentity,
		Regnumber:         c.Regnumber,
		Pincode:           c.Pincode,
		Startdate:         c.Startdate,
		Enddate:           c.Enddate,
		Alarmoff:          c.Alarmoff,
		Alarmon:           c.Alarmon,
		Timecode:          c.Timecode,
		Timecodetype:      c.Timecodetype,
		Timebookings:      c.Timebookings,
		Booktype:          c.Booktype,
		Rastamp:           c.Rastamp,
		Phoneshort:        c.Phoneshort,
		Phonetele:         c.Phonetele,
		Phoneteleall:      c.Phoneteleall,
		Phoneuseexternal:  c.Phoneuseexternal,
		Disability:        c.Disability,
		Fieldex1:          c.Fieldex1,
		Fieldex2:          c.Fieldex2,
		Cardidentityraw:   c.Cardidentityraw,
		Refguid:           c.Refguid,
		Selectpindatetime: c.Selectpindatetime,
		Apireference:      c.Apireference,
		Asciicard:         c.Asciicard,
		Showregister:      c.Showregister,
		Pinblocked:        c.Pinblocked,
		Isblocked:         c.Isblocked,
		Inherituseraccess: c.Inherituseraccess,
	}
}

// c2 fields take precedence over c1 fields (if c2 field is non-nil, use it).
func (c1 *Card) MergeCards(c2 *Card) *Card {
	return &Card{
		Id:                coalesceUint32(c2.Id, c1.Id),
		Fkusernumber:      coalesceUint32(c2.Fkusernumber, c1.Fkusernumber),
		Name:              coalesceString(c2.Name, c1.Name),
		Type:              coalesceUint32(c2.Type, c1.Type),
		Cardidentity:      coalesceString(c2.Cardidentity, c1.Cardidentity),
		Regnumber:         coalesceString(c2.Regnumber, c1.Regnumber),
		Pincode:           coalesceString(c2.Pincode, c1.Pincode),
		Startdate:         coalesceString(c2.Startdate, c1.Startdate),
		Enddate:           coalesceString(c2.Enddate, c1.Enddate),
		Alarmoff:          coalesceString(c2.Alarmoff, c1.Alarmoff),
		Alarmon:           coalesceString(c2.Alarmon, c1.Alarmon),
		Timecode:          coalesceUint32(c2.Timecode, c1.Timecode),
		Timecodetype:      coalesceUint32(c2.Timecodetype, c1.Timecodetype),
		Timebookings:      coalesceString(c2.Timebookings, c1.Timebookings),
		Booktype:          coalesceString(c2.Booktype, c1.Booktype),
		Rastamp:           coalesce(c2.Rastamp, c1.Rastamp), //nolint:all
		Phoneshort:        coalesceString(c2.Phoneshort, c1.Phoneshort),
		Phonetele:         coalesceString(c2.Phonetele, c1.Phonetele),
		Phoneteleall:      coalesceString(c2.Phoneteleall, c1.Phoneteleall),
		Phoneuseexternal:  coalesceUint32(c2.Phoneuseexternal, c1.Phoneuseexternal),
		Disability:        coalesceUint32(c2.Disability, c1.Disability),
		Fieldex1:          coalesceString(c2.Fieldex1, c1.Fieldex1),
		Fieldex2:          coalesceString(c2.Fieldex2, c1.Fieldex2),
		Cardidentityraw:   coalesceString(c2.Cardidentityraw, c1.Cardidentityraw),
		Refguid:           coalesceString(c2.Refguid, c1.Refguid),
		Selectpindatetime: coalesceString(c2.Selectpindatetime, c1.Selectpindatetime),
		Apireference:      coalesceString(c2.Apireference, c1.Apireference),
		Changedby:         coalesce(c2.Changedby, c1.Changedby),     //nolint:all
		Changeddate:       coalesce(c2.Changeddate, c1.Changeddate), //nolint:all
		Createdby:         coalesce(c2.Createdby, c1.Createdby),     //nolint:all
		Createddate:       coalesce(c2.Createddate, c1.Createddate), //nolint:all
		Asciicard:         coalesceBool(c2.Asciicard, c1.Asciicard),
		Showregister:      coalesceBool(c2.Showregister, c1.Showregister),
		Pinblocked:        coalesceBool(c2.Pinblocked, c1.Pinblocked),
		Isblocked:         coalesceBool(c2.Isblocked, c1.Isblocked),
		Inherituseraccess: coalesceBool(c2.Inherituseraccess, c1.Inherituseraccess),
		Expired:           coalesceBool(c2.Expired, c1.Expired),
		User:              coalesce(c2.User, c1.User),
		Parentcard:        coalesceUint32(c2.Parentcard, c1.Parentcard),
		Vacation:          coalesceUint32(c2.Vacation, c1.Vacation),

		//nolint:all
		Accessalarms: coalesce(c2.Accessalarms, c1.Accessalarms),

		//nolint:all
		Accessgroups: coalesceSlice(c2.Accessgroups, c1.Accessgroups),

		//nolint:all
		Readeraccess: coalesceSlice(c2.Readeraccess, c1.Readeraccess),
	}
}
