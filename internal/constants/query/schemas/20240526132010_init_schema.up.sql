-- enable uuid-ossp extension
CREATE EXTENSION "uuid-ossp";

CREATE TABLE users (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid() ,
                         "first_name" varchar NOT NULL,
                         "last_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT ((now())),
                         "updated_at"  timestamptz NOT NULL DEFAULT ((now()))
);

CREATE TABLE location (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid() ,
                         "country" varchar NOT NULL,
                         "region" varchar NOT NULL,
                         "zone" varchar NOT NULL,
                         "city" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT ((now())),
                         "updated_at"  timestamptz NOT NULL DEFAULT ((now()))
);

CREATE TABLE price_estimation (
                            "id" UUID PRIMARY KEY DEFAULT gen_random_uuid() ,
                            "product_name" varchar NOT NULL,
                            "user_id" uuid NOT NULL,
                            "price" decimal,
                            "location_id" UUID NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT ((now())),
                            "updated_at"  timestamptz NOT NULL DEFAULT ((now())),
                            CONSTRAINT fk_location_price_est FOREIGN KEY (location_id) REFERENCES location(id) ON DELETE CASCADE,
                            CONSTRAINT fk_users_price_est FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);