CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table "user" (
                        id uuid PRIMARY KEY default uuid_generate_v4(),
                        name varchar(500),
                        username varchar(500),
                        email varchar(500),
                        phonenumber varchar(500),
                        password varchar(500),
                        createdAt timestamptz,
                        updatedAt timestamptz,
                        deletedAt timestamptz default  null
);
