package main

import (
	"flag"
	"fmt"

	"github.com/huskarl10/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var F float64
var C float64
var K float64
var out string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene

	flag.Float64Var(&F, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&C, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&K, "K", 0.0, "temperatur i kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "", "tempraturskala for resultat")

	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	var res float64

	if C != 0 {
		if out == "F" {
			res = conv.CelsiusToFahrenheit(C)
		}
	} else if out == "K" {
		res = conv.CelsiusToKelvin(C)

	} else if F != 0 {
		if out == "C" {
			res = conv.FahrenheitToCelsius(F)
		}
	} else if out == "K" {
		res = conv.FahrenheitToKelvin(F)

	} else if K != 0 {
		if out == "F" {
			res = conv.KelvinToFahrenheit(K)
		}
	} else if out == "C" {
		res = conv.KelvinToCelsius(K)

	}

	fmt.Printf("%.2f\n", res)

	fmt.Println(isFlagPassed("out"))

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
