SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

# 外键约束由代码逻辑层完成
########## 老师表 #############
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `number` varchar(255) NOT NULL COMMENT '工号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE  KEY (`number`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 学生表 #############
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `number` varchar(255) NOT NULL COMMENT '学号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `klass_id` INT NOT NULL COMMENT '所属班级id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE  KEY (`number`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 班级表 #############
DROP TABLE IF EXISTS `klass`;
CREATE TABLE `klass` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(255) NOT NULL COMMENT '名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 题目表 #############
# 设计的非常sb，可以在根据不同题目类型建表，这里就算了
DROP TABLE IF EXISTS `question`;
CREATE TABLE `question` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `desc` varchar(255) NOT NULL COMMENT '题目描述',
  `answer` varchar(255) NOT NULL COMMENT '题目答案',
  `type` INT NOT NULL COMMENT '题目类型：0 未知，1 单选，2多选',
  `option_desc_A` varchar(255) COMMENT '选项 A，选项内容为空，即没有该选项',
  `option_desc_B` varchar(255) COMMENT '选项 B',
  `option_desc_C` varchar(255) COMMENT '选项 C',
  `option_desc_D` varchar(255) COMMENT '选项 D',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 考卷表 #############
DROP TABLE IF EXISTS `paper`;
CREATE TABLE `paper` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(255) NOT NULL COMMENT '考卷名称',
  `score_limit` INT NOT NULL COMMENT '分数上限',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 考试表 #############
DROP TABLE IF EXISTS `exam`;
CREATE TABLE `exam` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(255) NOT NULL COMMENT '考试名称',
  `begin_time` TIMESTAMP NOT NULL COMMENT '考试开始时间',
  `length` INT NOT NULL COMMENT '考试时长',
  `paper_id` INT NOT NULL COMMENT '考卷id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 老师-班级表 #############
DROP TABLE IF EXISTS `teacher_klass`;
CREATE TABLE `teacher_klass` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `teacher_id` INT NOT NULL COMMENT '老师id',
  `klass_id` INT NOT NULL COMMENT '班级id',

  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_teacher_id_klass_id`(`teacher_id`, `klass_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 班级-考试表 #############
DROP TABLE IF EXISTS `klass_exam`;
CREATE TABLE `klass_exam` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `klass_id` INT NOT NULL COMMENT '班级id',
  `exam_id` INT NOT NULL COMMENT '考试id',

  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_klass_id_exam_id`(`klass_id`, `exam_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;


########## 考卷-题目表 #############
DROP TABLE IF EXISTS `paper_question`;
CREATE TABLE `paper_question` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `paper_id` INT NOT NULL COMMENT '考卷id',
  `question_id` INT NOT NULL COMMENT '题目id',
  `question_score` INT NOT NULL COMMENT '题目得分',

  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_paper_id_question_id`(`paper_id`, `question_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;


########## 考试结果表 #############
DROP TABLE IF EXISTS `exam_result`;
CREATE TABLE `exam_result` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `exam_id` INT NOT NULL COMMENT '考试id',
  `student_id` INT NOT NULL COMMENT '学生id',
  `score` INT NOT NULL COMMENT '得分',

  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_exam_id_student_id`(`exam_id`, `student_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

########## 作答记录表 #############
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `student_id` INT NOT NULL COMMENT '学生id',
  `exam_id` INT NOT NULL COMMENT '考试id',
  `paper_id` INT NOT NULL COMMENT '考卷id',
  `question_id` INT NOT NULL COMMENT '题目id',
  `score` INT NOT NULL COMMENT '得分',
  `desc` varchar(255) NULL COMMENT '作答内容',

  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_key`(`student_id`, `exam_id`, `paper_id`, `question_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

##### 初始化数据
INSERT INTO teacher(name,number,password) VALUES('老师1','1','1');
INSERT INTO teacher(name,number,password) VALUES('老师2','2','2');
INSERT INTO klass(name) VALUES('班级1');
INSERT INTO klass(name) VALUES('班级2');
INSERT INTO klass(name) VALUES('班级3');
INSERT INTO klass(name) VALUES('班级4');
INSERT INTO teacher_klass(teacher_id,klass_id) VALUES(1,1);
INSERT INTO teacher_klass(teacher_id,klass_id) VALUES(1,2);
INSERT INTO teacher_klass(teacher_id,klass_id) VALUES(2,3);
INSERT INTO teacher_klass(teacher_id,klass_id) VALUES(2,4);
INSERT INTO student(name,number,password,klass_id) VALUES('学生1','1','1',1);
INSERT INTO student(name,number,password,klass_id) VALUES('学生2','2','2',1);
INSERT INTO student(name,number,password,klass_id) VALUES('学生3','3','3',2);
INSERT INTO student(name,number,password,klass_id) VALUES('学生4','4','4',3);
INSERT INTO student(name,number,password,klass_id) VALUES('学生5','5','5',4);
INSERT INTO question(`desc`,answer,type,option_desc_A,option_desc_B,option_desc_C,option_desc_D) VALUES('这是单选题题目描述','{"A"}',1,'答案A','B','C','D');
INSERT INTO question(`desc`,answer,type,option_desc_A,option_desc_B,option_desc_C,option_desc_D) VALUES('这是多选题题目描述','{"A,B"}',2,'答案A','答案B','C','D');