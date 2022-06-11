package main

import (
	"log"

	"radius-server/pkg/therouter"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc4679"
)

func main() {
	handler := func(w radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)
		// password := rfc2865.UserPassword_GetString(r.Packet)
		remoteId := rfc4679.ADSLAgentRemoteID_GetString(r.Packet)
		circuitId := rfc4679.ADSLAgentCircuitID_GetString(r.Packet)
		// nasId := cisco.CiscoNASPort_GetString
		nasId := rfc2865.NASIdentifier_GetString(r.Packet)
		log.Printf("User-name:[%v], NAS-Identifier:[%v], Remote-Id:[%v], Circuit-Id:[%v]", username, nasId, remoteId, circuitId)
		var code radius.Code
		// if username == "tim" && password == "12345" {
		code = radius.CodeAccessAccept
		// } else {
		// code = radius.CodeAccessReject
		// }
		log.Printf("Writing %v to %v", code, r.RemoteAddr)

		resp := r.Response(code)

		w.Write(resp)
		therouter.TherouterIngressCir_Set(r.Packet, 1000)
		w.Write(r.Response(code))
	}

	server := radius.PacketServer{
		Handler:      radius.HandlerFunc(handler),
		SecretSource: radius.StaticSecretSource([]byte(`$radiusSecret$`)),
	}

	log.Printf("Starting server on :1812")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
