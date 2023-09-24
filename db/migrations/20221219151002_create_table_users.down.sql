CREATE TABLE
  users (
    id INT NOT NULL AUTO_INCREMENT,
    full_name VARCHAR(10),
    email VARCHAR(10),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
  ) ENGINE = InnoDB;
