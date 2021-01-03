--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0 (Debian 13.0-1.pgdg100+1)
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-03 21:17:34 CET

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 201 (class 1259 OID 16387)
-- Name: creators; Type: TABLE; Schema: public; Owner: marko
--

CREATE TABLE public.creators (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text
);


ALTER TABLE public.creators OWNER TO marko;

--
-- TOC entry 200 (class 1259 OID 16385)
-- Name: creators_id_seq; Type: SEQUENCE; Schema: public; Owner: marko
--

CREATE SEQUENCE public.creators_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.creators_id_seq OWNER TO marko;

--
-- TOC entry 2967 (class 0 OID 0)
-- Dependencies: 200
-- Name: creators_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: marko
--

ALTER SEQUENCE public.creators_id_seq OWNED BY public.creators.id;


--
-- TOC entry 204 (class 1259 OID 16424)
-- Name: language_creators; Type: TABLE; Schema: public; Owner: marko
--

CREATE TABLE public.language_creators (
    language_id bigint NOT NULL,
    creator_id bigint NOT NULL
);


ALTER TABLE public.language_creators OWNER TO marko;

--
-- TOC entry 203 (class 1259 OID 16399)
-- Name: languages; Type: TABLE; Schema: public; Owner: marko
--

CREATE TABLE public.languages (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    year integer,
    wikipedia text,
    imperative boolean,
    object_oriented boolean,
    functional boolean,
    procedural boolean,
    generic boolean,
    reflective boolean
);


ALTER TABLE public.languages OWNER TO marko;

--
-- TOC entry 202 (class 1259 OID 16397)
-- Name: languages_id_seq; Type: SEQUENCE; Schema: public; Owner: marko
--

CREATE SEQUENCE public.languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.languages_id_seq OWNER TO marko;

--
-- TOC entry 2968 (class 0 OID 0)
-- Dependencies: 202
-- Name: languages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: marko
--

ALTER SEQUENCE public.languages_id_seq OWNED BY public.languages.id;


--
-- TOC entry 2815 (class 2604 OID 16390)
-- Name: creators id; Type: DEFAULT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.creators ALTER COLUMN id SET DEFAULT nextval('public.creators_id_seq'::regclass);


--
-- TOC entry 2816 (class 2604 OID 16402)
-- Name: languages id; Type: DEFAULT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.languages ALTER COLUMN id SET DEFAULT nextval('public.languages_id_seq'::regclass);


--
-- TOC entry 2958 (class 0 OID 16387)
-- Dependencies: 201
-- Data for Name: creators; Type: TABLE DATA; Schema: public; Owner: marko
--

COPY public.creators (id, created_at, updated_at, deleted_at, name) FROM stdin;
17	2021-01-03 20:10:46.411514+00	2021-01-03 20:10:46.411514+00	\N	Dennis Ritchie
18	2021-01-03 20:10:46.411514+00	2021-01-03 20:10:46.411514+00	\N	Bell Labs
21	2021-01-03 20:14:09.375091+00	2021-01-03 20:14:09.375091+00	\N	Maurice Vincent Wilkes
22	2021-01-03 20:14:09.375091+00	2021-01-03 20:14:09.375091+00	\N	David John Wheeler
23	2021-01-03 20:14:19.167386+00	2021-01-03 20:14:19.167386+00	\N	Brian Fox
24	2021-01-03 20:14:19.167386+00	2021-01-03 20:14:19.167386+00	\N	Richard Stallman
25	2021-01-03 20:14:28.304419+00	2021-01-03 20:14:28.304419+00	\N	Bjarne Stroustrup
26	2021-01-03 20:14:28.304419+00	2021-01-03 20:14:28.304419+00	\N	Douglas McIlroy
27	2021-01-03 20:14:38.115994+00	2021-01-03 20:14:38.115994+00	\N	Anders Hejlsberg
28	2021-01-03 20:14:38.115994+00	2021-01-03 20:14:38.115994+00	\N	Microsoft
29	2021-01-03 20:14:48.390129+00	2021-01-03 20:14:48.390129+00	\N	Howard Bromberg
30	2021-01-03 20:14:48.390129+00	2021-01-03 20:14:48.390129+00	\N	Norman Discount
31	2021-01-03 20:14:58.129168+00	2021-01-03 20:14:58.129168+00	\N	James Gosling
32	2021-01-03 20:14:58.129168+00	2021-01-03 20:14:58.129168+00	\N	Oracle Corporation
33	2021-01-03 20:15:11.38069+00	2021-01-03 20:15:11.38069+00	\N	Oracle Corporation
34	2021-01-03 20:15:11.38069+00	2021-01-03 20:15:11.38069+00	\N	IBM
35	2021-01-03 20:15:17.814086+00	2021-01-03 20:15:17.814086+00	\N	IBM
36	2021-01-03 20:15:17.814086+00	2021-01-03 20:15:17.814086+00	\N	Thomas E. Kurtz
37	2021-01-03 20:15:27.810274+00	2021-01-03 20:15:27.810274+00	\N	Arthur Whitney
38	2021-01-03 20:15:27.810274+00	2021-01-03 20:15:27.810274+00	\N	Kx Systems
\.


