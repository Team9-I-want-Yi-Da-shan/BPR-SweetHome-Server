--
-- PostgreSQL database dump
--

-- Dumped from database version 10.20
-- Dumped by pg_dump version 10.20

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
-- Name: created_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.created_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
	NEW.created_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.created_at_column() OWNER TO postgres;

--
-- Name: update_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.update_at_column() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: activity; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.activity (
    id integer NOT NULL,
    family_id integer,
    name character varying,
    description character varying,
    start_at integer,
    finish_at integer,
    reminder integer,
    person_id integer,
    isfinish integer DEFAULT 0,
    isrepeat integer DEFAULT 0,
    isfamily integer,
    isalarm integer
);


ALTER TABLE public.activity OWNER TO postgres;

--
-- Name: activity_family_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.activity_family_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.activity_family_id_seq OWNER TO postgres;

--
-- Name: activity_family_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.activity_family_id_seq OWNED BY public.activity.id;


--
-- Name: activity_participant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.activity_participant (
    user_id integer,
    activity_id integer
);


ALTER TABLE public.activity_participant OWNER TO postgres;

--
-- Name: article; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.article (
    id integer NOT NULL,
    user_id integer,
    title character varying,
    content text,
    updated_at integer,
    created_at integer
);


ALTER TABLE public.article OWNER TO postgres;

--
-- Name: article_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.article_id_seq OWNER TO postgres;

--
-- Name: article_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.article_id_seq OWNED BY public.article.id;


--
-- Name: bill; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bill (
    id integer NOT NULL,
    name character varying,
    family_id integer,
    amount integer,
    deadline integer
);


ALTER TABLE public.bill OWNER TO postgres;

--
-- Name: bill_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bill_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bill_id_seq OWNER TO postgres;

--
-- Name: bill_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bill_id_seq OWNED BY public.bill.id;


--
-- Name: family; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.family (
    id integer NOT NULL,
    name character varying,
    admin_id integer
);


ALTER TABLE public.family OWNER TO postgres;

--
-- Name: family_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.family_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.family_id_seq OWNER TO postgres;

--
-- Name: family_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.family_id_seq OWNED BY public.family.id;


--
-- Name: plan_family; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.plan_family (
    id integer NOT NULL,
    family_id integer,
    name character varying NOT NULL,
    description text,
    comment text,
    updated_at integer,
    created_at integer
);


ALTER TABLE public.plan_family OWNER TO postgres;

--
-- Name: plan_family_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.plan_family_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.plan_family_id_seq OWNER TO postgres;

--
-- Name: plan_family_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.plan_family_id_seq OWNED BY public.plan_family.id;


--
-- Name: plan_person; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.plan_person (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying,
    description text,
    comment text,
    updated_at integer,
    created_at integer
);


ALTER TABLE public.plan_person OWNER TO postgres;

--
-- Name: plan_person_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.plan_person_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.plan_person_id_seq OWNER TO postgres;

--
-- Name: plan_person_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.plan_person_id_seq OWNED BY public.plan_person.id;


--
-- Name: repeat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.repeat (
    id integer NOT NULL,
    activity_id integer,
    repeat_interval integer,
    repeat_start integer
);


ALTER TABLE public.repeat OWNER TO postgres;

--
-- Name: repeat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.repeat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.repeat_id_seq OWNER TO postgres;

--
-- Name: repeat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.repeat_id_seq OWNED BY public.repeat.id;


--
-- Name: response; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.response (
    id integer NOT NULL,
    admin_id integer NOT NULL,
    request_id integer NOT NULL,
    confirm integer NOT NULL,
    finish integer NOT NULL,
    created_at integer,
    updated_at integer
);


ALTER TABLE public.response OWNER TO postgres;

--
-- Name: response_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.response_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.response_id_seq OWNER TO postgres;

--
-- Name: response_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.response_id_seq OWNED BY public.response.id;


--
-- Name: smart_device; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.smart_device (
    id integer NOT NULL,
    family_id integer,
    device_id integer,
    device_name character varying,
    status integer,
    ssid integer,
    url integer
);


ALTER TABLE public.smart_device OWNER TO postgres;

--
-- Name: smart_device_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.smart_device_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.smart_device_id_seq OWNER TO postgres;

