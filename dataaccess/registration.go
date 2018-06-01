package dataaccess

/*
func GetCompetitiveBallroomEventRegistration(db *sql.DB, competitionID int, partnershipID int) (businesslogic.CompetitiveBallroomEventRegistration, error) {
	clause := SQLBUILDER.Select(`CE.COMPETITION_ID, ECBE.PARTNERSHIP_ID,
		CR.COUNTRY_ID,
		CR.STATE_ID,
		CR.SCHOOL_ID,
		CR.STUDIO_ID`).
		From("DAS.COMPETITION_ENTRY CE").
		Join("DAS.COMPETITION_REPRESENTATION CR ON CR.COMPETITION_ENTRY_ID = CE.ID").
		Join("DAS.EVENT E ON E.COMPETITION_ID = CE.COMPETITION_ID").
		Join("DAS.EVENT_COMPETITIVE_BALLROOM ECB ON ECB.EVENT_ID = E.ID").
		Join("DAS.EVENT_COMPETITIVE_BALLROOM_ENTRY ECBE ON ECBE.COMPETITIVE_BALLROOM_EVENT_ID = ECB.ID").
		Where(sq.Eq{"CE.COMPETITION_ID": competitionID}).
		Where(sq.Eq{"ECBE.PARTNERSHIP_ID": partnershipID})

	registration := businesslogic.CompetitiveBallroomEventRegistration{
		EventsAdded: make([]int, 0),
	}

	stmt, args, _ := clause.ToSql()
	if tx, txErr := db.Begin(); txErr != nil {
		return registration, txErr
	} else {
		tx.QueryRow(stmt, args...).Scan(
			&registration.CompetitionID,
			&registration.PartnershipID,
			&registration.CountryRepresented,
			&registration.StateRepresented,
			&registration.SchoolRepresented,
			&registration.StudioRepresented,
		)
		tx.Commit()
	}

	clause = SQLBUILDER.
		Select("ECBE.COMPETITIVE_BALLROOM_EVENT_ID").
		From("DAS.EVENT_COMPETITIVE_BALLROOM_ENTRY ECBE").
		Join("DAS.EVENT_COMPETITION_BALLROOM ECB ON ECBE.COMPETITIVE_BALLROOM_EVENT_ID = ECB.ID").
		Join("DAS.EVENT E ON ECB.EVENT_ID E.ID").
		Where(sq.Eq{"ECBE.PARTNERSHIP_ID": partnershipID}).
		Where(sq.Eq{"E.COMPETITION_ID": competitionID})

	rows, _ := clause.RunWith(db).Query()
	for rows.Next() {
		event := 0
		rows.Scan(
			&event,
		)
		registration.EventsAdded = append(registration.EventsAdded, event)
	}
	// find all the events entered
	return registration, nil
}
*/
