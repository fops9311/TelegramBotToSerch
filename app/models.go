// models.go

package main

type LocationResult struct {
	Type          string        `json:"type"`
	RegionNames   RegionNames   `json:"regionNames"`
	EssID         EssID         `json:"essId"`
	Coordinates   Coordinates   `json:"coordinates"`
	HierarchyInfo HierarchyInfo `json:"hierarchyInfo"`
}

type RegionNames struct {
	FullName             string `json:"fullName"`
	ShortName            string `json:"shortName"`
	DisplayName          string `json:"displayName"`
	PrimaryDisplayName   string `json:"primaryDisplayName"`
	SecondaryDisplayName string `json:"secondaryDisplayName"`
	LastSearchName       string `json:"lastSearchName"`
}

type EssID struct {
	SourceName string `json:"sourceName"`
	SourceID   string `json:"sourceId"`
}

type Coordinates struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type HierarchyInfo struct {
	Country Country `json:"country"`
	Airport Airport `json:"airport"`
}

type Country struct {
	Name     string `json:"name"`
	IsoCode2 string `json:"isoCode2"`
	IsoCode3 string `json:"isoCode3"`
}

type Airport struct {
	AirportCode string `json:"airportCode"`
	AirportID   string `json:"airportId"`
	Metrocode   string `json:"metrocode"`
	Multicity   string `json:"multicity"`
}

type LocationResponse struct {
	Query string           `json:"q"`
	RID   string           `json:"rid"`
	RC    string           `json:"rc"`
	SR    []LocationResult `json:"sr"`
}
