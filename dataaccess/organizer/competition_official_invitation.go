package organizer

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	dasCompetitionOfficialInvitationTable = "DAS.COMPETITION_OFFICIAL_INVITATION"
	columnInvitationRecipientID           = "RECIPIENT_ID"
	columnRoleAssigned                    = "ROLE_ASSIGNED"
	columnMessage                         = "MESSAGE"
	columnStatus                          = "STATUS"
	columnExpirationDate                  = "EXPIRATION_DATE"
)

type PostgresCompetitionOfficialInvitationRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionOfficialInvitationRepository) CreateCompetitionOfficialInvitationRepository(invitation *businesslogic.CompetitionOfficialInvitation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(dasCompetitionOfficialInvitationTable).
		Columns(
			common.COL_ORGANIZER_ID,
			columnInvitationRecipientID,
			common.COL_COMPETITION_ID,
			columnRoleAssigned,
			columnMessage,
			columnStatus,
			columnExpirationDate,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			invitation.Sender.ID,
			invitation.Recipient.ID,
			invitation.ServiceCompetition.ID,
			invitation.AssignedRoleID,
			invitation.Message,
			invitation.InvitationStatus,
			invitation.ExpirationDate,
			invitation.CreateUserId,
			invitation.DateTimeCreated,
			invitation.UpdateUserId,
			invitation.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)
	hasError := false
	clause, args, sqlErr := stmt.ToSql()
	if sqlErr != nil {
		log.Printf("[error] generating SQL clause: %v", sqlErr)
		hasError = true
	}
	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		log.Printf("[error] starting transaction: %v", txErr)
		hasError = true
	}
	row := repo.Database.QueryRow(clause, args...)
	if scanErr := row.Scan(&invitation.ID); scanErr != nil {
		log.Printf("[error] scanning ID of newly created Competition Official Invitation: %v", scanErr)
		hasError = true
	}
	if commitErr := tx.Commit(); commitErr != nil {
		log.Printf("[error] commiting transaction: %v", commitErr)
		hasError = true
	}
	if hasError {
		return errors.New("An error occurred while creating competition official invitation record")
	}
	return nil
}
func (repo PostgresCompetitionOfficialInvitationRepository) DeleteCompetitionOfficialInvitationRepository(invitation businesslogic.CompetitionOfficialInvitation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
func (repo PostgresCompetitionOfficialInvitationRepository) SearchCompetitionOfficialInvitationRepository(criteria businesslogic.SearchCompetitionOfficialInvitationCriteria) ([]businesslogic.CompetitionOfficialInvitation, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_ORGANIZER_ID,
			columnInvitationRecipientID,
			common.COL_COMPETITION_ID,
			columnRoleAssigned,
			columnMessage,
			columnStatus,
			columnExpirationDate,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated,
		)).From(dasCompetitionOfficialInvitationTable)
	if criteria.AssignedRoleID > 0 {
		stmt = stmt.Where(squirrel.Eq{columnRoleAssigned: criteria.AssignedRoleID})
	}
	if criteria.SenderID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_ORGANIZER_ID: criteria.SenderID})
	}
	if criteria.RecipientID > 0 {
		stmt = stmt.Where(squirrel.Eq{columnInvitationRecipientID: criteria.RecipientID})
	}
	if criteria.CreateUserID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnCreateUserID: criteria.CreateUserID})
	}

	invitations := make([]businesslogic.CompetitionOfficialInvitation, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := businesslogic.CompetitionOfficialInvitation{
			Sender:             businesslogic.Account{},
			Recipient:          businesslogic.Account{},
			ServiceCompetition: businesslogic.Competition{},
		}
		scanErr := rows.Scan(
			&each.ID,
			&each.Sender.ID,
			&each.Recipient.ID,
			&each.ServiceCompetition.ID,
			&each.AssignedRoleID,
			&each.Message,
			&each.InvitationStatus,
			&each.ExpirationDate,
			&each.CreateUserId,
			&each.DateTimeCreated,
			&each.UpdateUserId,
			&each.DateTimeUpdated)
		if scanErr != nil {
			log.Printf("[error] scanning Competition Official Invitation: %v", scanErr)
			return invitations, err
		}
		invitations = append(invitations, each)
	}
	rows.Close()
	return invitations, err
}
func (repo PostgresCompetitionOfficialInvitationRepository) UpdateCompetitionOfficialInvitationRepository(invitation businesslogic.CompetitionOfficialInvitation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
