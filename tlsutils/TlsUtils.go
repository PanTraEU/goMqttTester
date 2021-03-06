package TlsUtils

import (
	"../configUtil"
	"crypto/tls"
	"crypto/x509"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"log"
	"time"
)

func newTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certPool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("./certs/ca.crt")
	if err == nil {
		certPool.AppendCertsFromPEM(pemCerts)
	}

	clientCACert, err := ioutil.ReadFile("./certs/client1.crt")
	if err != nil {
		log.Fatal(err)
	}
	clientCACertPool := x509.NewCertPool()
	clientCACertPool.AppendCertsFromPEM(clientCACert)

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("./certs/client1.crt", "./certs/client1.key")
	if err != nil {
		panic(err)
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	//fmt.Println(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certPool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.RequireAndVerifyClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}

func MqttOpts(mqttHost string, prodId string, conf configUtil.Configuration, cleanSession bool) *mqtt.ClientOptions {
	optsProd := mqtt.NewClientOptions().
		AddBroker(mqttHost).
		SetClientID(prodId).
		SetAutoReconnect(true)
	optsProd.SetCleanSession(cleanSession)
	optsProd.SetKeepAlive(2 * time.Second)
	optsProd.SetPingTimeout(1 * time.Second)
	if conf.UseAuth {
		optsProd.SetUsername(conf.MqttUser)
		optsProd.SetPassword(conf.MqttPassword)
	}
	if conf.UseTls {
		optsProd.SetTLSConfig(newTLSConfig())
	}
	return optsProd
}
