--
-- PostgreSQL database dump
--

-- Dumped from database version 11.9 (Debian 11.9-0+deb10u1)
-- Dumped by pg_dump version 11.9 (Debian 11.9-0+deb10u1)

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

--
-- Name: creators; Type: TYPE; Schema: public; Owner: postgresql_2020102619yz7vctgkgqy6qre
--

CREATE TYPE public.creators AS (
	creator1 text,
	creator2 text
);


ALTER TYPE public.creators OWNER TO postgresql_2020102619yz7vctgkgqy6qre;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: prog_languages; Type: TABLE; Schema: public; Owner: postgresql_2020102619yz7vctgkgqy6qre
--

CREATE TABLE public.prog_languages (
    name text,
    year integer,
    wikipedia text,
    imperative boolean,
    object_oriented boolean,
    functional boolean,
    procedural boolean,
    generic boolean,
    reflective boolean,
    creators public.creators
);


ALTER TABLE public.prog_languages OWNER TO postgresql_2020102619yz7vctgkgqy6qre;

--
-- Data for Name: prog_languages; Type: TABLE DATA; Schema: public; Owner: postgresql_2020102619yz7vctgkgqy6qre
--

COPY public.prog_languages (name, year, wikipedia, imperative, object_oriented, functional, procedural, generic, reflective, creators) FROM stdin;
C	1972	C_(programming_language)	t	f	f	t	t	f	("Dennis Ritchie","Bell Labs")
Assembly	1949	Assembly_language	t	f	f	f	f	f	("Maurice Vincent Wilkes","David John Wheeler")
Bash	1989	Bash_(Unix_shell)	t	f	f	t	f	f	("Brian Fox","Richard Stallman")
C++	1985	C%2B%2B	t	t	t	t	t	f	("Bjarne Stroustrup","Douglas McIlroy")
C#	2000	C_Sharp_(programming_language)	t	t	t	t	t	t	("Anders Hejlsberg","Microsoft")
COBOL	1959	COBOL	t	t	f	t	f	f	("Howard Bromberg","Norman Discount")
Java	1995	Java_(programming_language)	t	t	t	t	t	t	("James Gosling","Oracle Corporation")
Fortran	1957	Fortran	t	t	t	t	t	t	("John Backus","IBM")
BASIC	1964	BASIC	t	f	f	t	f	f	("John G. Kemeny","Thomas E. Kurtz")
K	1993	K_(programming_language)	f	f	f	f	f	f	("Arthur Whitney","Kx Systems")
\.


--
-- PostgreSQL database dump complete
--

