package main

import (
	"log"

	"radius-server/pkg/therouter"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

func main() {
	handler := func(w radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)
		password := rfc2865.UserPassword_GetString(r.Packet)
		log.Printf("User-name:[%v], Password:[%v]", username, password)
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
