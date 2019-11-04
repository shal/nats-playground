CREATE TABLE wanted_vehicles(
    id TEXT UNIQUE NOT NULL,
    ovd TEXT,
    brand TEXT,
    model TEXT,
    kind TEXT,
    color TEXT,
    plates TEXT,
    body_number TEXT,
    chassis_number TEXT,
    engine_number TEXT,
    theft_date TIMESTAMP NOT NULL,
    insert_date TIMESTAMP NOT NULL,
    state TEXT NOT NULL
);

CREATE INDEX id_idx             ON wanted_vehicles(id);
CREATE INDEX body_number_idx    ON wanted_vehicles(body_number);
CREATE INDEX chassis_number_idx ON wanted_vehicles(chassis_number);
CREATE INDEX engine_number_idx  ON wanted_vehicles(engine_number);
