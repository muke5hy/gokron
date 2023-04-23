CREATE TABLE cron_jobs (
    id SERIAL PRIMARY KEY,
    command TEXT NOT NULL,
    schedule VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE cron_logs (
    id SERIAL PRIMARY KEY,
    cron_job_id INTEGER NOT NULL REFERENCES cron_jobs(id),
    output TEXT,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);