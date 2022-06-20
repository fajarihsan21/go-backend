package serve

import (
	"log"
	"net/http"
	"os"

	"github.com/fajarihsan21/go-backend/src/routers"
	"github.com/spf13/cobra"
)

var InitServe = &cobra.Command{
	Use:   "serve",
	Short: "start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = ":8080"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = "127.0.0.1:" + pr
		}

		log.Println("App running on " + addrs)

		if err := http.ListenAndServe(addrs, mainRoute); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}
