package serve

import (
	"log"
	"net/http"
	"os"

	"github.com/fajarihsan21/go-backend/src/routers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var InitServe = &cobra.Command{
	Use:   "serve",
	Short: "start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "127.0.0.1:8080"

		port := os.Getenv("PORT")

		if port != "" {
			addrs = ":" + port
		}

		log.Println("App running")

		c := cors.AllowAll()

		err := http.ListenAndServe(addrs, c.Handler(mainRoute))

		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}

	// 	if pr := os.Getenv("PORT"); pr != "" {
	// 		addrs = ":" + pr
	// 	}

	// 	log.Println("App running on " + addrs)

	// 	if err := http.ListenAndServe(addrs, mainRoute); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// } else {
	// 	return err
	// }
}
