package models

type TimePeriod struct {
	Periods []struct {
		Id    *uint32 `json:"id"`
		Text  *string `json:"text"`
		Days  *string `json:"days"`
		Times []struct {
			Start *string `json:"start"`
			End   *string `json:"end"`
		} `json:"times"`
	} `json:"periods"`
	Lock  bool `json:"lock"`
	Color bool `json:"color"`
}

type AccessGroup struct {
	Id           *uint32     `json:"id" validate:"gt=0" cmp:"skip"`
	Numbertype   *uint32     `json:"numbertype" validate:"required"`
	Name         *string     `json:"name" validate:"required"`
	Blocked      *uint32     `json:"blocked" validate:"required"`
	Startdate    *string     `json:"startdate" validate:"required"`
	Enddate      *string     `json:"enddate" validate:"required"`
	Timeperiod   *TimePeriod `json:"timeperiod" validate:"required"`
	Timecodetype *uint32     `json:"timecodetype" validate:"required"`
	Timecode     *uint32     `json:"timecode" validate:"required"`
	Alarmoff     *uint32     `json:"alarmoff" validate:"required"`
	Alarmon      *uint32     `json:"alarmon" validate:"required"`

	Elsassigned      *uint32 `json:"elsassigned" validate:"required"`
	Userassigned     *uint32 `json:"userassigned"`     // undocumentet
	Functionassigned *uint32 `json:"functionassigned"` // undocumentet
	Pincodeassigned  *uint32 `json:"pincodeassigned"`  // undocumentet

	Rasystem      *uint32 `json:"rasystem"`    // undocumentet
	Bookobjects   *string `json:"bookobjects"` // undocumentet
	Refguid       *string `json:"refguid" validate:"required"`
	Apidooraccess *uint32 `json:"apidooraccess" validate:"required"`

	Changedby   *string `json:"changedby" validate:"required" cmp:"skip"`
	Changeddate *string `json:"changeddate" validate:"required" cmp:"skip"`
	Createdby   *string `json:"createdby" validate:"required" cmp:"skip"`
	Createddate *string `json:"createddate" validate:"required" cmp:"skip"`
	Rastamp     *string `json:"rastamp" validate:"required" cmp:"skip"`

	Cards         []Card         `json:"cards"`
	Users         []User         `json:"users"`
	Units         []Unit         `json:"units"`
	Machinegroups []MachineGroup `json:"machinegroups"`
	Functions     []Function     `json:"functions"`

	//Activedays       *string     `json:"activedays" validate:"required"` // Removed in 1.5.4
	//Times            *string     `json:"times" validate:"required"` // Removed in 1.5.4
}

type AccessGroupList struct {
	Accessgroups []AccessGroup `json:"accessgroups"`
	Count        *int          `json:"count,omitempty"` // only included with ?count='true'
}
