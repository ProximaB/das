-- name: psql-create-das-division
CREATE TABLE IF NOT EXISTS DAS.DIVISION (
	ID SERIAL NOT NULL PRIMARY KEY ,
	NAME VARCHAR(32) NOT NULL,
	ABBREVIATION VARCHAR(16) NOT NULL,
	DESCRIPTION TEXT,
	NOTE TEXT,
	FEDERATION_ID INTEGER NOT NULL REFERENCES DAS.FEDERATION(ID),
	DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
	DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
	UNIQUE(NAME, FEDERATION_ID)
);

INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
	VALUES ('Amateur', 'Amateur division at USA DANCE', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'USA DANCE') ,'AM','N/A');
INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
	VALUES ('Professional','Professional division at USA DANCE', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'USA DANCE'),'PRO','N/A');
INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
	VALUES ('Teacher/Student','Professional and Amateur at USA DANCE', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'USA DANCE'),'TS','N/A');
INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
	VALUES ('Collegiate','Collegiate division at Collegiate Competition', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'COL'),'COL','N/A');
INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
		VALUES ('Amateur', 'Amateur division at Independent/Unaffiliated Competition', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'IND'), 'AM', 'N/A');
INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
	VALUES ('WDSF','Division for all WDSF Evens at USA DANCE', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'WDSF'),'WDSF','N/A');
INSERT INTO DAS.DIVISION (NAME, DESCRIPTION, FEDERATION_ID, ABBREVIATION, NOTE)
  VALUES ('Amateur Championship', 'Amateur division at NDCA', (SELECT F.ID FROM DAS.FEDERATION F WHERE F.ABBREVIATION = 'NDCA'), 'AM', 'N/A');