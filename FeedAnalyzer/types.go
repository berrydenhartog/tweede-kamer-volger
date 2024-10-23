package main

import (
	"encoding/xml"
)

type Element struct {
	XMLName xml.Name
	Content string `xml:",innerxml"` 
}

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Link    []struct {
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Updated string `xml:"updated"`
	Author  struct {
		Name  string `xml:"name"`
		URI   string `xml:"uri"`
		Email string `xml:"email"`
	} `xml:"author"`
	Rights string  `xml:"rights"`
	ID     string  `xml:"id"`
	Entry  []Entry `xml:"entry"`
    UnknownParts []Element `xml:",any"`
}

type Entry struct {
	Title  string `xml:"title"`
	ID     string `xml:"id"`
	Author struct {
		Name string `xml:"name"`
	} `xml:"author"`
	Updated  string `xml:"updated"`
	Category struct {
		Term string `xml:"term,attr"`
	} `xml:"category"`
	Link []struct {
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"` 
	Content Content `xml:"content"`
    UnknownParts []Element `xml:",any"`
}

type Content struct {
	Type    string `xml:"type,attr"`
    Activiteit *struct {
        Content string `xml:",innerxml"` 
    } `xml:"activiteit"`
    ActiviteitActor *struct {
        Content string `xml:",innerxml"` 
    } `xml:"activiteitActor"`
    Agendapunt *struct {
        Content string `xml:",innerxml"` 
    } `xml:"agendapunt"`
    Besluit *struct {
        Content string `xml:",innerxml"` 
    } `xml:"besluit"`
    Commissie *struct {
        Content string `xml:",innerxml"` 
    } `xml:"commissie"`
    CommissieContactinformatie *struct {
        Content string `xml:",innerxml"` 
    } `xml:"commissieContactinformatie"`
    CommissieZetel *struct {
        Content string `xml:",innerxml"` 
    } `xml:"commissieZetel"`
    CommissieZetelVastPersoon *struct {
        Content string `xml:",innerxml"` 
    } `xml:"commissieZetelVastPersoon"`
    CommissieZetelVervangerPersoon *struct {
        Content string `xml:",innerxml"` 
    } `xml:"commissieZetelVervangerPersoon"`
    CommissieZetelVervangerVacature *struct {
        Content string `xml:",innerxml"` 
    } `xml:"commissieZetelVervangerVacature"`
    Document *struct {
        Content string `xml:",innerxml"` 
    } `xml:"document"`
    DocumentActor *struct {
        Content string `xml:",innerxml"` 
    } `xml:"documentActor"`
    DocumentVersie *struct {
        Content string `xml:",innerxml"` 
    } `xml:"documentVersie"`
    Fractie *struct {
        Content string `xml:",innerxml"` 
    } `xml:"fractie"`
    FractieZetel *struct {
        Content string `xml:",innerxml"` 
    } `xml:"fractieZetel"`
    FractieZetelPersoon *struct {
        Content string `xml:",innerxml"` 
    } `xml:"fractieZetelPersoon"`
    Kamerstukdossier *struct {
        Content string `xml:",innerxml"` 
    } `xml:"kamerstukdossier"`
    Persoon *struct {
        Content string `xml:",innerxml"` 
    } `xml:"persoon"`
    PersoonGeschenk *struct {
        Content string `xml:",innerxml"` 
    } `xml:"persoonGeschenk"`
    PersoonNevenfunctie *struct {
        Content string `xml:",innerxml"` 
    } `xml:"persoonNevenfunctie"`
    PersoonNevenfunctieInkomsten *struct {
        Content string `xml:",innerxml"` 
    } `xml:"persoonNevenfunctieInkomsten"`
    PersoonReis *struct {
        Content string `xml:",innerxml"` 
    } `xml:"persoonReis"`
    Reservering *struct {
        Content string `xml:",innerxml"` 
    } `xml:"reservering"`
    Stemming *struct {
        Content string `xml:",innerxml"` 
    } `xml:"stemming"`
    Toezegging *struct {
        Content string `xml:",innerxml"` 
    } `xml:"toezegging"`
    Vergadering *struct {
        Content string `xml:",innerxml"` 
    } `xml:"vergadering"`
	Verslag *struct {
		ID            string `xml:"id,attr"`
		Verwijderd    string `xml:"verwijderd,attr"`
		Bijgewerkt    string `xml:"bijgewerkt,attr"`
		ContentType   string `xml:"contentType,attr"`
		ContentLength string `xml:"contentLength,attr"`
		Vergadering   struct {
			XSIType string `xml:"xsi:type,attr"`
			Ref     string `xml:"ref,attr"`
		} `xml:"vergadering"`
		Soort  string `xml:"soort"`
		Status string `xml:"status"`
	} `xml:"verslag"`
    Zaak *struct {
        Content string `xml:",innerxml"` 
    } `xml:"zaak"`
    ZaakActor *struct {
        Content string `xml:",innerxml"` 
    } `xml:"zaakActor"`
    Zaal *struct {
        Content string `xml:",innerxml"` 
    } `xml:"zaal"`

	UnknownParts []Element `xml:",any"`
}