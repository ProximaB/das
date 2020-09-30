package partnershipdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_PARTNERSHIP_REQUEST_TABLE              = "DAS.PARTNERSHIP_REQUEST"
	DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID      = "SENDER_ID"
	DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID  = "RECIPIENT_ID"
	DAS_PARTNERSHIP_REQUEST_COL_SENDER_ROLE    = "SENDER_ROLE"
	DAS_PARTNERSHIP_REQUEST_COL_RECIPIENT_ROLE = "RECIPIENT_ROLE"
	DAS_PARTNERSHIP_REQUEST_COL_MESSAGE        = "MESSAGE"
	DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS = "REQUEST_STATUS"
)

type PostgresPartnershipRequestRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestRepository) SearchPartnershipRequest(criteria businesslogic.SearchPartnershipRequestCriteria) ([]businesslogic.PartnershipRequest, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	requests := make([]businesslogic.PartnershipRequest, 0)
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID,
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIENT_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_MESSAGE,
		DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).From(DAS_PARTNERSHIP_REQUEST_TABLE).OrderBy(common.ColumnPrimaryKey)

	if criteria.Sender > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID: criteria.Sender})
	}
	if criteria.Recipient > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID: criteria.Recipient})
	}
	if criteria.Sender == 0 && criteria.Recipient == 0 {
		return requests, errors.New("either sender or recipient must be specified")
	}
	if criteria.RequestStatusID > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS: criteria.RequestStatusID})
	}
	if criteria.RequestID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.RequestID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return requests, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipRequest{}
		rows.Scan(
			&each.PartnershipRequestID,
			&each.SenderID,
			&each.RecipientID,
			&each.SenderRole,
			&each.RecipientRole,
			&each.Message,
			&each.Status,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		requests = append(requests, each)
	}
	rows.Close()
	return requests, err
}

func (repo PostgresPartnershipRequestRepository) CreatePartnershipRequest(request *businesslogic.PartnershipRequest) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_PARTNERSHIP_REQUEST_TABLE).Columns(
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ID,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIEINT_ID,
		DAS_PARTNERSHIP_REQUEST_COL_SENDER_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_RECIPIENT_ROLE,
		DAS_PARTNERSHIP_REQUEST_COL_MESSAGE,
		DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		request.SenderID,
		request.RecipientID,
		request.SenderRole,
		request.RecipientRole,
		request.Message,
		request.Status,
		request.CreateUserID,
		request.DateTimeCreated,
		request.UpdateUserID,
		request.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&request.PartnershipRequestID)
		err = tx.Commit()
	}
	return err
}

func (repo PostgresPartnershipRequestRepository) UpdatePartnershipRequest(request businesslogic.PartnershipRequest) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SqlBuilder.Update("").
		Table(DAS_PARTNERSHIP_REQUEST_TABLE).
		Set(DAS_PARTNERSHIP_REQUEST_COL_REQUEST_STATUS, request.Status).
		Set(common.ColumnUpdateUserID, request.RecipientID).
		Set(common.ColumnDateTimeUpdated, request.DateTimeUpdated).
		Where(squirrel.Eq{common.ColumnPrimaryKey: request.PartnershipRequestID})

	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresPartnershipRequestRepository) DeletePartnershipRequest(request businesslogic.PartnershipRequest) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

func (repo PostgresPartnershipRequestRepository) GetReceivedRequests(recipientID int) ([]businesslogic.PartnershipRequest, error) {
	stmt := `SELECT * FROM GET_RECEIVED_PARTNERSHIP_REQUESTS($1);`
	rows, err := repo.Database.Query(stmt, recipientID)
	requests := make([]businesslogic.PartnershipRequest, 0)
	if err != nil {
		return requests, err
	}
	var scanErr error
	for rows.Next() {
		each := businesslogic.PartnershipRequest{
			SenderAccount:    new(businesslogic.Account),
			RecipientAccount: new(businesslogic.Account),
		}
		scanErr = rows.Scan(
			&each.PartnershipRequestID,
			&each.SenderAccount.ID,
			&each.SenderAccount.FirstName,
			&each.SenderAccount.LastName,
			&each.SenderAccount.Email,
			&each.SenderRole,
			&each.RecipientAccount.ID,
			&each.RecipientAccount.FirstName,
			&each.RecipientAccount.LastName,
			&each.RecipientAccount.Email,
			&each.RecipientRole,
			&each.Message,
			&each.Status,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		requests = append(requests, each)
		if scanErr != nil {
			return requests, scanErr
		}
	}
	err = rows.Close()
	return requests, err
}
func (repo PostgresPartnershipRequestRepository) GetSentRequests(senderID int) ([]businesslogic.PartnershipRequest, error) {
	stmt := `SELECT * FROM GET_SENT_PARTNERSHIP_REQUESTS($1);`
	rows, err := repo.Database.Query(stmt, senderID)
	requests := make([]businesslogic.PartnershipRequest, 0)
	if err != nil {
		return requests, err
	}
	var scanErr error
	for rows.Next() {
		each := businesslogic.PartnershipRequest{
			SenderAccount:    new(businesslogic.Account),
			RecipientAccount: new(businesslogic.Account),
		}
		scanErr = rows.Scan(
			&each.PartnershipRequestID,
			&each.SenderAccount.ID,
			&each.SenderAccount.FirstName,
			&each.SenderAccount.LastName,
			&each.SenderAccount.Email,
			&each.SenderRole,
			&each.RecipientAccount.ID,
			&each.RecipientAccount.FirstName,
			&each.RecipientAccount.LastName,
			&each.RecipientAccount.Email,
			&each.RecipientRole,
			&each.Message,
			&each.Status,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		requests = append(requests, each)
		if scanErr != nil {
			return requests, scanErr
		}
	}
	err = rows.Close()
	return requests, err
}
