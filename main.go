package main

import (
	kvc "key-value-client/client"
	"log"
	"math/rand"
	"sync"
)

const MAX_CLIENTS = 50
const MAX_ITERS = 100
const VALUES_LEN = 50
const ADDR = "192.168.1.36:6379"

var VALUES = [VALUES_LEN]string{
"oHueSG7IH8",
"hDpqFkb691",
"XOdkubnhWU",
"J9M8MEwIw7",
"RYLbh0DleP",
"vyi3XanhI3",
"N1OccDG91v",
"lbzEK3CYBh",
"EZZkaONy7C",
"ItIEJrKkas",
"4uopGFX7s0",
"qntsuBmyHZ",
"idOQrJSTVB",
"K8cu9sbDUn",
"Yn1FC2K6J7",
"YQZInpCbyo",
"tfoH9xNcFg",
"h2X0neFPJJ",
"rkaHMGaJlD",
"bM3kJqX7TX",
"d87AJ9kiuX",
"c6SRY5l0il",
"LKFEkfF4x5",
"4ltyJ2FOmB",
"XwzV2Wdhya",
"VeoFjyDPe4",
"sa2Rpm1GxL",
"UERIurh9su",
"usPMa8Rsca",
"enmYzUq7wn",
"Gp0NqZTyuw",
"ZNQgiQUWUf",
"iCfsVroTLM",
"Y06gJnrYex",
"FckJSiCdjh",
"sY1RCnMvRX",
"1uOVdYxcdq",
"lpA2vZ8tkf",
"XrZjK5DPJu",
"WKG1SvNdam",
"QEdzhaxG07",
"8fkW89LOUz",
"hrOPafwhlc",
"s3GKur6MCA",
"XScFHEcDK7",
"KKEnEQOlMd",
"0XrMbyRk6C",
"MKCrPp91SZ",
"h82GVTZ5UP",
"7f0RVvFwcC",
}

func main() {
	var wg sync.WaitGroup
	var setters [MAX_CLIENTS]*kvc.KeyValueClient
	var getters [MAX_CLIENTS]*kvc.KeyValueClient

	for i := range setters {
    newClient := kvc.New(ADDR)
		setters[i] = &newClient
		err := setters[i].Connect()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Setters", i, "Connected")
	}

	for i := range getters {
    newClient := kvc.New(ADDR)
		getters[i] = &newClient
		err := getters[i].Connect()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Getters", i, "Connected")
	}

	wg.Add(1)
	go emit(setters, MAX_ITERS, &wg)
	wg.Add(1)
	go consume(getters, MAX_ITERS, &wg)

	wg.Wait()
}

func emit(setters [MAX_CLIENTS]*kvc.KeyValueClient, iters int, wg *sync.WaitGroup) {
	defer wg.Done()
	iter := 0
	for {
		for i := range setters {
			randKeyIdx := rand.Intn(VALUES_LEN)
			randValIdx := rand.Intn(VALUES_LEN)

			_, _ = setters[i].Set(VALUES[randKeyIdx], VALUES[randValIdx])
		}
		iter++

		if iter == iters {
			return
		}
	}
}

func consume(getters [MAX_CLIENTS]*kvc.KeyValueClient, iters int, wg *sync.WaitGroup) {
	defer wg.Done()
	iter := 0
	for {
		for i := range getters {
			randKeyIdx := rand.Intn(VALUES_LEN)

      _, _ = getters[i].Get(VALUES[randKeyIdx])
		}
		iter++

		if iter == iters {
			return
		}
	}
}
