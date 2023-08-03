CREATE TABLE Users
(
    id SERIAL PRIMARY KEY,
    phone           VARCHAR(13) UNIQUE NOT NULL,
    fullName        VARCHAR(60),
    password        VARCHAR(64)
);

INSERT INTO Users (phone, fullName, password)
VALUES ('+621234567890', 'John Doe', 'hashedpassword1');

INSERT INTO Users (phone, fullName, password)
VALUES ('+629876543210', 'Jane Smith', 'hashedpassword2');

INSERT INTO Users (phone, fullName, password)
VALUES ('+628765432109', 'Michael Johnson', 'hashedpassword3');

INSERT INTO Users (phone, fullName, password)
VALUES ('+621234567891', 'Emily Williams', 'hashedpassword4');

INSERT INTO Users (phone, fullName, password)
VALUES ('+628765432108', 'David Lee', 'hashedpassword5');

INSERT INTO Users (phone, fullName, password)
VALUES ('+628765432107', 'Sarah Brown', 'hashedpassword6');

INSERT INTO Users (phone, fullName, password)
VALUES ('+621234567892', 'Robert Martin', 'hashedpassword7');

INSERT INTO Users (phone, fullName, password)
VALUES ('+628765432106', 'Karen Johnson', 'hashedpassword8');

INSERT INTO Users (phone, fullName, password)
VALUES ('+628765432105', 'Daniel Wilson', 'hashedpassword9');

INSERT INTO Users (phone, fullName, password)
VALUES ('+628765432104', 'Michelle Anderson', 'hashedpassword10');
