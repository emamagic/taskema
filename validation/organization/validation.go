package orgvalidation

type FileRepo interface {
	DoesFileExist(hash string) (bool, error)
}

type OrgRepo interface {
	DoesOrganizationExist(organizationID uint) (bool, error)
}

type Validation struct {
	fileRepo FileRepo
	orgRepo  OrgRepo
}

func New(fileRepo FileRepo, orgRepo OrgRepo) Validation {
	return Validation{fileRepo: fileRepo, orgRepo: orgRepo}
}
