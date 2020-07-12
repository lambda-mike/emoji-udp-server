package server

import (
	"net"
	"testing"
)

type mockHandler struct {
	state chan string
}

func (h *mockHandler) Handle(s string) {
	h.state <- s
}

func TestListen(t *testing.T) {
	t.Log("Listen")
	{
		t.Log("Given new server with mocked handler")
		{
			t.Log("When called with some msg")
			{
				msg := "testing message"
				handler := mockHandler{}
				handler.state = make(chan string)
				sut, port := CreateUDPServer(0, &handler)
				go sut.Listen()

				udpAddr := net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: port, Zone: ""}
				conn, err := net.DialUDP("udp", nil /*(laddr)*/, &udpAddr)
				defer conn.Close()
				if err != nil {
					t.Fatal("Could not establish connection to server", err)
				}
				t.Log("Established connection with local server at port:", port)
				conn.Write([]byte(msg))
				state := <-handler.state
				if state != msg {
					t.Fatal("Handler did not receive proper msg! Got:", state, "should be:", msg)
				}
				t.Log("Handler should receive proper message in its state")
			}
		}
	}
}
