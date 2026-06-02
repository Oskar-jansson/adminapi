package models

type Department struct {
	Id             *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Departmentname *string `json:"departmentname" validate:"required"`
	Rastamp        *string `json:"rastamp" validate:"required" cmp:"skip"`
	Departmentdata *string `json:"departmentdata"`
	Users          []User  `json:"users"`
}

type DepartmentInput struct {
	Departmentname *string `json:"departmentname,omitempty"`
	Rastamp        *string `json:"rastamp,omitempty"`
	Departmentdata *string `json:"departmentdata,omitempty"`
}

type DepartmentList struct {
	Departments []Department `json:"departments"`
	Count       *int         `json:"count,omitempty"` // only included with ?count='true'
}

func (u Department) ConvToDepartmentInput() *DepartmentInput {
	return &DepartmentInput{

		Departmentname: u.Departmentname,
		Rastamp:        u.Rastamp,
		Departmentdata: u.Departmentdata,
	}
}

func (u DepartmentInput) ConvToDepartment() *Department {
	return &Department{

		Departmentname: u.Departmentname,
		Rastamp:        u.Rastamp,
		Departmentdata: u.Departmentdata,
	}
}

// u2 fields take precedence over u1 fields (if u2 field is non-nil, use it).
func (u1 *Department) MergeDepartments(u2 *Department) *Department {
	return &Department{
		Id:             coalesceUint32(u2.Id, u1.Id),
		Departmentname: coalesceString(u2.Departmentname, u1.Departmentname),
		Rastamp:        coalesce(u2.Rastamp, u1.Rastamp), //nolint:all
		Departmentdata: coalesceString(u2.Departmentdata, u1.Departmentdata),
	}
}
