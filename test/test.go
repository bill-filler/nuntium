package main

import (
	"fmt"
	"os"

	"launchpad.net/go-dbus/v1"
)

func main() {
	var pdu string
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [busname]\n", os.Args[0])
		os.Exit(1)
	} else if len(os.Args) == 3 {
		pdu = os.Args[2]
	} else {
		pdu = "argentina-personal"
	}

	var err error
	var conn *dbus.Connection
	if conn, err = dbus.Connect(dbus.SystemBus); err != nil {
		fmt.Println("Connection error:", err)
		os.Exit(1)
	}

	obj := conn.Object(os.Args[1], "/nuntium")

	var data []byte
	switch pdu {
	case "spain-vodaphone":
		//VodaPhone España
		data = []byte{
			0x00, 0x06, 0x26, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
			0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d, 0x6d, 0x73,
			0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf, 0x84, 0xb4, 0x81,
			0x8d, 0xdf, 0x8c, 0x82, 0x98, 0x4e, 0x4f, 0x4b, 0x35, 0x43, 0x64, 0x7a, 0x30,
			0x38, 0x42, 0x41, 0x73, 0x77, 0x61, 0x62, 0x77, 0x55, 0x48, 0x00, 0x8d, 0x90,
			0x89, 0x18, 0x80, 0x2b, 0x33, 0x34, 0x36, 0x30, 0x30, 0x39, 0x34, 0x34, 0x34,
			0x36, 0x33, 0x2f, 0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00,
			0x8a, 0x80, 0x8e, 0x02, 0x74, 0x00, 0x88, 0x05, 0x81, 0x03, 0x02, 0xa3, 0x00,
			0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6d, 0x6d, 0x31, 0x66, 0x65,
			0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x6c, 0x65, 0x74, 0x73, 0x2f, 0x4e, 0x4f,
			0x4b, 0x35, 0x43, 0x64, 0x7a, 0x30, 0x38, 0x42, 0x41, 0x73, 0x77, 0x61, 0x62,
			0x77, 0x55, 0x48, 0x00,
		}
	case "argentina-personal":
		//Personal Argentina
		data = []byte{
			0x01, 0x06, 0x26, 0x61, 0x70, 0x70, 0x6C, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
			0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d, 0x6d, 0x73,
			0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf, 0x84, 0xb4, 0x86,
			0xc3, 0x95, 0x8c, 0x82, 0x98, 0x6d, 0x30, 0x34, 0x42, 0x4b, 0x6b, 0x73, 0x69,
			0x6d, 0x30, 0x35, 0x40, 0x6d, 0x6d, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
			0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x00, 0x8d, 0x90,
			0x89, 0x19, 0x80, 0x2b, 0x35, 0x34, 0x33, 0x35, 0x31, 0x35, 0x39, 0x32, 0x34,
			0x39, 0x30, 0x36, 0x2f, 0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e,
			0x00, 0x8a, 0x80, 0x8e, 0x02, 0x74, 0x00, 0x88, 0x05, 0x81, 0x03, 0x02, 0xa2,
			0xff, 0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x37, 0x32, 0x2e,
			0x32, 0x35, 0x2e, 0x37, 0x2e, 0x31, 0x33, 0x31, 0x2f, 0x3f, 0x6d, 0x65, 0x73,
			0x73, 0x61, 0x67, 0x65, 0x2d, 0x69, 0x64, 0x3d, 0x6d, 0x30, 0x34, 0x42, 0x4b,
			0x68, 0x34, 0x33, 0x65, 0x30, 0x33, 0x00,
		}
	case "usa-att":
		//USA AT&T
		data = []byte{
			0x01, 0x06, 0x27, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
			0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d, 0x6d, 0x73,
			0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf, 0x84, 0x8d, 0x01,
			0x82, 0xb4, 0x84, 0x8c, 0x82, 0x98, 0x44, 0x32, 0x30, 0x34, 0x30, 0x37, 0x31,
			0x36, 0x35, 0x36, 0x32, 0x34, 0x36, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30,
			0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x00, 0x8d, 0x90, 0x89, 0x18, 0x80,
			0x2b, 0x31, 0x37, 0x37, 0x34, 0x32, 0x37, 0x30, 0x30, 0x36, 0x35, 0x39, 0x2f,
			0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00, 0x96, 0x02, 0xea,
			0x00, 0x8a, 0x80, 0x8e, 0x02, 0x80, 0x00, 0x88, 0x05, 0x81, 0x03, 0x05, 0x46,
			0x00, 0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x36, 0x36, 0x2e,
			0x32, 0x31, 0x36, 0x2e, 0x31, 0x36, 0x36, 0x2e, 0x36, 0x37, 0x3a, 0x38, 0x30,
			0x30, 0x34, 0x2f, 0x30, 0x34, 0x30, 0x37, 0x31, 0x36, 0x35, 0x36, 0x32, 0x34,
			0x36, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
			0x30, 0x30, 0x00,
		}
	default:
		fmt.Println("Choose between argentina-personal and spain-vodaphone")
	}

	info := map[string]*dbus.Variant{"LocalSentTime": &dbus.Variant{"2014-02-05T08:29:55-0300"},
		"Sender": &dbus.Variant{"+543515924906"}}

	reply, err := obj.Call("org.ofono.PushNotificationAgent", "ReceiveNotification", data, info)
	if err != nil || reply.Type == dbus.TypeError {
		fmt.Printf("Notification error: %s", err)
		os.Exit(1)
	}
}