--
-- Name: smart_device_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.smart_device_id_seq OWNED BY public.smart_device.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    email character varying,
    password character varying,
    name character varying,
    updated_at integer,
    created_at integer,
    familyid integer DEFAULT 0
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: activity id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity ALTER COLUMN id SET DEFAULT nextval('public.activity_family_id_seq'::regclass);


--
-- Name: article id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article ALTER COLUMN id SET DEFAULT nextval('public.article_id_seq'::regclass);


--
-- Name: bill id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bill ALTER COLUMN id SET DEFAULT nextval('public.bill_id_seq'::regclass);


--
-- Name: family id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.family ALTER COLUMN id SET DEFAULT nextval('public.family_id_seq'::regclass);


--
-- Name: plan_family id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_family ALTER COLUMN id SET DEFAULT nextval('public.plan_family_id_seq'::regclass);


--
-- Name: plan_person id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_person ALTER COLUMN id SET DEFAULT nextval('public.plan_person_id_seq'::regclass);


--
-- Name: repeat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.repeat ALTER COLUMN id SET DEFAULT nextval('public.repeat_id_seq'::regclass);


--
-- Name: response id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.response ALTER COLUMN id SET DEFAULT nextval('public.response_id_seq'::regclass);


--
-- Name: smart_device id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.smart_device ALTER COLUMN id SET DEFAULT nextval('public.smart_device_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: activity; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.activity (id, family_id, name, description, start_at, finish_at, reminder, person_id, isfinish, isrepeat, isfamily, isalarm) FROM stdin;
1	1	1FamilyActivityTest	\N	\N	\N	\N	\N	\N	\N	\N	\N
3	\N	FamilyRepeatActivityTest	No desc	\N	\N	\N	17	\N	\N	\N	\N
4	1	name	\N	\N	\N	\N	\N	\N	\N	1	\N
5	1	name2	\N	\N	\N	\N	\N	\N	\N	1	\N
6	\N	activityName	Some des	0	0	0	17	0	0	0	\N
7	\N	activityName	Some des	0	0	0	17	0	0	0	\N
13	1	activityName Family1	Some des	0	0	0	\N	0	0	1	\N
14	1	activityName Family2	Some des	0	0	0	\N	0	0	1	\N
15	1	activityName Family2 with time	Some des	1653481087	0	23333	\N	0	0	1	\N
16	1	activityName Family3	Some des	1653481087	1653481227	23333	\N	0	0	1	\N
20	\N	activityName Family3	Some des	1653481087	1653481227	3600	17	0	1	0	\N
22	\N	activity Person 5	Some des	1653481087	1653481227	3600	17	0	1	0	\N
23	\N	activity Person 6	Some des	1653481087	1653481227	3600	17	0	1	0	\N
19	\N	activityName Family4	Some des2	1653481027	1653481027	23133	17	0	0	0	\N
35	\N	1family act		0	0	0	17	0	0	0	\N
37	\N	1family act		0	0	10	17	0	1	0	\N
38	\N	1family act		0	0	10	17	0	1	0	\N
39	\N	1family act		0	0	10	17	0	1	0	\N
40	1	1family act		0	0	10	\N	0	1	1	\N
53	1	5family act		0	0	10	\N	0	0	1	\N
54	1	name2	\N	\N	\N	\N	\N	\N	\N	1	\N
55	1	6family act		0	0	10	\N	0	0	1	\N
56	1	7family act		0	0	10	\N	0	0	1	\N
57	1	7family act		0	0	10	\N	0	0	1	\N
58	1	8family act		0	0	10	\N	0	0	1	\N
61	1	9family act		0	0	10	\N	0	0	1	\N
62	1	10family act		0	0	10	\N	0	0	1	\N
63	1	11family act T		0	0	10	\N	0	0	1	\N
64	1	Family act 5/28	5/28	1653682140	1653688140	0	\N	0	1	1	\N
75	1	2Family act 5/28	\N	1653689140	\N	0	\N	0	0	1	\N
21	\N	do some thing bob	new des	0	0	0	17	1	0	0	0
79	\N	213213	32132	1653972540	0	-1	17	0	0	0	0
80	\N	xiu jie kou	ddddd	1654030920	0	-1	17	0	0	0	0
81	\N	xie bao gao	xiexiexiexie	1654094040	1654058100	-1	17	0	0	0	0
\.


--
-- Data for Name: activity_participant; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.activity_participant (user_id, activity_id) FROM stdin;
17	53
32	53
31	53
17	55
32	55
31	55
17	56
32	56
31	56
17	57
32	57
31	57
17	58
32	58
31	58
17	64
30	64
\.


--
-- Data for Name: article; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.article (id, user_id, title, content, updated_at, created_at) FROM stdin;
5	17	articleTitle	articleContent	1653179480	1653179480
6	17	啊、	才	1653179499	1653179499
\.


--
-- Data for Name: bill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bill (id, name, family_id, amount, deadline) FROM stdin;
\.


--
-- Data for Name: family; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.family (id, name, admin_id) FROM stdin;
1	TestHome	17
4	Test Home	32
5	Test	30
0	defaultFamily	0
6	STQ' family	29
\.


--
-- Data for Name: plan_family; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.plan_family (id, family_id, name, description, comment, updated_at, created_at) FROM stdin;
1	1	FamilyPlanName	FamilyPlanDes	FamilyCom	1653339500	1653339500
2	1		des	comment 1	1653860427	1653860427
3	1		des	comment 1	1653860495	1653860495
4	1		xxx	xxxx	1654034505	1654034505
5	1		xxx	xxxx	1654036532	1654036532
6	1		xxx	xxxx	1654038263	1654038263
7	1		xxxx	xxxx	1654038429	1654038429
8	1		gfd	hujj	1654038472	1654038472
\.


--
-- Data for Name: plan_person; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.plan_person (id, user_id, name, description, comment, updated_at, created_at) FROM stdin;
3	17	testPlan2	des2	\N	1652990304	1652988993
5	17	33name	des1	com	1653124825	1653124825
21	17	do some thing bob	new des	update new comment	1653421380	1653239132
22	17	personal plan1			1653510475	1653510475
23	17	personal plan1	des	comment 1	1653510587	1653510587
24	17	2do some thing bob			1653591677	1653591677
25	17		androidx.lifecycle.MutableLiveData@3d7b7f2	androidx.lifecycle.MutableLiveData@4113e43	1653781040	1653781040
26	17		mPPDescription	mPPComment	1653830832	1653830832
27	17		mPPDescription	mPPComment	1653831085	1653831085
28	17		mPPDescription	mPPComment	1653831255	1653831255
29	17		mPPDescription	mPPComment	1653831675	1653831675
30	17		asdasd	fgasds	1654009383	1654009383
31	17		des	comment 1	1654009707	1654009707
32	17		sss	rrrrr	1654010818	1654010818
33	17		xxxsx	xxxax	1654023047	1654023047
34	17		xxx	xxx	1654035033	1654035033
35	17		xxx	xxx	1654035040	1654035040
36	17		xxx	xxx	1654035095	1654035095
\.


--
-- Data for Name: repeat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.repeat (id, activity_id, repeat_interval, repeat_start) FROM stdin;
3	1	3600	166456435
4	4	3600	12421412
5	5	46300	1653481057
13	37	360000	0
14	38	360000	0
15	39	360000	0
16	40	360000	0
17	64	172800	1653682140
\.


--
-- Data for Name: response; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.response (id, admin_id, request_id, confirm, finish, created_at, updated_at) FROM stdin;
6	17	31	1	1	1653597527	1653601977
5	17	32	1	1	1653597527	1653603676
\.


--
-- Data for Name: smart_device; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.smart_device (id, family_id, device_id, device_name, status, ssid, url) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, email, password, name, updated_at, created_at, familyid) FROM stdin;
16	testEmail@via.dk	$2a$10$uEsXIS/KfhvFEQBe6x1r7Om1Rpj6TNRb9c9bq73.6IbE8p7TL5NYO	haocheng	1652956177	1652956177	\N
17	123@123.dk	$2a$10$FyHonH0PneZXAJrtBX77u.NYB9Wy1FEX/Y6Cb.mBrcpQnicNlI.US	haocheng	1653315936	1652956820	1
20	123@via.dk	$2a$10$lBc/.eUe6E5fhza4AnoPLOkz5rBE8n/aslxZSZPQ1D6HBVY9wSxNG	haochengZ	1653316295	1653316295	0
23	email@via.dk	$2a$10$PRaC7OkWj99B8c4CUJGlUeUrFHWm2wPR5A2dZSxVnVjerUwBfiBdm	username	1653417549	1653417549	0
24	test002@via.dk	$2a$10$a.dTecTbOJRo5faY5ezgge7ErIfvKYk7x73FjOC8ILjO0el5sEade	test	1653418962	1653418962	0
25	HBW@via.dk	$2a$10$XtuRUQAaIHJLUxfUk0rVdOx9X0ZbVn30kyNsTy.7vDxpO0OsFnJL.	HBWWWW	1653419299	1653419299	0
26	hbwww@via.dk	$2a$10$clXmpMYwi651QJOAQzcb6.BKxVXuVuvk1PG.MR/S/vfObl0NEwETm	hbwwww	1653421060	1653421060	0
27	bowen@via.dk	$2a$10$r4KFFee6gg/GzvzpxOxf.uNJf0kbV6ftCSsXHEFsDmTFntInup8zy	hbwwwwwwwwww	1653421272	1653421272	0
28	bowenxxx@via.dk	$2a$10$7XJLDJtxi8OF/wthAG2J7e1IxxjjLjGJ6wkg88WwAMk7WE21Boks2	bowenxxx	1653484255	1653484255	0
29	stq@via.dk	$2a$10$WyVqZXP4vi970evlYxbSmOIGm5vmUT61./loVjiYSM./Y0UnnEiiG	STQSTQ	1653484403	1653484403	0
30	haocheng@via.dk	$2a$10$aWxPRXp3zGo7WZwtBymzH.w5szwRT9imLDYYprR0doMvJug9BzXue	haochengzhangaaa	1653488440	1653488440	0
31	asd@via.dk	$2a$10$DkW0AdK5xW4pf/fMfqJJju9sagM8qwC2vaWKRrUm1h5whbHS81Mv2	asdafag	1653488766	1653488766	0
33	285511@VIIII.dk	$2a$10$iMRv1mv7LgXHXp.OjpqBTOVRQV3M32RKOIxmmmAC3LDJK4wEnOy4O	asdasd	1653513284	1653513284	0
34	hbwtest@via.dk	$2a$10$KwMGSes5hmWNTElr/4khJuzLAdHfed45XCaNQPMp.Q1GOrRArfhp.	hbwtest	1653567040	1653567040	0
0	DefaultUser	\N	\N	1653588008	1653588008	0
32	123123@via.dk	$2a$10$X/nGox3VDxQWp6stiEcD2eeiWaXdeDIyGG.RF3i3saLza8TBDwnTq	xxxxa\ntestininte	1653603676	1653489979	1
35	xxxx@123.dk	$2a$10$bTucutwpBtBTlcPLRaB.7eOJYs1056jXnQ5Q0/W.jgJhoEiyU9W7C	testgg	1653781526	1653781526	0
\.


