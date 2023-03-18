CREATE TABLE IF NOT EXISTS "stat" (
    "id" UUID NOT NULL UNIQUE,
    "user_id" BIGINT NOT NULL REFERENCES "user"("user_id"),
    "message" TEXT,
    "created_at" TIMESTAMP NOT NULL
)