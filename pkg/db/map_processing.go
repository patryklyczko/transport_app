package db

import (
	"context"
	"io"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"github.com/qedus/osmpbf"
)

func (d *DBController) ProcessMap(path *structures.MapRequest) error {
	startTime := time.Now()
	collectionNodes := d.db.Collection("Nodes")
	collectionRelations := d.db.Collection("Relations")

	file, err := os.Open(path.Path)
	if err != nil {
		d.log.Debugf("Error opening path %v", err)
		return err
	}
	defer file.Close()

	// Creating a new decoder
	decoder := osmpbf.NewDecoder(file)
	err = decoder.Start(runtime.NumCPU())
	if err != nil {
		d.log.Debugf("Error creating decoder %v", err)
		return err
	}

	nodesPositionChannel := make(chan interface{}, 7000)
	nodesRelationsChannel := make(chan interface{}, 7000)
	nodesPositionsInsert := make([]interface{}, 0)
	nodesRelationsInsert := make([]interface{}, 0)

	// WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range nodesPositionChannel {
				nodesPositionsInsert = append(nodesPositionsInsert, item.(structures.NodePositions))

				if len(nodesPositionsInsert) >= 8000 {
					_, err = collectionNodes.InsertMany(context.Background(), nodesPositionsInsert)
					if err != nil {
						d.log.Debugf("error while inserting positions: %v", err)
					}
					nodesPositionsInsert = nodesPositionsInsert[:0]
				}
			}

			if len(nodesPositionsInsert) > 0 {
				_, err = collectionNodes.InsertMany(context.Background(), nodesPositionsInsert)
				if err != nil {
					d.log.Debugf("error while inserting positions")
				}
				nodesPositionsInsert = nodesPositionsInsert[:0]
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range nodesRelationsChannel {
				nodesRelationsInsert = append(nodesRelationsInsert, item.(structures.NodesRelations))

				if len(nodesRelationsInsert) >= 8000 {
					_, err = collectionRelations.InsertMany(context.Background(), nodesRelationsInsert)
					if err != nil {
						d.log.Debugf("error while inserting relations")
					}
					nodesRelationsInsert = nodesRelationsInsert[:0]
				}
			}

			if len(nodesRelationsInsert) > 0 {
				_, err = collectionRelations.InsertMany(context.Background(), nodesRelationsInsert)
				if err != nil {
					d.log.Debugf("error while inserting relations")
				}
				nodesRelationsInsert = nodesRelationsInsert[:0]
			}
		}()
	}

	for {
		if v, err := decoder.Decode(); err != nil {
			if err == io.EOF {
				break
			}
			d.log.Debugf("decoder rip %v", err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				// Node
				if containTag(v.Tags, "highway") {
					position := structures.NodePositions{
						Parent: v.ID,
						Position: structures.Position{
							Lat: float32(v.Lat),
							Lon: float32(v.Lon),
						},
					}
					nodesPositionChannel <- position
				}
			case *osmpbf.Way:
				// Way
				if containTag(v.Tags, "highway") {
					connections := structures.NodesRelations{
						Parent:   v.ID,
						Children: v.NodeIDs,
						MaxSpeed: v.Tags["maxspeed"],
					}
					nodesRelationsChannel <- connections
				}
			case *osmpbf.Relation:
				// Relation
			default:
			}
		}
	}

	close(nodesPositionChannel)
	close(nodesRelationsChannel)
	wg.Wait()
	endTime := time.Now()
	d.log.Debugf("Function took %v to execute\n", endTime.Sub(startTime))
	return nil
}

func containTag(m map[string](string), target string) bool {
	for k := range m {
		if k == target {
			return true
		}
	}
	return false
}
