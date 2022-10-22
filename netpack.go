package main

import(
	"fmt"
	"log"
	"os"
	"net"
)

func main(){

	domain_info()

}


func domain_info(){

	if len(os.Args) != 2 {
		log.Fatal("you didn't provide any argument!")
	}

	arg := os.Args[1]

// look up host ips
	fmt.Println("checking for IP")
	fmt.Println(" ")
	host_ips, err := net.LookupHost(arg)
	if err != nil {
		log.Fatal(err)
	}

	for _, host_ip := range host_ips{
		fmt.Println("Ip address found .....")
		fmt.Println(" ")
		fmt.Println(host_ip)
		fmt.Println(" ")

	}

// look up mx records
	fmt.Println("checking for mx records")
	fmt.Println(" ")
	mxRecords , err := net.LookupMX(arg)
	if err != nil {
		log.Fatal(err)
	}

	for _, mx := range mxRecords{
		fmt.Println("mx records found .....")
		fmt.Println(" ")
		fmt.Printf("Host: %s\t Prefrence: %d", mx.Host, mx.Pref)
		fmt.Println(" ")
	}

// look up name servers
	fmt.Println("checking for nameservers ")
	fmt.Println(" ")
	nameservers, err := net.LookupNS(arg)
	if err != nil{
		log.Fatal(err)
	}

	for _, nameserver := range nameservers{
		fmt.Println("nameservers found .....")
		fmt.Println(" ")
		fmt.Println(nameserver.Host)
		fmt.Println(" ")
	}

	fmt.Println(" DONE ..... ")

}