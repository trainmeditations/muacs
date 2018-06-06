package muacs

import (
	"encoding/xml"
	"net/http"
)

//ADRequest represents the values in a POX autodiscover request
type ADRequest struct {
}

//ADResponse represents the values needed for a POX autodisover response
type ADResponse struct {
	XMLName xml.Name `xml:"http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a Response"`
}

//Autodiscover represents both message types for POX autodiscover
type Autodiscover struct {
	XMLName  xml.Name     `xml:"http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006 Autodiscover"`
	Request  []ADRequest  `xml:",omitempty"`
	Response []ADResponse `xml:",omitempty"`
}

//AutodiscoverHandler handles the http transaction for a POX Autodiscover request
type AutodiscoverHandler struct{}

func (ah AutodiscoverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

//AutodiscoverResponse should return an ADResponse based on the email
//address and server configs
//
//To avoid the ability to fish for valid emails, the validity of the email
//isn't checked, the email is just used in generating the response based
//on the domain
func AutodiscoverResponse(email string, configs []ServerConfig) ADResponse {
	return ADResponse{}
}

//AutodiscoverResponseXML produces the POX Autodiscover Response xml.
func AutodiscoverResponseXML(r ADResponse) ([]byte, error) {
	a := Autodiscover{
		Response: []ADResponse{r},
	}
	return xml.Marshal(a)
}
