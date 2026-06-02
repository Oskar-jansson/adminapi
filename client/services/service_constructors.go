package services

func NewCardService(c clientInterface) *CardService {
	return &CardService{sc: c}
}

func NewAdministratorService(c clientInterface) *AdministratorService {
	return &AdministratorService{sc: c}
}

func NewAuthService(c clientInterface) *AuthService {
	return &AuthService{sc: c}
}

func NewAccessGroupService(c clientInterface) *AccessGroupService {
	return &AccessGroupService{sc: c}
}

func NewConnectionService(c clientInterface) *ConnectionService {
	return &ConnectionService{sc: c}
}

func NewDateService(c clientInterface) *DateService {
	return &DateService{sc: c}
}

func NewDomainService(c clientInterface) *DomainService {
	return &DomainService{sc: c}
}

func NewFunctionService(c clientInterface) *FunctionService {
	return &FunctionService{sc: c}
}

func NewFloorService(c clientInterface) *FloorService {
	return &FloorService{sc: c}
}

func NewMachineGroupService(c clientInterface) *MachineGroupService {
	return &MachineGroupService{sc: c}
}

func NewMachineGroupTypeService(c clientInterface) *MachineGroupTypeService {
	return &MachineGroupTypeService{sc: c}
}

func NewOfflineUnitService(c clientInterface) *OfflineUnitService {
	return &OfflineUnitService{sc: c}
}

func NewPreselectionService(c clientInterface) *PreselectionService {
	return &PreselectionService{sc: c}
}

func NewReaderAccessService(c clientInterface) *ReaderAccessService {
	return &ReaderAccessService{sc: c}
}

func NewSettingService(c clientInterface) *SettingService {
	return &SettingService{sc: c}
}

func NewSystemService(c clientInterface) *SystemService {
	return &SystemService{sc: c}
}

func NewTimezoneService(c clientInterface) *TimezoneService {
	return &TimezoneService{sc: c}
}

func NewUnitService(c clientInterface) *UnitService {
	return &UnitService{sc: c}
}

func NewUserService(c clientInterface) *UserService {
	return &UserService{sc: c}
}

func NewUserGroupService(c clientInterface) *UserGroupService {
	return &UserGroupService{sc: c}
}

func NewDepartmentService(c clientInterface) *DepartmentService {
	return &DepartmentService{sc: c}
}

func NewEventService(c clientInterface) *EventService {
	return &EventService{sc: c}
}
