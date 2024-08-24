CREATE TABLE IF NOT EXISTS posts
            (
              "id"       SERIAL       NOT NULL
                CONSTRAINT posts_pkey
                PRIMARY KEY,
              "text"    VARCHAR(255) NOT NULL,
              "created_at" integer(8)
            );"