CREATE TABLE IF NOT EXISTS clips(
    id    INTEGER   PRIMARY KEY,
    value TEXT      NOT NULL, 
    saved DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP
);
