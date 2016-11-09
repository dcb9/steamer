package database

import "github.com/rubenv/sql-migrate"

var migrations = &migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		&migrate.Migration{
			Id: "1",
			Up: []string{`CREATE TABLE IF NOT EXISTS resource_info (
  id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  site VARCHAR(256) NULL DEFAULT NULL,
  title VARCHAR(1024) NULL DEFAULT NULL,
  url VARCHAR(2048) NULL DEFAULT NULL,
  streams TEXT NULL DEFAULT NULL,
  PRIMARY KEY (id))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4`},
			Down: []string{`DROP TABLE IF EXISTS resource_info`},
		},
		&migrate.Migration{
			Id: "2",
			Up: []string{`CREATE TABLE IF NOT EXISTS download_task (
  id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  resource_info_id INT(10) UNSIGNED NULL DEFAULT NULL,
  stream_index VARCHAR(16) NOT NULL,
  stdout TEXT NULL DEFAULT NULL,
  stderr TEXT NULL DEFAULT NULL,
  status TINYINT(4) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  INDEX fk_download_task_resource_info_idx (resource_info_id ASC),
  CONSTRAINT fk_download_task_resource_info
    FOREIGN KEY (resource_info_id)
    REFERENCES steamer.resource_info (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;`},
			Down: []string{`DROP TABLE IF EXISTS download_task`},
		},
	},
}