--
-- TOC entry 2961 (class 0 OID 16424)
-- Dependencies: 204
-- Data for Name: language_creators; Type: TABLE DATA; Schema: public; Owner: marko
--

COPY public.language_creators (language_id, creator_id) FROM stdin;
6	17
6	18
8	21
8	22
9	23
9	24
10	25
10	26
11	27
11	28
12	29
12	30
13	31
13	32
14	33
14	34
15	35
15	36
16	37
16	38
\.


--
-- TOC entry 2960 (class 0 OID 16399)
-- Dependencies: 203
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: marko
--

COPY public.languages (id, created_at, updated_at, deleted_at, name, year, wikipedia, imperative, object_oriented, functional, procedural, generic, reflective) FROM stdin;
6	2021-01-03 20:10:46.365136+00	2021-01-03 20:10:46.365136+00	\N	C	1972	C_(programming_language)	t	f	f	t	t	f
8	2021-01-03 20:14:09.374845+00	2021-01-03 20:14:09.374845+00	\N	Assembly	1949	Assembly_language	t	f	f	f	f	f
9	2021-01-03 20:14:19.167159+00	2021-01-03 20:14:19.167159+00	\N	Bash	1989	Bash_(Unix_shell)	t	f	f	t	f	f
10	2021-01-03 20:14:28.304224+00	2021-01-03 20:14:28.304224+00	\N	C++	1985	C%2B%2B	t	f	t	t	t	f
11	2021-01-03 20:14:38.115787+00	2021-01-03 20:14:38.115787+00	\N	C#	2000	C_Sharp_(programming_language)	t	f	t	t	t	t
12	2021-01-03 20:14:48.389906+00	2021-01-03 20:14:48.389906+00	\N	COBOL	1959	COBOL	t	f	f	t	f	f
13	2021-01-03 20:14:58.128993+00	2021-01-03 20:14:58.128993+00	\N	Java	1995	Java_(programming_language)	t	f	t	t	t	t
14	2021-01-03 20:15:11.380484+00	2021-01-03 20:15:11.380484+00	\N	Fortran	1957	Fortran	t	f	t	t	t	t
15	2021-01-03 20:15:17.813844+00	2021-01-03 20:15:17.813844+00	\N	BASIC	1964	BASIC	t	f	f	t	f	f
16	2021-01-03 20:15:27.810063+00	2021-01-03 20:15:27.810063+00	\N	K	1993	K_(programming_language)	f	f	f	f	f	f
\.


--
-- TOC entry 2969 (class 0 OID 0)
-- Dependencies: 200
-- Name: creators_id_seq; Type: SEQUENCE SET; Schema: public; Owner: marko
--

SELECT pg_catalog.setval('public.creators_id_seq', 38, true);


--
-- TOC entry 2970 (class 0 OID 0)
-- Dependencies: 202
-- Name: languages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: marko
--

SELECT pg_catalog.setval('public.languages_id_seq', 16, true);


--
-- TOC entry 2818 (class 2606 OID 16395)
-- Name: creators creators_pkey; Type: CONSTRAINT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.creators
    ADD CONSTRAINT creators_pkey PRIMARY KEY (id);


--
-- TOC entry 2824 (class 2606 OID 16428)
-- Name: language_creators language_creators_pkey; Type: CONSTRAINT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.language_creators
    ADD CONSTRAINT language_creators_pkey PRIMARY KEY (language_id, creator_id);


--
-- TOC entry 2822 (class 2606 OID 16407)
-- Name: languages languages_pkey; Type: CONSTRAINT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);


--
-- TOC entry 2819 (class 1259 OID 16396)
-- Name: idx_creators_deleted_at; Type: INDEX; Schema: public; Owner: marko
--

CREATE INDEX idx_creators_deleted_at ON public.creators USING btree (deleted_at);


--
-- TOC entry 2820 (class 1259 OID 16408)
-- Name: idx_languages_deleted_at; Type: INDEX; Schema: public; Owner: marko
--

CREATE INDEX idx_languages_deleted_at ON public.languages USING btree (deleted_at);


--
-- TOC entry 2826 (class 2606 OID 16434)
-- Name: language_creators fk_language_creators_creator; Type: FK CONSTRAINT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.language_creators
    ADD CONSTRAINT fk_language_creators_creator FOREIGN KEY (creator_id) REFERENCES public.creators(id);


--
-- TOC entry 2825 (class 2606 OID 16429)
-- Name: language_creators fk_language_creators_language; Type: FK CONSTRAINT; Schema: public; Owner: marko
--

ALTER TABLE ONLY public.language_creators
    ADD CONSTRAINT fk_language_creators_language FOREIGN KEY (language_id) REFERENCES public.languages(id);


-- Completed on 2021-01-03 21:17:38 CET

--
-- PostgreSQL database dump complete
--

