CREATE TABLE users (
  id bigserial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,
  account_rub bigserial not null,
  account_usd bigserial not null,
  account_eur bigserial not null,
  account_balance_rub bigserial not null,
  account_balance_usd bigserial not null,
  account_balance_eur bigserial not null
);
