package main

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Printf(
		`
 ▄▄▄· .▄▄ ·  ▄▄· ▪  ▪      ▄▄▌   ▄▄▄·  ▐ ▄
▐█ ▀█ ▐█ ▀. ▐█ ▌▪██ ██     ██•  ▐█ ▀█ •█▌▐█
▄█▀▀█ ▄▀▀▀█▄██ ▄▄▐█·▐█·    ██▪  ▄█▀▀█ ▐█▐▐▌
▐█ ▪▐▌▐█▄▪▐█▐███▌▐█▌▐█▌    ▐█▌▐▌▐█ ▪▐▌██▐█▌
 ▀  ▀  ▀▀▀▀ ·▀▀▀ ▀▀▀▀▀▀    .▀▀▀  ▀  ▀ ▀▀ █▪

`)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	var (
		resultBytes []byte
		errorBytes  []byte
	)

	s, err := nmap.NewScanner(
		nmap.WithContext(ctx),
		nmap.WithServiceInfo(),
		nmap.WithScripts("vulners"),
		nmap.WithTargets("172.16.1.*"),
		nmap.WithStylesheet("oX-html.xsl"),
		nmap.WithFilterHost(func(h nmap.Host) bool {
			for idx := range h.Ports {
				if h.Ports[idx].Status() == "open" {
					return true
				}
			}
			return false
		}),
	)

	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	fmt.Println("Begining nmap scan")

	// Executes asynchronously, allowing results to be streamed in real time.
	if err := s.RunAsync(); err != nil {
		panic(err)
	}

	// Connect to stdout of scanner.
	stdout := s.GetStdout()

	// Connect to stderr of scanner.
	stderr := s.GetStderr()

	// Goroutine to watch for stdout and print to screen. Additionally it stores
	// the bytes intoa variable for processiing later.
	go func() {
		for stdout.Scan() {
			//fmt.Println(stdout.Text())
			resultBytes = append(resultBytes, stdout.Bytes()...)
		}
	}()

	// Goroutine to watch for stderr and print to screen. Additionally it stores
	// the bytes intoa variable for processiing later.
	go func() {
		for stderr.Scan() {
			errorBytes = append(errorBytes, stderr.Bytes()...)
		}
	}()

	// Blocks main until the scan has completed.
	if err := s.Wait(); err != nil {
		panic(err)
	}

	// Parsing the results into corresponding structs.
	result, err := nmap.Parse(resultBytes)

	// Parsing the results into the NmapError slice of our nmap Struct.
	result.NmapErrors = strings.Split(string(errorBytes), "\n")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n\nXML\n", len(result.Hosts), result.Stats.Finished.Elapsed)

	err = nmapWebReport(ctx, result)
	if err != nil {
		log.Fatal("Failed to generte web report", err)
	}

	generateDiagram(ctx, result)
	ascii2png(ctx)
}

func generatePortInfo(ctx context.Context, host nmap.Host) string {
	var portASCII []string
	for _, port := range host.Ports {
		product := port.Service.Product
		if len(port.Service.Product) > 8 {
			product = port.Service.Product[0:8]
		}
		details := fmt.Sprintf("[+] %d:%s %s", port.ID, strings.Replace(port.Service.Name, "-", " ", -1), product)
		details = genSpacing(details)
		portASCII = append(portASCII, fmt.Sprintf("%s", details))
	}
	return strings.Join(portASCII, "\n\t\t| ")
}
func generateHostBox(ctx context.Context, host nmap.Host) string {
	if len(host.Ports) == 0 || len(host.Addresses) == 0 {
		return ""
	}

	portInfoASCII := generatePortInfo(ctx, host)
	hostinfo := "    <" + host.Addresses[0].Addr + ">"
	hostinfo = genSpacing(hostinfo)
	hostBoxASCII := fmt.Sprintf(`
		+-----------------------------+
		| %s
		| %s
		|                             |
		+---------------+-------------+
`, hostinfo, portInfoASCII)
	return hostBoxASCII
}

func generateDiagram(ctx context.Context, result *nmap.Run) error {
	// Use the results to print an example output
	var networkASCII []string

	for _, host := range result.Hosts {
		networkASCII = append(networkASCII, generateHostBox(ctx, host))
	}

	diagramAll := strings.Join(networkASCII[0:], "\t\t\t\t|\n\t\t\t\t|\n\t\t\t\t|\n\t\t\t\tV")
	fmt.Printf("%s", diagramAll)
	return ioutil.WriteFile("ASCII_LAN.txt", []byte(diagramAll), 0666)

}

func nmapWebReport(ctx context.Context, result *nmap.Run) error {
	err := result.ToFile("LANscan.xml")
	if err != nil {
		return err
	}

	xml2html := exec.CommandContext(ctx, "bash", "-c", "xsltproc LANscan.xml -o LAN_WebReport.html")
	fmt.Println("Converting xml scan data into html report")

	xml2html.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func ascii2png(ctx context.Context) error {
	cmdstr := "diagram -in ASCII_LAN.txt -out LAN_DRAWING.png -preview=\"false\""

	ascii2png := exec.CommandContext(ctx, "bash", "-c", cmdstr)
	fmt.Println("Drawing Network Diagram to LAN_DRAWING.png")

	stdout, err := ascii2png.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))
	return nil
}

func genSpacing(curLine string) string {
	remaningLength := 28 - len(curLine)
	i := 0

	for i < remaningLength {
		curLine = curLine + " "
		i++
	}
	curLine = curLine + "|"
	return curLine
}
