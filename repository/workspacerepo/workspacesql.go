package workspacerepo

import (
	"database/sql"
	"taskema/datasource/mysql"
	"taskema/entity"
	"taskema/pkg/richerror"
)

type workspaceSqlRepo struct {
	db *mysql.MYSQL
}

func New(db *mysql.MYSQL) Workspace {
	return &workspaceSqlRepo{db: db}
}

// TODO - priority not implemented yet
func (repo *workspaceSqlRepo) CreateWorkspace(workspace entity.Workspace) (uint, error) {
	op := "workspacesql.CreateWorkspace"

	tx, tErr := repo.db.Conn().Begin()
	if tErr != nil {

		return 0, richerror.New(op).WithError(tErr).WithCode(richerror.CodeUnexpected)
	}
	result, err := tx.Exec("INSERT INTO workspaces (title, avatar, creator_user_id, organization_id) VALUES (?, ?, ?, ?)",
		workspace.Title, workspace.Avatar, workspace.CreatorUserID, workspace.OrganizationID)
	if err != nil {
		tx.Rollback()
		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	workspaceID, lErr := result.LastInsertId()
	if lErr != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}
	_, eErr := tx.Exec("INSERT INTO user_workspace (user_id, workspace_id) VALUES (?, ?)", workspace.CreatorUserID, workspaceID)
	if eErr != nil {
		tx.Rollback()
		return 0, richerror.New(op).WithError(eErr).WithCode(richerror.CodeUnexpected)
	}

	if err := tx.Commit(); err != nil {

		return 0, richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return uint(workspaceID), nil
}

func (repo *workspaceSqlRepo) GetAllWorkspaceByOrganizationID(organizationID uint, userID uint) ([]entity.Workspace, error) {
	op := "workspacesql.GetAllWorkspaceByOrganizationID"

	rows, qErr := repo.db.Conn().Query(
		"SELECT workspaces.id, workspaces.title, workspaces.avatar, workspaces.organization_id, workspaces.create_at, workspaces.update_at "+
			"FROM workspaces "+
			"JOIN user_workspace ON workspaces.id = user_workspace.workspace_id "+
			"JOIN users ON user_workspace.user_id = users.id "+
			"WHERE users.id = ? AND workspaces.organization_id = ?",
		userID, organizationID)
	if qErr != nil {

		return nil,
			richerror.New(op).
				WithMessage(qErr.Error()).
				WithCode(richerror.CodeUnexpected)
	}
	defer rows.Close()

	list := make([]entity.Workspace, 0)
	for rows.Next() {
		var workspace entity.Workspace
		err := rows.Scan(
			&workspace.ID,
			&workspace.Title,
			&workspace.Avatar,
			&workspace.OrganizationID,
			&workspace.CreateAt,
			&workspace.UpdateAt,
		)
		if err != nil {

			return nil,
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeUnexpected)
		}
		list = append(list, workspace)
	}

	if rows.Err() != nil {

		return nil,
			richerror.New(op).
				WithMessage(rows.Err().Error()).
				WithCode(richerror.CodeUnexpected).
				WithMeta(map[string]interface{}{"meta": "error happend while scanning one of the row of organizations"})
	}

	return list, nil
}

func (repo *workspaceSqlRepo) DeleteWorkspaceByID(workspaceID uint) error {
	op := "workspacesql.DeleteWorkspaceByID"

	result, err := repo.db.Conn().Exec("DELETE FROM workspaces WHERE id = ?", workspaceID)
	if err != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	count, aErr := result.RowsAffected()
	if aErr != nil {

		return richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	if count == 0 {

		return richerror.New(op).WithMessage("no such workspace exist").WithCode(richerror.CodeNotFound)
	}

	return nil
}

func (repo *workspaceSqlRepo) GetWorkspaceByID(workspaceID uint) (entity.Workspace, error) {
	op := "workspacesql.GetWorkspaceByID"

	row := repo.db.Conn().QueryRow("SELECT * FROM workspaces WHERE id = ?", workspaceID)

	var workspace entity.Workspace
	if err := row.Scan(
		&workspace.ID,
		&workspace.Title,
		&workspace.Avatar,
		&workspace.CreatorUserID,
		&workspace.OrganizationID,
		&workspace.Priority,
		&workspace.CreateAt,
		&workspace.UpdateAt,
	); err != nil {
		if err == sql.ErrNoRows {

			return entity.Workspace{},
				richerror.New(op).
					WithMessage(err.Error()).
					WithCode(richerror.CodeNotFound)
		}
		return entity.Workspace{},
			richerror.New(op).WithError(err).WithCode(richerror.CodeUnexpected)
	}

	return workspace, nil
}
