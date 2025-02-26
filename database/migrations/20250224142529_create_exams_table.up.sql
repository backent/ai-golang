CREATE TABLE exams (
  id BIGINT NOT NULL AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  question_id BIGINT NOT NULL,
  score INT NOT NULL,
  submissions TEXT NOT NULL,
  exam_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX exams_username (username),
  INDEX exams_question_id (question_id)
)