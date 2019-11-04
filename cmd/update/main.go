package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/opencars/wanted/pkg/config"
	"github.com/opencars/wanted/pkg/storage"
	"github.com/opencars/wanted/pkg/storage/postgres"
	"github.com/opencars/wanted/pkg/wanted"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	newPath := flag.Arg(0)

	// Check file extension.
	if filepath.Ext(newPath) != ".json" {
		fmt.Fprintln(os.Stderr, "invalid file extension")
		os.Exit(1)
	}

	// Get configuration.
	conf, err := config.New(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Register postgres adapter.
	db, err := postgres.New(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	store := storage.New(db)

	// Load all the vehicles from DB.
	transport, err := store.WantedVehicles()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	sort.Sort(transport)
	fmt.Printf("Amount of vehicle in database: %d\n", len(transport))

	f, err := os.Open(newPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dec := json.NewDecoder(f)

	// Read the array open bracket.
	if _, err := dec.Token(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	newTransport := make([]wanted.Vehicle, 0, 100)
	for dec.More() {
		var tmp wanted.Vehicle
		err := dec.Decode(&tmp)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if res := transport.Search(tmp.ID); res == -1 {
			newTransport = append(newTransport, tmp)
		}
	}

	// TODO: Find transport that was removed.

	// New stolen cars.
	fmt.Println("Total:", len(newTransport))

	// TODO: Parse wanted.Transport into storage.Transport.
	// TODO: Parse and insert.
	// fmt.Println(v.ID)
	tmps := make([]storage.WantedVehicle, len(newTransport))
	for i, v := range newTransport {
		tmps[i].ID = v.ID
		tmps[i].OVD = v.OVD
		tmps[i].Color = v.Color
		tmps[i].TheftDate = v.TheftData
		tmps[i].State = "stolen"
		tmps[i].InsertDate = v.InsertDate
		brand, model, kind := v.ParseBrand()

		// TODO: Check "-" , "---".

		if brand != "" {
			tmps[i].Brand = &brand
		}

		if model != "" {
			tmps[i].Model = &model
		}

		if kind != "" {
			tmps[i].Kind = &kind
		}

		if v.BodyNumber != "" {
			tmps[i].BodyNumber = &v.BodyNumber
		}

		if v.ChassisNumber != "" {
			tmps[i].ChassisNumber = &v.ChassisNumber
		}

		if v.EngineNumber != "" {
			tmps[i].EngineNumber = &v.EngineNumber
		}

		if v.VehicleNumber != "" {
			tmps[i].Plates = &v.VehicleNumber
		}

		fmt.Println(tmps[i].Brand)
		if err := store.InsertWantedVehicle(&tmps[i]); err != nil {
			fmt.Println(err)
		}
	}
}
