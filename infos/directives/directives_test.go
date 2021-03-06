package directives

import (
	"context"
	"log"
	"testing"

	"github.com/huangjunwen/sqlw-mysql/datasrc"

	"github.com/huangjunwen/tstsvc/mysql"
)

var (
	loader *datasrc.Loader
)

func exec(query string, args ...interface{}) {
	_, err := loader.Conn().ExecContext(context.Background(), query, args...)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	log.Printf("Starting testing MySQL server.\n")
	res, err := tstmysql.Run(nil)
	if err != nil {
		log.Panic(err)
	}
	defer res.Close()
	log.Printf("Testing MySQL server up.\n")

	loader, err = datasrc.NewLoader(res.DSN())
	if err != nil {
		log.Panic(err)
	}
	defer loader.Close()

	m.Run()

}
