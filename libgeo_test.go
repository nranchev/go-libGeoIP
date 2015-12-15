package libgeo

import (
	"io/ioutil"
	"testing"
)

const (
	GEO_IP_FILE_LOC = "GeoLiteCity.dat"
)

func TestNoFileExist(t *testing.T) {
	if _, err := Load("does/not/exist"); err == nil {
		t.Fatalf("expected err when not such file exists")
	}
}

func TestFileExists(t *testing.T) {
	if len(GEO_IP_FILE_LOC) == 0 {
		t.Fatalf("please fill out the path for GEO_IP_FILE_LOC")
	}
	if _, err := Load(GEO_IP_FILE_LOC); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestReadFromMem(t *testing.T) {
	if len(GEO_IP_FILE_LOC) == 0 {
		t.Fatalf("please fill out the path for GEO_IP_FILE_LOC")
	}
	if data, err := ioutil.ReadFile(GEO_IP_FILE_LOC); err != nil {
		t.Fatalf(err.Error())
	} else {
		if _, err := LoadData(data); err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func TestBasicFromFile(t *testing.T) {
	if geoip, err := Load(GEO_IP_FILE_LOC); err != nil {
		t.Fatalf(err.Error())
	} else {
		ip := "72.229.28.185"
		loc := geoip.GetLocationByIP(ip)
		if loc.City != "New York" {
			t.Fatalf("expected city to equal New York but got %s instead", loc.City)
		}
		if loc.CountryCode != "US" {
			t.Fatalf("expected country code to equal US but got %s instead", loc.CountryCode)
		}
		if loc.CountryName != "United States" {
			t.Fatalf("expected country name to equal United States but got %s instead", loc.CountryName)
		}
	}
}

func TestBasicFromMem(t *testing.T) {
	if data, err := ioutil.ReadFile(GEO_IP_FILE_LOC); err != nil {
		t.Fatalf(err.Error())
	} else {
		if geoip, err := LoadData(data); err != nil {
			t.Fatalf(err.Error())
		} else {
			ip := "72.229.28.185"
			loc := geoip.GetLocationByIP(ip)
			if loc.City != "New York" {
				t.Fatalf("expected city to equal New York but got %s instead", loc.City)
			}
			if loc.CountryCode != "US" {
				t.Fatalf("expected country code to equal US but got %s instead", loc.CountryCode)
			}
			if loc.CountryName != "United States" {
				t.Fatalf("expected country name to equal United States but got %s instead", loc.CountryName)
			}
		}
	}
}
