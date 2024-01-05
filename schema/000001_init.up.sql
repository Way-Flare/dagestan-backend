CREATE TABLE place 
(
    id          primary key, 
    title       varchar(255), 
    descrip     varchar(255),
    typeId      int references placeType (id) on delete cascade not null,
    latitude    decimal,
    longitude   decimal 
);

CREATE TABLE placeType 
(
    id          primary key, 
    typeName    varchar(255)
);

CREATE TABLE user 
(
    id              primary key,
    username        varchar(255) not null unique,
    email           varchar(255) not null,
    password_hash   varchar(255) not null
);

CREATE TABLE userPlace
(
    id      primary key,
    userId  int references user (id) on delete cascade   not null,
    placeId int references place (id) on delete cascade not null
);

CREATE TABLE image 
(
    id       primary key,
    imageUrl varchar(255)
);

CREATE TABLE placeImage 
(
    id      primary key,
    placeId int references place (id) on delete cascade not null,
    imageId int references image (id) on delete cascade not null 
);
