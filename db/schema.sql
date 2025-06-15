SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: chat_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.chat_type AS ENUM (
    'ChatAll',
    'ChatTeam',
    'ChatSquad',
    'ChatAdmin'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: bans; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.bans (
    ban_id bigint NOT NULL,
    user_id integer NOT NULL,
    player_id integer NOT NULL,
    reason character varying(255) NOT NULL,
    note text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    end_at timestamp with time zone
);


--
-- Name: bans_ban_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.bans_ban_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: bans_ban_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.bans_ban_id_seq OWNED BY public.bans.ban_id;


--
-- Name: logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.logs (
    log_id bigint NOT NULL,
    log_type character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    server_id integer,
    user_id integer,
    player_id integer,
    victim_id integer,
    attacker_id integer,
    player_ip character varying,
    squad_id character varying,
    squad_name text,
    team_id character varying,
    chat_type public.chat_type,
    is_teamkill boolean,
    message text,
    map text
);


--
-- Name: logs_log_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.logs_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: logs_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.logs_log_id_seq OWNED BY public.logs.log_id;


--
-- Name: players; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.players (
    player_id integer NOT NULL,
    name character varying(64) NOT NULL,
    eos_id character varying(32) NOT NULL,
    steam_id character varying(17) NOT NULL,
    first_seen timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_seen timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: players_player_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.players_player_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: players_player_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.players_player_id_seq OWNED BY public.players.player_id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying NOT NULL
);


--
-- Name: servers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.servers (
    server_id integer NOT NULL,
    "user" text NOT NULL,
    host text NOT NULL,
    password text NOT NULL,
    rcon_password text NOT NULL,
    rcon_port text NOT NULL,
    admin_log_file_path text NOT NULL,
    game_log_file_path text NOT NULL
);


--
-- Name: servers_server_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.servers_server_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: servers_server_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.servers_server_id_seq OWNED BY public.servers.server_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    steam_id character varying(17) NOT NULL,
    name character varying(255) NOT NULL,
    avatar text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: bans ban_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bans ALTER COLUMN ban_id SET DEFAULT nextval('public.bans_ban_id_seq'::regclass);


--
-- Name: logs log_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs ALTER COLUMN log_id SET DEFAULT nextval('public.logs_log_id_seq'::regclass);


--
-- Name: players player_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players ALTER COLUMN player_id SET DEFAULT nextval('public.players_player_id_seq'::regclass);


--
-- Name: servers server_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.servers ALTER COLUMN server_id SET DEFAULT nextval('public.servers_server_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Name: bans bans_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bans
    ADD CONSTRAINT bans_pkey PRIMARY KEY (ban_id);


--
-- Name: logs logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (log_id);


--
-- Name: players players_eos_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_eos_id_key UNIQUE (eos_id);


--
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (player_id);


--
-- Name: players players_steam_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_steam_id_key UNIQUE (steam_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: servers servers_game_log_file_path_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.servers
    ADD CONSTRAINT servers_game_log_file_path_key UNIQUE (game_log_file_path);


--
-- Name: servers servers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.servers
    ADD CONSTRAINT servers_pkey PRIMARY KEY (server_id);


--
-- Name: servers servers_rcon_port_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.servers
    ADD CONSTRAINT servers_rcon_port_key UNIQUE (rcon_port);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: users users_steam_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_steam_id_key UNIQUE (steam_id);


--
-- Name: bans bans_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bans
    ADD CONSTRAINT bans_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(player_id);


--
-- Name: bans bans_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bans
    ADD CONSTRAINT bans_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: logs logs_attacker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_attacker_id_fkey FOREIGN KEY (attacker_id) REFERENCES public.players(player_id);


--
-- Name: logs logs_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(player_id);


--
-- Name: logs logs_server_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_server_id_fkey FOREIGN KEY (server_id) REFERENCES public.servers(server_id);


--
-- Name: logs logs_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: logs logs_victim_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_victim_id_fkey FOREIGN KEY (victim_id) REFERENCES public.players(player_id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20250519001544'),
    ('20250525160158'),
    ('20250525162512'),
    ('20250525163527'),
    ('20250525163550');
