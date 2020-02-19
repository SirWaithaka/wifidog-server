package http

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

//0 - AUTH_DENIED - User firewall users are deleted and the user removed.
//6 - AUTH_VALIDATION_FAILED - User email validation timeout has occured and user/firewall is deleted
//1 - AUTH_ALLOWED - User was valid, add firewall rules if not present
//5 - AUTH_VALIDATION - Permit user access to email to get validation email under default rules
//-1 - AUTH_ERROR - An error occurred during the validation process
type AuthTypes struct {
	AuthDenied int
	AuthValidationFailed int
	AuthAllowed int
	AuthValidation int
	AuthError int
}

var authTypes = AuthTypes{0,6,1,5,-1}

type WDLoginProtocol struct {
	GwAddress string `json:"gw_address"`
	GwID      string `json:"gw_id"`
	GwPort    string `json:"gw_port"`
	URL       string `json:"url"`
}

type LoginParams struct {
	GwAddress string
	GwPort	string
	GwId	string
	Ip		string
	MacAddress string
	URL		string
}

type AuthParams struct {
	Stage string
	Ip	string
	MacAddress	string
	Token 	string
	Incoming string
	Outgoing string
	GwId string
}

func NewRouter(protocol *WDLoginProtocol) http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/ping", protocol.Ping).Methods("GET", "POST")
	router.HandleFunc("/login", protocol.Login).Methods("GET")
	router.HandleFunc("/auth", protocol.Auth).Methods("GET")

	return router
}


func (wdlogin *WDLoginProtocol) Ping(w http.ResponseWriter, r *http.Request) {
	res := "Pong"
	log.Println("answer:", res)
	w.Write([]byte(res))
}

func (wdlogin *WDLoginProtocol) Login(w http.ResponseWriter, r *http.Request) {

	params := LoginParams{
		GwAddress: r.URL.Query().Get("gw_address"),
		GwPort: r.URL.Query().Get("gw_port"),
		GwId: r.URL.Query().Get("gw_id"),
		Ip: r.URL.Query().Get("ip"),
		MacAddress: r.URL.Query().Get("mac"),
		URL: r.URL.Query().Get("url"),
		}

	redirectURL := url.URL{
		Scheme:     "http",
		Host:       params.GwAddress + ":" + params.GwPort,
		Path:       "/wifidog/auth",
	}

	query := redirectURL.Query()
	query.Set("token", "1234")
	redirectURL.RawQuery = query.Encode()

	//res := "Login page"
	log.Println("Page request", params)
	log.Println("Redirect url", redirectURL.String())
	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}

func (wdlogin *WDLoginProtocol) Auth(w http.ResponseWriter, r *http.Request) {

	params := AuthParams{
		Stage:      r.URL.Query().Get("stage"),
		Ip:         r.URL.Query().Get("ip"),
		MacAddress: r.URL.Query().Get("mac"),
		Token:      r.URL.Query().Get("token"),
		Incoming:   r.URL.Query().Get("incoming"),
		Outgoing:   r.URL.Query().Get("outgoing"),
		GwId:       r.URL.Query().Get("gw_id"),
	}

	log.Println("Auth route", params)

	//response := "Auth: " + string(authTypes.AuthAllowed)
	response := "Auth: 1"
	log.Println("Auth response", response)
	w.Write([]byte(response))
}