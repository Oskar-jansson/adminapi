package models

type Function struct {
	Id            *uint32      `json:"id" validate:"gt=0" cmp:"skip"`
	Fkaccessgroup *uint32      `json:"fkaccessgroup"`
	Fkunit        *uint32      `json:"fkunit"`
	Times         *string      `json:"times" validate:"required"`
	Timeperiods   *string      `json:"timeperiods"`
	Activedays    *string      `json:"activedays"`
	Type          *uint32      `json:"type" validate:"required"`
	Comment       *string      `json:"comment"`
	Rastamp       *string      `json:"rastamp" validate:"required" cmp:"skip"`
	Unit          *Unit        `json:"unit"`
	Accessgroup   *AccessGroup `json:"Accessgroup"`
}

type FunctionList struct {
	Functions []Function `json:"functions"`
	Count     *int       `json:"count,omitempty"` // only included with ?count='true'
}

type FunctionInput struct {
	Fkaccessgroup *uint32      `json:"fkaccessgroup,omitempty"`
	Fkunit        *uint32      `json:"fkunit,omitempty"`
	Times         *string      `json:"times,omitempty"`
	Timeperiods   *string      `json:"timeperiods,omitempty"`
	Activedays    *string      `json:"activedays,omitempty"`
	Type          *uint32      `json:"type,omitempty"`
	Comment       *string      `json:"comment,omitempty"`
	Rastamp       *string      `json:"rastamp,omitempty"`
	Unit          *Unit        `json:"unit,omitempty"`
	Accessgroup   *AccessGroup `json:"Accessgroup,omitempty"`
}

func (f *Function) ConvToFunctionInput() *FunctionInput {
	return &FunctionInput{
		Fkaccessgroup: f.Fkaccessgroup,
		Fkunit:        f.Fkunit,
		Times:         f.Times,
		Timeperiods:   f.Timeperiods,
		Activedays:    f.Activedays,
		Type:          f.Type,
		Comment:       f.Comment,
		Rastamp:       f.Rastamp,
	}
}

func (f *FunctionInput) ConvToFunction() *Function {
	return &Function{
		Fkaccessgroup: f.Fkaccessgroup,
		Fkunit:        f.Fkunit,
		Times:         f.Times,
		Timeperiods:   f.Timeperiods,
		Activedays:    f.Activedays,
		Type:          f.Type,
		Comment:       f.Comment,
		Rastamp:       f.Rastamp,
	}
}

// f2 fields take precedence over f1 fields (if f2 field is non-nil, use it).
func (f1 *Function) MergeFunctions(f2 *Function) *Function {
	return &Function{
		Id:            coalesceUint32(f2.Id, f1.Id),
		Fkaccessgroup: coalesceUint32(f2.Fkaccessgroup, f1.Fkaccessgroup),
		Fkunit:        coalesceUint32(f2.Fkunit, f1.Fkunit),
		Times:         coalesceString(f2.Times, f1.Times),
		Timeperiods:   coalesceString(f2.Timeperiods, f1.Timeperiods),
		Activedays:    coalesceString(f2.Activedays, f1.Activedays),
		Type:          coalesceUint32(f2.Type, f1.Type),
		Comment:       coalesceString(f2.Comment, f1.Comment),
		Rastamp:       coalesce(f2.Rastamp, f1.Rastamp), //nolint:all
		Unit:          coalesce(f2.Unit, f1.Unit),
		Accessgroup:   coalesce(f2.Accessgroup, f1.Accessgroup),
	}
}
