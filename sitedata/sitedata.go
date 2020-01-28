package sitedata

type MenuEntry struct {
	Title string
	Dest  string
}

type Page struct {
	Address string
	Title   string
	Menu []MenuEntry
}