--
-- Name: activity_family_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.activity_family_id_seq', 81, true);


--
-- Name: article_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.article_id_seq', 6, true);


--
-- Name: bill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bill_id_seq', 1, false);


--
-- Name: family_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.family_id_seq', 6, true);


--
-- Name: plan_family_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.plan_family_id_seq', 8, true);


--
-- Name: plan_person_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.plan_person_id_seq', 36, true);


--
-- Name: repeat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.repeat_id_seq', 17, true);


--
-- Name: response_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.response_id_seq', 6, true);


--
-- Name: smart_device_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.smart_device_id_seq', 1, false);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 37, true);


--
-- Name: activity activity_family_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_family_pk PRIMARY KEY (id);


--
-- Name: article article_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_id PRIMARY KEY (id);


--
-- Name: family family_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.family
    ADD CONSTRAINT family_pk PRIMARY KEY (id);


--
-- Name: plan_family plan_family_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_family
    ADD CONSTRAINT plan_family_pk PRIMARY KEY (id);


--
-- Name: plan_person plan_person_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_person
    ADD CONSTRAINT plan_person_pk PRIMARY KEY (id);


--
-- Name: repeat repeat_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.repeat
    ADD CONSTRAINT repeat_pk PRIMARY KEY (id);


--
-- Name: response response_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.response
    ADD CONSTRAINT response_pk PRIMARY KEY (id);


