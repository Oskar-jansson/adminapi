package models

type MachineGroupType struct {
	Id             *uint32       `json:"id" validate:"gt=0" cmp:"skip"`
	Rastamp        *string       `json:"rastamp" validate:"required" cmp:"skip"`
	Name           *string       `json:"name" validate:"required"`
	Fkpreselection *uint32       `json:"fkpreselection" validate:"required"`
	Fkcategorie    *uint32       `json:"fkcategorie" validate:"required"`
	Localid        *uint32       `json:"localid" validate:"required"`
	Fkconnr        *uint32       `json:"fkconnr" validate:"required"`
	Rasystem       *uint32       `json:"rasystem" validate:"required"`
	Preselection   *Preselection `json:"preselection"`
	Machinegroups  *MachineGroup `json:"machinegroups"`
}

type MachineGroupTypeList struct {
	Machinegrouptypes []MachineGroupType `json:"machinegrouptypes"`
	Count             *int               `json:"count,omitempty"` // only included with ?count='true'
}
