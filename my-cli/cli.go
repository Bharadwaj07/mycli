package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"
	app.Version = "1.0.0"


    myFlags := []cli.Flag{
        &cli.StringFlag{
            Name:  "host",
            Value: "tutorialedge.net",
        },
    }

    
    app.Commands = []*cli.Command{
        {
            Name:  "ns",
            Usage: "Looks Up the NameServers for a Particular Host",
            Flags: myFlags,
            
            Action: func(c *cli.Context) error {
               
                ns, err := net.LookupNS(c.String("host"))
                if err != nil {
                    return err
                }
				fmt.Println("Names server are :")
                for i := 0; i < len(ns); i++ {
                    fmt.Println(ns[i].Host)
                }
                return nil
            },
		},
		{
            Name:  "ip",
            Usage: "Looks Up the IPs for a Particular Host",
            Flags: myFlags,
            
            Action: func(c *cli.Context) error {
               
                ip, err := net.LookupIP(c.String("host"))
                if err != nil {
                    return err
                }
				fmt.Println("IPs are :")
                for i := 0; i < len(ip); i++ {
                    fmt.Println(ip[i])
                }
                return nil
            },
		},
		{
            Name:  "cname",
            Usage: "Looks Up the CNAME for a Particular Host",
            Flags: myFlags,
            
            Action: func(c *cli.Context) error {
               
                cname, err := net.LookupCNAME(c.String("host"))
                if err != nil {
                    return err
                }
				fmt.Println("CName :")
                fmt.Println(cname)
                return nil
            },
		},
		{
            Name:  "mx",
            Usage: "Looks Up the MX record for a Particular Host",
            Flags: myFlags,
            
            Action: func(c *cli.Context) error {
               
                mx, err := net.LookupMX(c.String("host"))
                if err != nil {
                    return err
                }
				fmt.Println("MX records  are :")
                for i := 0; i < len(mx); i++ {
                    fmt.Println(mx[i].Host,mx[i].Pref)
                }
                return nil
            },
        },
    }

   
    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
