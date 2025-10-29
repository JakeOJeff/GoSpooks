package main

type Room struct {
    Name    string
    Desc    string
    Exits   map[string]string
    Locked  bool
    Key     string
    Items   []string
}