--
-- Name: smart_device smart_device_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.smart_device
    ADD CONSTRAINT smart_device_pk PRIMARY KEY (id);


--
-- Name: user user_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_id PRIMARY KEY (id);


--
-- Name: activity_family_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX activity_family_id_uindex ON public.activity USING btree (id);


--
-- Name: bill_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX bill_id_uindex ON public.bill USING btree (id);


--
-- Name: family_admin_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX family_admin_id_uindex ON public.family USING btree (admin_id);


--
-- Name: family_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX family_id_uindex ON public.family USING btree (id);


--
-- Name: plan_family_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX plan_family_id_uindex ON public.plan_family USING btree (id);


--
-- Name: repeat_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX repeat_id_uindex ON public.repeat USING btree (id);


--
-- Name: response_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX response_id_uindex ON public.response USING btree (id);


--
-- Name: smart_device_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX smart_device_id_uindex ON public.smart_device USING btree (id);


--
-- Name: user_email_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_email_uindex ON public."user" USING btree (email);


--
-- Name: article create_article_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_article_created_at BEFORE INSERT ON public.article FOR EACH ROW EXECUTE PROCEDURE public.created_at_column();


--
-- Name: plan_family create_plan_person_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_plan_person_created_at BEFORE INSERT ON public.plan_family FOR EACH ROW EXECUTE PROCEDURE public.created_at_column();


