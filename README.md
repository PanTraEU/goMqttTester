# Simple MQTT Tester based on GO

## Test Idee
Der Test Producer sender jede Sekunde einen frischen Int64 (counter++). Die Idee ist testen zu können ob jeder einzelne Client wirklich keine Messages verpasst.

Jeder ESP32 Client muss sich mit einer eindeutigen statischen ClientID beim MQTT anmelden. Die Session muss mit der Eigenschaft cleanSession = false und QoS 1 oder 2 initialisiert werden. 

Nachdem der Client sich das erste Mal angemeldet hat, sollte der MQTT Server ab dem Moment bezogen auf die ClientID alle ab dem Moment anfallenden Messages speichern, bis der Client sie sich abgeholt hat. Der Server bewahrt die Messages im Moment pro Client max. 7 Tage auf (das kann man global fest legen). 

Der Test sollte nun also so aussehen:

### 1. Verbinden und für einen Moment Messages lesen:

167
168

69
170

### 2. dann Verbindung schließen, Moment warten und dann erneute verbinden und lesen, dann sollte es natlos weiter gehen:

171
172
173
174
....


Das ganze mit 2-3 Clients parallel und jeder sollte alle Messages lückenlos bekommen.

Der Server fängt nach dem ersten Connect eines Clients (der sich über seine ClientID und nicht den Usernamen) bezüglich der Persitenz identifiziert an die Messages zu speichern.
Also die ab dem ersten Connect anfallenden neuen Messages für den Client zu bewahren, bis dieser wieder vorbeikommt. 

QoS 1 = es kann eine Message doppelt ausgeliefert werden, aber min. 1
QoS 2 = jede Message wird jedem Client genau einmal ausgeliefert

Das script "testIt-tls.sh" bildet den Ablauf eines solchen Tests grob ab. 
