CREATE TABLE urls (
                      id BIGSERIAL PRIMARY KEY,
                      original_url TEXT NOT NULL,
                      short_code VARCHAR(20) NOT NULL UNIQUE,
                      created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                      updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                      expires_at TIMESTAMP NULL,
                      clicks BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_urls_short_code ON urls(short_code);