--
-- Name: plan_person create_plan_person_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_plan_person_created_at BEFORE INSERT ON public.plan_person FOR EACH ROW EXECUTE PROCEDURE public.created_at_column();


--
-- Name: response create_plan_person_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_plan_person_created_at BEFORE INSERT ON public.response FOR EACH ROW EXECUTE PROCEDURE public.created_at_column();


--
-- Name: user create_user_created_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER create_user_created_at BEFORE INSERT ON public."user" FOR EACH ROW EXECUTE PROCEDURE public.created_at_column();


--
-- Name: article update_article_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_article_updated_at BEFORE UPDATE ON public.article FOR EACH ROW EXECUTE PROCEDURE public.update_at_column();


--
-- Name: plan_family update_plan_person_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_plan_person_updated_at BEFORE UPDATE ON public.plan_family FOR EACH ROW EXECUTE PROCEDURE public.update_at_column();


--
-- Name: plan_person update_plan_person_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_plan_person_updated_at BEFORE UPDATE ON public.plan_person FOR EACH ROW EXECUTE PROCEDURE public.update_at_column();


--
-- Name: response update_plan_person_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_plan_person_updated_at BEFORE UPDATE ON public.response FOR EACH ROW EXECUTE PROCEDURE public.update_at_column();


--
-- Name: user update_user_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON public."user" FOR EACH ROW EXECUTE PROCEDURE public.update_at_column();


--
-- Name: activity activity_family_family_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_family_family_id_fk FOREIGN KEY (family_id) REFERENCES public.family(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: activity_participant activity_participant_activity_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity_participant
    ADD CONSTRAINT activity_participant_activity_fk FOREIGN KEY (activity_id) REFERENCES public.activity(id) ON DELETE CASCADE;


--
-- Name: activity_participant activity_participant_user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity_participant
    ADD CONSTRAINT activity_participant_user_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: activity activity_person_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_person_fk FOREIGN KEY (person_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: article article_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: bill bill_family_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bill
    ADD CONSTRAINT bill_family_fk FOREIGN KEY (family_id) REFERENCES public.family(id);


--
-- Name: family family_admin_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.family
    ADD CONSTRAINT family_admin_fk FOREIGN KEY (admin_id) REFERENCES public."user"(id);


--
-- Name: plan_family family_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_family
    ADD CONSTRAINT family_fk FOREIGN KEY (family_id) REFERENCES public.family(id) ON DELETE SET NULL;


--
-- Name: user family_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT family_id_fk FOREIGN KEY (familyid) REFERENCES public.family(id);


--
-- Name: repeat repeat_activity_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.repeat
    ADD CONSTRAINT repeat_activity_fk FOREIGN KEY (activity_id) REFERENCES public.activity(id) ON DELETE CASCADE;


--
-- Name: response response_admin_user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.response
    ADD CONSTRAINT response_admin_user_fk FOREIGN KEY (admin_id) REFERENCES public."user"(id);


--
-- Name: response response_request_user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.response
    ADD CONSTRAINT response_request_user_fk FOREIGN KEY (request_id) REFERENCES public."user"(id);


--
-- Name: smart_device smart_device_family_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.smart_device
    ADD CONSTRAINT smart_device_family_fk FOREIGN KEY (family_id) REFERENCES public.family(id) ON DELETE CASCADE;


--
-- Name: plan_person user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_person
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public."user"(id) DEFERRABLE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

