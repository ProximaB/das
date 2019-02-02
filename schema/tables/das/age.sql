CREATE TABLE IF NOT EXISTS DAS.AGE (
  ID SERIAL PRIMARY KEY NOT NULL,
  NAME VARCHAR (16) NOT NULL,
  DESCRIPTION TEXT,
  DIVISION_ID INTEGER NOT NULL REFERENCES DAS.DIVISION(ID),
  ENFORCED BOOLEAN NOT NULL DEFAULT FALSE,
  MINIMUM_AGE INTEGER NOT NULL DEFAULT (-1),
  MAXIMUM_AGE INTEGER NOT NULL DEFAULT (-1),
  CREATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (DIVISION_ID, NAME),
  CHECK (MINIMUM_AGE <= MAXIMUM_AGE)
);
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID) 
  VALUES('Pre Teen I','First level of Pre-Teen',true,0,9, 
         (SELECT D.ID FROM DAS.DIVISION D 
           JOIN DAS.FEDERATION F 
             ON D.FEDERATION_ID = F.ID 
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID) 
  VALUES('Pre Teen II','Second level of Pre-Teen',true,10,11,
         (SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Junior I','First level of Junior',true,12,13,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID) 
  VALUES('Junior II','Second level of Junior',true,14,15,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Youth','Youth level',true,16,18,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('U21','Under 21 Adult and Youth',true,16,20,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Adult','All adult',true,19,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Senior I','First level of Senior',true,35,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Senior II','Second level of Senior',true,45,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Senior III','Third level of Senior',true,55,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Senior IV','Fourth level of Senior',true,65,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Senior V','Fifth level of Senior',true,75,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'USA DANCE'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
VALUES('Collegiate','Default Collegiate Age',false,0,99,(SELECT D.ID FROM DAS.DIVISION D
           JOIN DAS.FEDERATION F
             ON D.FEDERATION_ID = F.ID
         WHERE D.ABBREVIATION = 'COL' AND F.ABBREVIATION = 'COL'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Pre-Teen I','First level of Pre-Teen',true,0,9,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Pre-Teen II','Second level of Pre-Teen',true,10,11,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Pre-Teen','Comprehensive level of Pre-Teen',true,0,11,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Junior I','First level of Junior',true,12,13,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Junior II','Second level of Junior',true,14,15,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Junior','Comprehensive level of Junior',true,12,15,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Youth','Youth competitors of NDCA',true,16,18,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Under 21','Amateur competitors under 21 in NDCA',true,16,21,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Adult','Adult competitors in NDCA',true,19,99,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Senior I','First level of Senior',true,35,99,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Senior II','Second level of Senior',true,45,99,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Senior III','Third level of Senior',true,55,99,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));
INSERT INTO DAS.AGE(NAME, DESCRIPTION, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, DIVISION_ID)
  VALUES('Senior','Comprehensive level of Senior',true,35,99,
    (SELECT D.ID FROM DAS.DIVISION D
      JOIN DAS.FEDERATION F
        ON D.FEDERATION_ID = F.ID
      WHERE D.ABBREVIATION = 'AM' AND F.ABBREVIATION = 'NDCA'));