-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_entries (
    id BIGSERIAL PRIMARY KEY,
    workout_id BIGINT NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    exercise_name VARCHAR(50) NOT NULL,
    sets INTEGER NOT NULL,
    reps INTEGER,
    duration_seconds INTEGER,
    weight DECIMAL(5,2)
    notes TEXT,
    order_index INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP

    CONSTRAINT valid_workout_entry CHECK (
        (reps is NOT NULL OR duration_seconds is NOT NULL) AND
        (reps is NULL OR duration_seconds is NULL)
    )


);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workout_entries;
-- +goose StatementEnd


