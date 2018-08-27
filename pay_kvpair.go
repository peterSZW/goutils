package goutils

//package main

import (
	"net/url"
	"sort"
	"strings"
)

type kvpair struct {
	k, v string
}

type kvpairs []kvpair

func (t kvpairs) Less(i, j int) bool {
	return t[i].k < t[j].k
}

func (t kvpairs) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t kvpairs) Len() int {
	return len(t)
}

func (t kvpairs) Sort() {
	sort.Sort(t)
}

func (t kvpairs) UrlEncode() (t2 kvpairs) {
	for _, kv := range t {

		kv.v = url.QueryEscape(kv.v)
		t2 = append(t2, kv)

	}
	return
}

func (t kvpairs) RemoveEmpty() (t2 kvpairs) {
	for _, kv := range t {
		if kv.v != "" {
			t2 = append(t2, kv)
		}
	}
	return
}

func (t kvpairs) Join() string {
	var strs []string
	for _, kv := range t {
		strs = append(strs, kv.k+"="+kv.v)
	}
	return strings.Join(strs, "&")
}
