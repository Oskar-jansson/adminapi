package adminapiinterface

import (
	"context"

	"github.com/Oskar-jansson/adminapi/models"
)

//
//	Generic interface for exposed services.
//	Copy into app.
//

type Auth interface {
	Login(ctx context.Context, credentials models.Credentials) error
	Logout(ctx context.Context) error
}

type AccessGroup interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.AccessGroup, error)
	List(ctx context.Context, params map[string]string) (*models.AccessGroupList, error)
}

type Administrator interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Administrator, error)
	List(ctx context.Context, params map[string]string) (*models.AdministratorList, error)
}

type Card interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Card, error)
	List(ctx context.Context, params map[string]string) (*models.CardList, error)
	Edit(ctx context.Context, id int, changes models.CardInput) (*models.Card, error)
	Create(ctx context.Context, card models.CardInput) (*models.Card, error)
	Delete(ctx context.Context, id int, params map[string]string) error
	AssignAccessGroup(ctx context.Context, cardId int, accessGroupId int) error
	RemoveAccessGroup(ctx context.Context, cardId int, accessGroupId int) error
}

type Connection interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Connection, error)
	List(ctx context.Context, params map[string]string) (*models.ConnectionList, error)
}

type Date interface {
	List(ctx context.Context, params map[string]string) (*models.DateList, error)
}

type Department interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Department, error)
	List(ctx context.Context, params map[string]string) (*models.DepartmentList, error)
	Create(ctx context.Context, newDepartment models.DepartmentInput) (*models.Department, error)
	Edit(ctx context.Context, id int, changes models.DepartmentInput) (*models.Department, error)
	Delete(ctx context.Context, id int) error
}

type Domain interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Domain, error)
	List(ctx context.Context, params map[string]string) (*models.DomainList, error)
}

type Event interface {
}

type Floor interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Floor, error)
	List(ctx context.Context, params map[string]string) (*models.FloorList, error)
	Edit(ctx context.Context, id int, changes models.FloorInput) (*models.Floor, error)
}

type Function interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Function, error)
	List(ctx context.Context, params map[string]string) (*models.FunctionList, error)
	Create(ctx context.Context, newFunction models.DepartmentInput) (*models.Function, error)
	Edit(ctx context.Context, id int, changes models.DepartmentInput) (*models.Function, error)
	Delete(ctx context.Context, id int) error
}

type MachineGroup interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.MachineGroup, error)
	List(ctx context.Context, params map[string]string) (*models.MachineGroupList, error)
}

type MachineGroupType interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.MachineGroupType, error)
	List(ctx context.Context, params map[string]string) (*models.MachineGroupTypeList, error)
}

type OfflineUnit interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.OfflineUnit, error)
	List(ctx context.Context, params map[string]string) (*models.OfflineUnitList, error)
	StepAccessVersion(ctx context.Context, id int) error
}

type Preselection interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Preselection, error)
	List(ctx context.Context, params map[string]string) (*models.PreselectionList, error)
}

type ReaderAccess interface {
	List(ctx context.Context, params map[string]string) (*models.ReaderAccessList, error)
	Create(ctx context.Context, cardId int, newReaderAccess models.ReaderAccessInput) (*models.ReaderAccess, error)
	Delete(ctx context.Context, cardId int, unitId int) error
}

type Setting interface {
	List(ctx context.Context, params map[string]string) (*models.SettingList, error)
}

type System interface {
	List(ctx context.Context, params map[string]string) (*models.System, error)
}

type Timezone interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Timezone, error)
	List(ctx context.Context, params map[string]string) (*models.TimezoneList, error)
}

type Unit interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.Unit, error)
	List(ctx context.Context, params map[string]string) (*models.UnitList, error)
}

type User interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.User, error)
	List(ctx context.Context, params map[string]string) (*models.UserList, error)
	Create(ctx context.Context, newUser models.UserInput) (*models.User, error)
	Edit(ctx context.Context, id int, changes models.UserInput) (*models.User, error)
	Delete(ctx context.Context, id int) error
	AssignAccessGroup(ctx context.Context, userId int, accessGroupId int) error
	RemoveAccessGroup(ctx context.Context, userId int, accessGroupId int) error
}

type UserGroup interface {
	Get(ctx context.Context, id int, params map[string]string) (*models.UserGroup, error)
	List(ctx context.Context, params map[string]string) (*models.UserGroupList, error)
	Create(ctx context.Context, newUserGroup models.DepartmentInput) (*models.UserGroup, error)
	Edit(ctx context.Context, id int, changes models.DepartmentInput) (*models.UserGroup, error)
	Delete(ctx context.Context, id int) error
}

type Version interface {
	Get(ctx context.Context) (*models.Version, error)
}
