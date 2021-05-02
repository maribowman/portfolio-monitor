create table transactions
(
    transaction_id  serial       not null,
    isin            varchar(100),
    ticker          varchar(100) not null,
    title           varchar(200) not null,
    actual_price    double       not null,
    total_price     double       not null,
    effective_price double       not null,
    date            date         not null,
    comment         text
);

create table watchlist
(
    watchlist_id       serial       not null,
    isin               varchar(100),
    ticker             varchar(100) not null,
    title              varchar(200) not null,
    entry_price        double       not null,
    notification_price double       not null
);