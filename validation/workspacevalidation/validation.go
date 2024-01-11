package workspacevalidation

type FileRepo interface {
	DoesFileExist(hash string) (bool, error)
}

type OrganizationRepo interface {
	DoesOrganizationExist(organizationID uint) (bool, error)
}

type Validation struct {
	fileRepo FileRepo
	orgRepo  OrganizationRepo
}

func New(
	fileRepo FileRepo,
	orgRepo OrganizationRepo,
) Validation {
	return Validation{
		fileRepo: fileRepo,
		orgRepo:  orgRepo}
}
