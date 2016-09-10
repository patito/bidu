CREATE TABLE public.site(
    site_id serial NOT NULL,
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    physical_address character varying(50) NOT NULL,
    comments text,
    CONSTRAINT site_id PRIMARY KEY (site_id),
    CONSTRAINT site_name UNIQUE (name),
    CONSTRAINT site_slug UNIQUE (slug)

);
ALTER TABLE public.site OWNER TO postgres;
