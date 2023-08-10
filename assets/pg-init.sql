CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" text,
  "price" bigint
);

CREATE TABLE "cart_items" (
  "product_id" bigserial PRIMARY KEY,
  "quantity" bigint,
  FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" text,
  "password" text,
  "role" text
);

INSERT INTO "products" ("name", "price") VALUES ('pet rock', 30);
INSERT INTO "products" ("name", "price") VALUES ('iPhone 9', 800);

INSERT INTO "cart_items" ("product_id", "quantity") VALUES (1, 3);
INSERT INTO "cart_items" ("product_id", "quantity") VALUES (2, 1);

INSERT INTO "accounts" ("username", "password", "role") VALUES ('root', 'password', 'admin');
INSERT INTO "accounts" ("username", "password", "role") VALUES ('guest', 'password', 'user');
