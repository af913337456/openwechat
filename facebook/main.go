package main

//
//import (
//	"encoding/json"
//	"os"
//
//	"github.com/tamboto2000/facebook"
//)
//
//func main() {
//	// get new client
//	fb := facebook.New()
//
//	// set facebook login cookie
//	fb.SetCookieStr(`your_facebook_cookie`)
//
//	// initiate client
//	if err := fb.Init(); err != nil {
//		panic(err.Error())
//	}
//
//	username := "franklin.tamboto.3"
//
//	// get profile
//	profile, err := fb.Profile(username)
//	if err != nil {
//		panic(err.Error())
//	}
//	Errorprofile.About
//	// before getting all data from "About" section, make sure to call Profile.SyncAbout first
//	//if err := profile.About.SyncAbout(); err != nil {
//	//	panic(err.Error())
//	//}
//
//	if err := profile.About.SyncWorkAndEducation(); err != nil {
//		panic(err.Error())
//	}
//
//	if err := profile.About.SyncPlacesLived(); err != nil {
//		panic(err.Error())
//	}
//
//	if err := profile.About.SyncContactAndBasicInfo(); err != nil {
//		panic(err.Error())
//	}
//
//	if err := profile.About.SyncFamilyAndRelationships(); err != nil {
//		panic(err.Error())
//	}
//
//	if err := profile.About.SyncDetails(); err != nil {
//		panic(err.Error())
//	}
//
//	if err := profile.About.SyncLifeEvents(); err != nil {
//		panic(err.Error())
//	}
//
//	// save profile to a file
//	f, err := os.Create(username + ".json")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	defer f.Close()
//
//	if err := json.NewEncoder(f).Encode(profile); err != nil {
//		panic(err.Error())
//	}
//}
