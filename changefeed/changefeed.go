package changefeed

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

const (
	watchQuery = "SELECT * FROM %s.%s where COSMOS_CHANGEFEED_START_TIME() = '%s'"
)

type ChangeFeed struct {
	table    string
	keyspace string
	session  *gocql.Session
}

func Create(table string, keyspace string, session *gocql.Session) ChangeFeed {
	return ChangeFeed{
		table:    table,
		keyspace: keyspace,
		session:  session,
	}
}

func (ins *ChangeFeed) WatchChangeFeed() error {
	startTime := time.Now().UTC()
	pageState := []byte{}
	for {
		log.Println("[ChangeFeed] Starting query")
		// Build our query
		query := ins.session.Query(fmt.Sprintf(watchQuery, ins.keyspace, ins.table, startTime.Format(time.RFC3339)))

		// If we are continuing from where we were, set the page state
		if len(pageState) > 0 {
			query = query.PageState(pageState)
		}

		// Exec the query
		iter := query.Iter()

		// Store the page state for next time
		pageState = iter.PageState()

		// Read all of the changes and log them
		for {
			// An empty interface map is used to show all rows
			//  It is important that each row that we call MapScan() on is new
			row := map[string]interface{}{}
			if !iter.MapScan(row) {
				break
			}
			log.Printf("[ChangeFeed] Change detected: %v", row)
		}

		err := iter.Close()
		if err != nil {
			return err
		}
	}
}
