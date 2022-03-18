package types

type Item struct {
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Price    float32  `json:"price"`
	Author   string   `json:"author"`
	Director string   `json:"director"`
	Artist   string   `json:"artist"`
	Tracks   []Track  `json:"tracks"`
	Chapters []string `json:"chapters"`
}

type Track struct {
	Seconds int    `json:"seconds"`
	Name    string `json:"name"`
}

/*
item_type varchar(10) not null,
title varchar(100) not null,
author varchar(100) not null,
director varchar(100) not null,
year int not null,
minutes decimal not null,
tracks jsonb,
chapters jsonb*/
