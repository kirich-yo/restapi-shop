BEGIN;


ALTER TABLE IF EXISTS public.cart DROP CONSTRAINT IF EXISTS cart_item_id_fkey;

ALTER TABLE IF EXISTS public.cart DROP CONSTRAINT IF EXISTS cart_user_id_fkey;

ALTER TABLE IF EXISTS public.item_category DROP CONSTRAINT IF EXISTS category_fkey;

ALTER TABLE IF EXISTS public.item_category DROP CONSTRAINT IF EXISTS item_fkey;

ALTER TABLE IF EXISTS public."order" DROP CONSTRAINT IF EXISTS order_item_id_fkey;

ALTER TABLE IF EXISTS public."order" DROP CONSTRAINT IF EXISTS order_user_id_fkey;

ALTER TABLE IF EXISTS public.review DROP CONSTRAINT IF EXISTS review_item_id_fkey;

ALTER TABLE IF EXISTS public.review DROP CONSTRAINT IF EXISTS review_user_id_fkey;

ALTER TABLE IF EXISTS public."user" DROP CONSTRAINT IF EXISTS user_role_id_fkey;

ALTER TABLE IF EXISTS public.warehouse_item DROP CONSTRAINT IF EXISTS warehouse_item_item_id_fk;

ALTER TABLE IF EXISTS public.warehouse_item DROP CONSTRAINT IF EXISTS warehouse_item_warehouse_id_fk;



DROP TABLE IF EXISTS public.cart;

CREATE TABLE IF NOT EXISTS public.cart
(
    id serial NOT NULL,
    user_id integer NOT NULL,
    item_id integer NOT NULL,
    amount integer NOT NULL,
    CONSTRAINT cart_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.category;

CREATE TABLE IF NOT EXISTS public.category
(
    id serial NOT NULL,
    name character varying(32) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT category_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.item;

CREATE TABLE IF NOT EXISTS public.item
(
    id serial NOT NULL,
    name character varying(128) COLLATE pg_catalog."default" NOT NULL,
    price real,
    sale_price real,
    photo_url character varying(128) COLLATE pg_catalog."default",
    CONSTRAINT item_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.item_category;

CREATE TABLE IF NOT EXISTS public.item_category
(
    id serial NOT NULL,
    item_id integer NOT NULL,
    category_id integer NOT NULL,
    CONSTRAINT item_category_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public."order";

CREATE TABLE IF NOT EXISTS public."order"
(
    id serial NOT NULL,
    user_id integer NOT NULL,
    item_id integer NOT NULL,
    amount integer NOT NULL,
    ship_date date NOT NULL,
    status character varying(15) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT order_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.review;

CREATE TABLE IF NOT EXISTS public.review
(
    id serial NOT NULL,
    rating integer NOT NULL,
    advantages text COLLATE pg_catalog."default",
    disadvantages text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    user_id integer NOT NULL,
    item_id integer NOT NULL,
    CONSTRAINT review_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.role;

CREATE TABLE IF NOT EXISTS public.role
(
    id serial NOT NULL,
    name character varying(32) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT role_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public."user";

CREATE TABLE IF NOT EXISTS public."user"
(
    id serial NOT NULL,
    username character varying(32) COLLATE pg_catalog."default" NOT NULL,
    first_name character varying(32) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(32) COLLATE pg_catalog."default",
    date_of_birth date NOT NULL,
    photo_url character varying(128) COLLATE pg_catalog."default",
    role_id integer NOT NULL,
    password character(32) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.warehouse;

CREATE TABLE IF NOT EXISTS public.warehouse
(
    id serial NOT NULL,
    address character varying(128) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT warehouse_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS public.warehouse_item;

CREATE TABLE IF NOT EXISTS public.warehouse_item
(
    id serial NOT NULL,
    warehouse_id integer NOT NULL,
    item_id integer NOT NULL,
    amount integer NOT NULL,
    CONSTRAINT warehouse_item_pkey PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.cart
    ADD CONSTRAINT cart_item_id_fkey FOREIGN KEY (item_id)
    REFERENCES public.item (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.cart
    ADD CONSTRAINT cart_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES public."user" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.item_category
    ADD CONSTRAINT category_fkey FOREIGN KEY (category_id)
    REFERENCES public.category (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.item_category
    ADD CONSTRAINT item_fkey FOREIGN KEY (item_id)
    REFERENCES public.item (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public."order"
    ADD CONSTRAINT order_item_id_fkey FOREIGN KEY (item_id)
    REFERENCES public.item (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public."order"
    ADD CONSTRAINT order_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES public."user" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.review
    ADD CONSTRAINT review_item_id_fkey FOREIGN KEY (item_id)
    REFERENCES public.item (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.review
    ADD CONSTRAINT review_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES public."user" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public."user"
    ADD CONSTRAINT user_role_id_fkey FOREIGN KEY (role_id)
    REFERENCES public.role (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.warehouse_item
    ADD CONSTRAINT warehouse_item_item_id_fk FOREIGN KEY (item_id)
    REFERENCES public.item (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS public.warehouse_item
    ADD CONSTRAINT warehouse_item_warehouse_id_fk FOREIGN KEY (warehouse_id)
    REFERENCES public.warehouse (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

END;
