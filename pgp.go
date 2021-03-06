package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	//"regexp"
)

func pgpSearch(mailSet mapset.Set) mapset.Set {
	fmt.Println("==== PGP SEARCH ====")
	mailIterator := mailSet.Iterator()
	for mail := range mailIterator.C {
		pgpSet := mapset.NewSet()
		fmt.Println("[+] pgp search for " + mail.(string))
		domain := "http://pgp.mit.edu/pks/lookup?search=" + mail.(string)
		body := retriveRequestBody(domain)
		pgpSet = findMailInText(body, pgpSet)
		if pgpSet != nil {
			pgpIterator := pgpSet.Iterator()
			for email := range pgpIterator.C {
				fmt.Println(email)
			}
			mailSet = mailSet.Union(pgpSet)
		} else {
			fmt.Println("No result found")
		}
	}
	return mailSet
}
