CREATE TABLE IF NOT EXISTS Codes (
    CodeID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    CodeType VARCHAR(255),
    Code VARCHAR(255)
);