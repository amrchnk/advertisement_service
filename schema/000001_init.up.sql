CREATE TABLE advert (
                        id serial primary key,
                        title varchar(200) not null,
                        description varchar(1000),
                        price int not null,
                        created timestamp not null default now()
);

CREATE TABLE photos (
                        id serial primary key,
                        link varchar(200) not null,
                        first bool default false,
                        advert_id int,
                        FOREIGN KEY (advert_id) references advert (id) ON DELETE CASCADE
);