--name: drop-students-table
DROP TABLE students;
--name: create-students-table

CREATE TABLE students (
  id INT(11) NOT NULL AUTO_INCREMENT,
  roll_id INT(11) NOT NULL,
  std_name VARCHAR(244) NOT NULL,
  ather_name VARCHAR(244) NOT NULL,
  mother_name VARCHAR(244) NOT NULL,
  board VARCHAR(244) NOT NULL,
  degree VARCHAR(244) NOT NULL,
  institute VARCHAR(244) NOT NULL,
  sub VARCHAR(244) NOT NULL,
  mark INT(3) NOT NULL,
  year INT(4) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
  PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- name: find-users-by-degree-year-board-roll
SELECT * FROM students WHERE roll_id = ? AND board = ? And degree = ? AND year = ?;
