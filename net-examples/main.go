package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const (
	udpPacketBufSize = 65536
	udpRecvBufSize   = 2 * 1024 * 1024
)

type runType string

const (
	clientRunType runType = "client"
	serverRunType runType = "server"
)

type conn struct {
	tcp *net.TCPListener
	udp *net.UDPConn
}

type packet struct {
	buf       []byte
	from      net.Addr
	timestamp time.Time
}

type msgType int8

const (
	clientMsgType msgType = 1
	serverMsgType msgType = 2
)

type message struct {
	MsgType msgType `json:"msg_type"`
	Msg     string  `json:"msg"`
}

var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "", log.LstdFlags)

	var r string
	var addr string
	flag.StringVar(&r, "type", "server", "run type")
	flag.StringVar(&addr, "addr", "", "server addr")
	flag.Parse()
	if runType(r) == serverRunType {
		serverMain()
	} else {
		if addr == "" {
			panic("server addr is required")
		}
		clientMain(addr)
	}
}

func serverMain() {
	c, err := buildConnection()
	if err != nil {
		panic(err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	tcpCh := make(chan *net.TCPConn, 10)
	packetCh := make(chan packet, 10)
	go tcpListen(ctx, c.tcp, tcpCh)
	go udpListen(ctx, c.udp, packetCh)
	go func() {
		defer close(tcpCh)
		for {
			select {
			case con := <-tcpCh:
				go readTcp(ctx, con)
			}
		}
	}()
	go func() {
		defer close(packetCh)
		for {
			select {
			case <-ctx.Done():
				return
			case packet := <-packetCh:
				var m message
				err = json.Unmarshal(packet.buf, &m)
				if err != nil {
					logger.Printf("failed to json.Unmarshal udp :%v", err)
				} else {
					logger.Printf("recieved upd from %s msg %v", packet.from.String(), m)
				}
			}

		}
	}()
	logger.Printf("done setup")
	<-ctx.Done()
	stop()
	time.Sleep(2 * time.Second)
	logger.Printf("done server")
}

func clientMain(addr string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	go func() {
		c, err := net.Dial("tcp", addr)
		if err != nil {
			return
		}
		i := 0
		for {
			time.Sleep(2 * time.Second)
			select {
			case <-ctx.Done():
				return
			default:
				m := message{
					MsgType: clientMsgType,
					Msg:     fmt.Sprintf("tcp msg %d", i),
				}
				b, err := json.Marshal(&m)
				if err != nil {
					logger.Printf("failed to json.Marshal :%v", err)
				}
				_, err = c.Write(b)
				if err != nil {
					logger.Printf("failed to Write :%v", err)
				}
				logger.Printf("write tcp msg %d", i)
			}
			i++
		}
	}()
	go func() {
		c, err := net.Dial("udp", addr)
		if err != nil {
			return
		}
		i := 0
		for {
			time.Sleep(3 * time.Second)
			select {
			case <-ctx.Done():
				return
			default:
				m := message{
					MsgType: clientMsgType,
					Msg:     fmt.Sprintf("udp msg %d", i),
				}
				b, err := json.Marshal(&m)
				if err != nil {
					logger.Printf("failed to json.Marshal :%v", err)
				}
				_, err = c.Write(b)
				if err != nil {
					logger.Printf("failed to Write :%v", err)
				}
				logger.Printf("write udp msg %d", i)
			}
			i++
		}
	}()
	logger.Printf("done setup")
	<-ctx.Done()
	stop()
	logger.Printf("done client")
}

func buildConnection() (*conn, error) {
	ip := net.ParseIP("0.0.0.0")
	tcpAddr := &net.TCPAddr{IP: ip, Port: 0}
	tcpLn, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to start TCP listener on : %v", err)
	}

	port := tcpLn.Addr().(*net.TCPAddr).Port

	udpAddr := &net.UDPAddr{IP: ip, Port: port}
	udpLn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to start UDP listener on : %v", err)
	}
	if err := setUDPRecvBuf(udpLn); err != nil {
		return nil, fmt.Errorf("failed to resize UDP buffer: %v", err)
	}
	logger.Printf("server port %d", port)
	return &conn{
		tcp: tcpLn,
		udp: udpLn,
	}, nil

}

func setUDPRecvBuf(c *net.UDPConn) error {
	size := udpRecvBufSize
	var err error
	for size > 0 {
		if err = c.SetReadBuffer(size); err == nil {
			return nil
		}
		size = size / 2
	}
	return err
}

func tcpListen(ctx context.Context, tcpLn *net.TCPListener, ch chan *net.TCPConn) {
	const baseDelay = 5 * time.Millisecond
	const maxDelay = 1 * time.Second

	var loopDelay time.Duration
	for {
		conn, err := tcpLn.AcceptTCP()
		if err != nil {
			select {
			case <-ctx.Done():
				break
			default:
				if loopDelay == 0 {
					loopDelay = baseDelay
				} else {
					loopDelay *= 2
				}

				if loopDelay > maxDelay {
					loopDelay = maxDelay
				}
				logger.Printf("failed to accept TCP: %v", err)
				time.Sleep(loopDelay)
				continue
			}
		}
		loopDelay = 0

		ch <- conn
	}
}

func udpListen(ctx context.Context, udpLn *net.UDPConn, ch chan packet) {
	defer udpLn.Close()
	for {
		logger.Printf("read udp")
		buf := make([]byte, udpPacketBufSize)
		n, addr, err := udpLn.ReadFrom(buf)
		ts := time.Now()
		if err != nil {
			select {
			case <-ctx.Done():
				break
			default:
				logger.Printf("failed to read UPD packet: %v", err)
				continue
			}
		}

		if n < 1 {
			logger.Printf("error UDP packet too short (%d bytes) from %s", len(buf), addr.String())
			continue
		}

		ch <- packet{
			buf:       buf[:n],
			from:      addr,
			timestamp: ts,
		}
	}
}

func readTcp(ctx context.Context, conn *net.TCPConn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			l, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}
				logger.Printf("failed to conn.Read :%v", err)
				continue
			}
			var m message
			err = json.Unmarshal(buf[:l], &m)
			if err != nil {
				logger.Printf("failed to json.Unmarshal tcp :%v", err)
				continue
			}
			logger.Printf("received tcp msg %v", m)
		}
	}
}
