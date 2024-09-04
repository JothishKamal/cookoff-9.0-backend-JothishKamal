CREATE TABLE "user" (
	"id" UUID NOT NULL UNIQUE,
	"email" TEXT NOT NULL UNIQUE, 
	"regNo" TEXT NOT NULL UNIQUE,
	"password" TEXT NOT NULL,
	"role" TEXT NOT NULL,
	"roundQualified" INTEGER NOT NULL,
	"score" INTEGER,
	"name" TEXT NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "questions" (
	"id" UUID NOT NULL UNIQUE,
	"description" TEXT,
	"title" TEXT,
	"inputFormat" TEXT,
	"points" INTEGER,
	"round" INTEGER NOT NULL,
	"constraints" TEXT,
	"outputFormat" TEXT,
	"testcases" UUID,
	PRIMARY KEY("id")
);

CREATE TABLE "submissions" (
	"id" UUID NOT NULL UNIQUE,
	"question_id" UUID NOT NULL,
	"testcases_passed" INTEGER,
	"testcases_failed" INTEGER,
	"runtime" DECIMAL,
	"sub time" TIMESTAMP,
	"testcases_id" UUID,
	"language_id" INTEGER NOT NULL,
	"description" TEXT,
	"memory" INTEGER,
	"user_id" UUID,
	PRIMARY KEY("id")
);

CREATE TABLE "testcases" (
	"id" UUID NOT NULL UNIQUE,
	"expected_output" TEXT,
	"memory" TEXT,
	"input" TEXT,
	"hidden" BOOLEAN,
	"runtime" TIME,
	PRIMARY KEY("id")
);

ALTER TABLE "submissions"
ADD FOREIGN KEY("question_id") REFERENCES "questions"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "questions"
ADD FOREIGN KEY("testcases") REFERENCES "testcases"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "submissions"
ADD FOREIGN KEY("testcases_id") REFERENCES "testcases"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "submissions"
ADD FOREIGN KEY("user_id") REFERENCES "user"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
