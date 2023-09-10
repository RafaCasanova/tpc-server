package server

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	netListen  net.Listener
	quitchan   chan struct{}
	msgchan    chan Message
}

type Message struct {
	from    string
	payload []byte
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitchan:   make(chan struct{}),
		msgchan:    make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	nl, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer nl.Close()

	s.netListen = nl

	go s.accetpLoop()

	<-s.quitchan

	close(s.msgchan)

	return nil
}

func (s *Server) accetpLoop() {
	for {
		conn, err := s.netListen.Accept()
		if err != nil {
			fmt.Println("Error aceitavel", err)
			continue
		}
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error", err)
			continue
		}

		// comand := strings.ReplaceAll(strings.ReplaceAll(string(buf[:n]), "\n", ""), " ", "")

		// fmt.Println(strings.Compare(comand, "exit"))
		// if strings.Compare()comand == "exit" {
		// 	conn.Write([]byte("Fechando connexão \n"))
		// 	net.Conn.Close(conn)
		// } else {
		s.msgchan <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
		conn.Write([]byte("Menssagem recebida \n"))
		// }
	}
}

func (s *Server) GetMensagemChan() {
	go func() {
		for msg := range s.msgchan {
			fmt.Printf("(%s):%s", msg.from, string(msg.payload))
		}
	}()
}
