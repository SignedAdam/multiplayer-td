package unit_types

import (
	"errors"

	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai/npc"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai/tower"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	npcunit "github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit/npc"
	towerunit "github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit/tower"
)

type UnitType struct {
	UnitClass
	UnitTypeID int
	MaxHealth  int
	Damage     int
	AI         ai.AI
	CreateFunc func(npcTypeID, x, y, health, damage int, AI ai.AI) unit.Unit
}

type UnitClass int

var (
	UnitClassTower     UnitClass = 0
	UnitClassNPC       UnitClass = 1
	UnitClassStructure UnitClass = 2

	unitTypes map[int]UnitType = map[int]UnitType{
		//towers
		0: {
			UnitClass:  UnitClassTower,
			UnitTypeID: 0,
			MaxHealth:  300,
			Damage:     34,
			AI:         tower.NewTowerAI(),
			CreateFunc: towerunit.NewTower,
		},
		//npcs
		100: {
			UnitClass:  UnitClassNPC,
			UnitTypeID: 100,
			MaxHealth:  100,
			Damage:     10,
			AI:         npc.NewNPCAI(),
			CreateFunc: npcunit.NewNPC,
		},
		101: {
			UnitClass:  UnitClassNPC,
			UnitTypeID: 101,
			MaxHealth:  150,
			Damage:     10,
			AI:         npc.NewNPCAI(),
			CreateFunc: npcunit.NewNPC,
		},
		//structures
		200: {
			UnitClass:  UnitClassStructure,
			UnitTypeID: 200,
			MaxHealth:  300,
			Damage:     0,
			//AI:
			//CreateFunc:
		},
	}
)

func GetUnitType(unitTypeID int) (*UnitType, error) {
	unitType, ok := unitTypes[unitTypeID]
	if !ok {
		return nil, errors.New("unit type ID doesn't exist")
	}

	return &unitType, nil

}

func (unitType *UnitType) CreateUnit(x, y int) (unit.Unit, error) {
	return unitType.CreateFunc(unitType.UnitTypeID, x, y, unitType.MaxHealth, unitType.Damage, unitType.AI), nil
}
