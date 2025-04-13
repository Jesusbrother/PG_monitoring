CREATE TABLE monitoring_logs (
    id SERIAL PRIMARY KEY,
    collected_at TIMESTAMP DEFAULT NOW(),
    data JSONB NOT NULL
);