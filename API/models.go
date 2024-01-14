package main

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/jett65/cakes_by_kindraV2/internal/datebase"
)

// Cake
type Cake struct {
	ID           uuid.UUID `json:"id"`
	Type         string    `json:"type"`
	Layer_number int       `json:"layer_number"`
	Tiere_number int       `json:"tiere_number"`
	Size         string    `json:"size"`
	Price        float64   `json:"price"`
}

func databaseCakeToCake(dbCake datebase.Cake) (Cake, error) {
    floatPrice, err := strconv.ParseFloat(dbCake.Price, 64)
    if err != nil {
        return Cake{}, err 
    }

    
    return Cake {
        ID: dbCake.ID,
        Type: dbCake.Type,
        Layer_number: int(dbCake.LayerNumber),
        Tiere_number: int(dbCake.TiereNumber),
        Size: dbCake.Size,
        Price: floatPrice,
    }, nil

}

func databaseCakesToCakes(dbCake []datebase.Cake) ([]Cake, error) {
    cakes := []Cake{}
    for _, cake := range dbCake {
        conCake, err := databaseCakeToCake(cake)
        if err != nil {
            return cakes, err
        }
        
        cakes = append(cakes, conCake)
    }

    return cakes, nil
}

// Flavor
type Flavor struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    AddPrice float64 `json:"add_price"`
}

func databaseFlavorToFlavor(dbFlavor datebase.Flavor) (Flavor, error) {
    floatAddPrice, err := strconv.ParseFloat(dbFlavor.AddPrice, 64)
    if err != nil {
        return Flavor{}, err
    }
    
    return Flavor {
        ID: dbFlavor.ID,
        Name: dbFlavor.Name,
        AddPrice: floatAddPrice,
    }, nil
}

func databaseFlavorsToFlavors(dbFlavors []datebase.Flavor) ([]Flavor, error) {
    flavors := []Flavor{}
    for _, flavor := range dbFlavors {
        conFlavor, err := databaseFlavorToFlavor(flavor)
        if err != nil {
            return flavors, err
        }
        
        flavors = append(flavors, conFlavor)
    }

    return flavors, nil
}

// Frosting
type Frosting struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    AddPrice float64 `json:"add_price"`
}

func databaseFrostingToFrosting(dbFlavor datebase.Frosting) (Frosting, error) {
    floatAddPrice, err := strconv.ParseFloat(dbFlavor.AddPrice, 64)
    if err != nil {
        return Frosting{}, err
    }
    
    return Frosting {
        ID: dbFlavor.ID,
        Name: dbFlavor.Name,
        AddPrice: floatAddPrice,
    }, nil
}

func databaseFrostingsToFrostngs(dbFrosting []datebase.Frosting) ([]Frosting, error) {
    frostings := []Frosting{}
    for _, frosting := range dbFrosting {
        conFrosting, err := databaseFrostingToFrosting(frosting)
        if err != nil {
            return frostings, err
        }
        
        frostings = append(frostings, conFrosting)
    }

    return frostings, nil
}

// Filling
type Filliong struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
}

func databaseFillingToFilling(dbFlavor datebase.Filling) (Filliong, error) {
    floatPrice, err := strconv.ParseFloat(dbFlavor.Price, 64)
    if err != nil {
        return Filliong{}, err
    }
    
    return Filliong {
        ID: dbFlavor.ID,
        Name: dbFlavor.Name,
        Price: floatPrice,
    }, nil
}

func databaseFillingsToFilligs(dbFrosting []datebase.Filling) ([]Filliong, error) {
    fillings := []Filliong{}
    for _, filling := range dbFrosting {
        conFrosting, err := databaseFillingToFilling(filling)
        if err != nil {
            return fillings, err
        }
        
        fillings = append(fillings, conFrosting)
    }

    return fillings, nil
}
