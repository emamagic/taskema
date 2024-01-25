package workspacevalidation

type FileRepo interface {
	DoesFileExist(hash string) (bool, error)
}

type OrganizationRepo interface {
	DoesOrganizationExist(organizationID uint) (bool, error)
}

type WorkspaceRepo interface {
	DoesWorkspaceExist(workspaceID uint) (bool, error)
}

type Validation struct {
	fileRepo      FileRepo
	orgRepo       OrganizationRepo
	workspaceRepo WorkspaceRepo
}

func New(
	fileRepo FileRepo,
	orgRepo OrganizationRepo,
	workspaceRepo WorkspaceRepo,
) Validation {
	return Validation{
		fileRepo:      fileRepo,
		orgRepo:       orgRepo,
		workspaceRepo: workspaceRepo}
}
