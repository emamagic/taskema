package orgrepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type orgSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) Organization {
	return &orgSqlRepo{db: db}
}

func (repo *orgSqlRepo) CreateOrganization(org entity.Organization) (uint, error) {
	op := "orgsql.CreateOrganization"

	tx, tErr := repo.db.Conn().Begin()
	if tErr != nil {

		return 0, richerror.New(op).WithError(tErr).WithCode(richerror.CodeUnexpected)
	}

	result, err := tx.Exec("INSERT INTO organizations (title, avatar, creator_user_id) VALUES (?, ?, ?)",
		org.Title, org.Avatar, org.CreatorUserID)
	if err != nil {
		tx.Rollback()
		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	orgID, lErr := result.LastInsertId()
	if lErr != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	_, eErr := tx.Exec("INSERT INTO user_organization (user_id, organization_id) VALUES (?, ?)", org.CreatorUserID, orgID)
	if eErr != nil {
		tx.Rollback()
		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	if err := tx.Commit(); err != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return uint(orgID), nil
}

func (repo *orgSqlRepo) GetAllOrganizationByUserID(userID uint) ([]entity.Organization, error) {
	op := "orgsql.GetAllOrganization"

	organizations := make([]entity.Organization, 0)

	rows, qErr := repo.db.Conn().Query(
		"SELECT organizations.id, organizations.title, organizations.avatar, organizations.creator_user_id, organizations.create_at, organizations.update_at "+
			"FROM organizations "+
			"JOIN user_organization ON organizations.id = user_organization.organization_id "+
			"JOIN users ON users.id = user_organization.user_id "+
			"WHERE users.id = ?",
		userID,
	)
	if qErr != nil {

		return nil,
			richerror.New(op).
				WithMessage(qErr.Error()).
				WithCode(richerror.CodeUnexpected).
				WithMeta(map[string]interface{}{"meta": "query failed"})
	}
	defer rows.Close()

	for rows.Next() {
		var organization entity.Organization

		err := rows.Scan(
			&organization.ID,
			&organization.Title,
			&organization.Avatar,
			&organization.CreatorUserID,
			&organization.CreateAt,
			&organization.UpdateAt,
		)
		if err != nil {

			return nil,
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeUnexpected).
					WithMeta(map[string]interface{}{"meta": "one row failed"})
		}
		organizations = append(organizations, organization)
	}

	if rows.Err() != nil {

		return nil,
			richerror.New(op).
				WithMessage(rows.Err().Error()).
				WithCode(richerror.CodeUnexpected).
				WithMeta(map[string]interface{}{"meta": "error happend while scanning one of the row of organizations"})
	}

	return organizations, nil
}

func (repo *orgSqlRepo) DeleteOrganizationByID(orgID uint) error {
	op := "orgsql.DeleteOrganizationByID"

	result, err := repo.db.Conn().Exec("DELETE FROM organizations WHERE id = ?", orgID)
	if err != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	count, aErr := result.RowsAffected()
	if aErr != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	if count == 0 {

		return richerror.New(op).WithMessage("no such organization exist").WithCode(richerror.CodeNotFound)
	}

	return nil
}

func (repo *orgSqlRepo) GetOrganizationByID(orgID uint) (entity.Organization, error) {
	op := "orgsql.GetOrganizationByID"

	row := repo.db.Conn().QueryRow("SELECT * FROM organizations WHERE id = ?", orgID)

	var organization entity.Organization
	if err := row.Scan(
		&organization.ID,
		&organization.Title,
		&organization.Avatar,
		&organization.CreatorUserID,
		&organization.CreateAt,
		&organization.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {

			return entity.Organization{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}
		return entity.Organization{},
			richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return organization, nil
}
