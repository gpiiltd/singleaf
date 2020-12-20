package main

import (
	"singleaf/driver"
	"singleaf/middlewares"
	
	"singleaf/user/handler"
	"singleaf/user/repo"
	"singleaf/user/usecase"

	appHandler "singleaf/apps/handler"
	appRepo "singleaf/apps/repo"
	appUsecase "singleaf/apps/usecase"

	ustbeHandler "singleaf/ustbe/handler"
	ustbeRepo "singleaf/ustbe/repo"
	ustbeUsecase "singleaf/ustbe/usecase"

	bestaHandler "singleaf/besta/handler"
	bestaRepo "singleaf/besta/repo"
	bestaUsecase "singleaf/besta/usecase"

	enterprisesHandler "singleaf/enterprises/handler"
	enterprisesRepo "singleaf/enterprises/repo"
	enterprisesUsecase "singleaf/enterprises/usecase"

	// armsHandler "singleaf/arms/handler"
	// armsRepo "singleaf/arms/repo"
	// armsUsecase "singleaf/arms/usecase"

	"flag"
	"log"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ServeStatic(router *mux.Router) {
	//var dir string

	//flag.StringVar(&dir, "dir", viper.GetString("file.path"), "the directory to serve files from. Defaults to the current dir")
	//flag.Parse()
	staticPaths := map[string]string{
		"images": viper.GetString("file.path"),
		"ustbephoto": viper.GetString("ustbe.path"),
		"publisherphoto":           viper.GetString("publisher.path"),
		"bestaphoto":          viper.GetString("besta.path"),
		"appsphoto":          viper.GetString("apps.path"),
	}
	for pathName, pathValue := range staticPaths {
		flag.StringVar(&pathValue, pathValue, pathValue, "the directory to serve files from. Defaults to the current dir")
		flag.Parse()
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
	}
}

func main() {
	db := driver.Config()

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.Logging)
	router.Use(middlewares.SetMiddlewareJSON)


	ServeStatic(router)
		

	//router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(dir))))

	userRepo := repo.CreateRepo(db)
	userUsecase := usecase.CreateUsecase(userRepo)
	handler.CreateHandler(router, userUsecase)

	appRepo_ := appRepo.CreateRepo(db)
	appUsecase := appUsecase.CreateUsecase(appRepo_)
	appHandler.CreateHandler(router, appUsecase)

	ustbeRepo_ := ustbeRepo.CreateRepo(db)
	ustbeUsecase := ustbeUsecase.CreateUsecase(ustbeRepo_)
	ustbeHandler.CreateHandler(router, ustbeUsecase)

	bestaRepo_ := bestaRepo.CreateRepo(db)
	bestaUsecase := bestaUsecase.CreateUsecase(bestaRepo_)
	bestaHandler.CreateHandler(router, bestaUsecase)

	enterprisesRepo_ := enterprisesRepo.CreateRepo(db)
	enterprisesUsecase := enterprisesUsecase.CreateUsecase(enterprisesRepo_)
	enterprisesHandler.CreateHandler(router, enterprisesUsecase)

	// armsRepo_ := armsRepo.CreateRepo(db)
	// armsUsecase := armsUsecase.CreateUsecase(armsRepo_)
	// armsHandler.CreateHandler(router, armsUsecase)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Unified Access Framework - GPII"))
	})

	router.HandleFunc("/getip", handleRequest)

	router.HandleFunc("/getname", handleNameRequest)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})

	log.Println("Server starts")
	//logrus.Fatal(http.ListenAndServe(viper.GetString("server.port"), "/etc/letsencrypt/live/concerto.my-gpi.io/fullchain.pem", "/etc/letsencrypt/live/concerto.my-gpi.io/privkey.pem" ,handlers.CORS(headersOk, originsOk, methodsOk)(router)))

	// ssl := map[string]string{
 //        	"cert": "/etc/letsencrypt/live/concerto.my-gpi.io/fullchain.pem",
 //        	"key":  "/etc/letsencrypt/live/concerto.my-gpi.io/privkey.pem",
	// }


	errs := make(chan error)

    // Starting HTTP server
    go func() {
        log.Printf("Staring HTTP service on %s ...", viper.GetString("server.openport"))

        //logrus.Fatal(http.ListenAndServe(viper.GetString("server.port"), handlers.CORS(headersOk, originsOk, methodsOk)(router)))
        if err := http.ListenAndServe(viper.GetString("server.openport"), handlers.CORS(headersOk, originsOk, methodsOk)(router)); err != nil {
        //if err := logrus.Fatal(http.ListenAndServe(viper.GetString("server.port"), handlers.CORS(headersOk, originsOk, methodsOk)(router))) {
            errs <- err
        }
        
    }()

    // Starting HTTPS server
    go func() {
        log.Printf("Staring HTTPS service on %s ...", viper.GetString("server.closedport"))

        //logrus.Fatal(http.ListenAndServe(viper.GetString("server.port"), "/etc/letsencrypt/live/concerto.my-gpi.io/fullchain.pem", "/etc/letsencrypt/live/concerto.my-gpi.io/privkey.pem" ,handlers.CORS(headersOk, originsOk, methodsOk)(router)))
        
        // if err := http.ListenAndServeTLS(viper.GetString("server.closedport"), ssl["cert"], ssl["key"], handlers.CORS(headersOk, originsOk, methodsOk)(router)); err != nil {
        // 	errs <- err
        // }
    }()


	//errs := Run(viper.GetString("server.openport"), viper.GetString("server.closeport"), map[string]string{
    //    "cert": "/etc/letsencrypt/live/concerto.my-gpi.io/fullchain.pem",
    //    "key":  "/etc/letsencrypt/live/concerto.my-gpi.io/privkey.pem",
    //})


    // This will run forever until channel receives error
    select {
    case err := <-errs:
    	log.Fatalln(err)
    	//logrus.Fatal()
        //logrus.Fatal("Could not start serving service due to (error: %s)" + err.Error())
    }

}

// func getHost(w *http.ResponseWriter, r *http.Request) string {
//     if r.URL.IsAbs() {
//         host := r.Host
//         // Slice off any port information.
//         if i := strings.Index(host, ":"); i != -1 {
//             host = host[:i]
//         }
//         return host
//     }
//     return r.URL.Host
// }

func getIP(r *http.Request) (string, error) {
    // //Get IP from the X-REAL-IP header
    // ip := r.Header.Get("X-REAL-IP")
    // netIP := net.ParseIP(ip)
    // if netIP != nil {
    //     return ip, nil
    // }

    // //Get IP from X-FORWARDED-FOR header
    // ips := r.Header.Get("X-FORWARDED-FOR")
    // splitIps := strings.Split(ips, ",")
    // for _, ip := range splitIps {
    //     netIP := net.ParseIP(ip)
    //     if netIP != nil {
    //         return ip, nil
    //     }
    // }

    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return "", err
    }
    netIP := net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }
    return "", fmt.Errorf("No valid ip found")
}

func handleNameRequest(w http.ResponseWriter, r *http.Request) {
	var host = getHost(r)
    w.WriteHeader(200)
    w.Write([]byte(host))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    ip, err := getIP(r)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("No valid ip"))
    }
    w.WriteHeader(200)
    w.Write([]byte(ip))
}

func getHost(r *http.Request) string {
	fmt.Println(r.URL.Host)
	fmt.Println("----------------")
	fmt.Println(r.Host)
    if r.URL.IsAbs() {
        host := r.Host
        // Slice off any port information.
        if i := strings.Index(host, ":"); i != -1 {
            host = host[:i]
        }
        return host
    }
    return r.Host
}