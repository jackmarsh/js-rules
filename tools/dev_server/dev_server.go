package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/thought-machine/go-flags"
)

var opts struct {
	Port   int    `short:"p" long:"port" description:"Port to run the server on" default:"8080"`
	Dist   string `short:"d" long:"dist" description:"Distribution to serve."`
	Static string `short:"s" long:"static" description:"Static files to serve."`
	Proxy  string `long:"proxy" description:"Path to proxy configuration file"`
}

// ProxyConfig represents the structure of the proxy configuration
type ProxyConfig struct {
	Proxy       string            `json:"proxy"`
	Host        string            `json:"host"`
	Protocol    string            `json:"protocol"`
	Port        int               `json:"port"`
	Headers     map[string]string `json:"headers"`
	PathRewrite map[string]string `json:"pathRewrite"`
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func main() {
	start := time.Now()
	p := flags.NewParser(&opts, flags.Default)
	_, err := p.Parse()
	if err != nil {
		os.Exit(1)
	}

	// Load proxy configuration if provided.
	var proxyConfig *ProxyConfig
	if opts.Proxy != "" {
		proxyConfig, err = loadProxyConfig(opts.Proxy)
		if err != nil {
			log.Fatalf("Failed to load proxy configuration: %v", err)
		}
	}

	// Set up the file server and routes.
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(opts.Dist))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	if proxyConfig != nil {
		setupProxyRoutes(mux, proxyConfig)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Prevent the catch-all handler from intercepting proxy paths
		if proxyConfig != nil {
			if r.URL.Path == proxyConfig.Proxy || strings.HasPrefix(r.URL.Path, proxyConfig.Proxy+"/") {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, fmt.Sprintf("%s/index.html", opts.Static))
	})

	// Start the server.
	port := opts.Port
	address := fmt.Sprintf(":%d", port)
	loopbackURL := fmt.Sprintf("http://localhost%s", address)
	fmt.Printf("%s[js-rules dev-server] Server is running at the following addresses:\n", ColorGreen)
	fmt.Printf("%s[js-rules dev-server] \t%sLoopback: %s%s\n", ColorGreen, ColorReset, ColorCyan, loopbackURL)
	if ipv4, ok := getLocalIPv4(); ok {
		fmt.Printf("%s[js-rules dev-server] \t%sOn Your Network (IPv4): %shttp://%s%s\n", ColorGreen, ColorReset, ColorCyan, ipv4, address)
	}
	fmt.Printf("%s[js-rules dev-server] Static content being served from '%s%s%s' directory\n", ColorGreen, ColorCyan, opts.Dist, ColorGreen)
	fmt.Printf("%s[js-rules dev-server] 404s will fallback to '%s%s%s'\n", ColorGreen, ColorCyan, fmt.Sprintf("%s/index.html", opts.Static), ColorGreen)
	fmt.Printf("%s[js-rules dev-server] Compiled successfully in %s%s\n", ColorGreen, ColorCyan, time.Since(start))

	openBrowser(loopbackURL)
	if err := http.ListenAndServe(address, mux); err != nil {
		panic(err)
	}
}

func getLocalIPv4() (string, bool) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("%sgetting network interfaces:", ColorRed, err)
		return "", false
	}
	for _, iface := range interfaces {
		// Skip interfaces that are down or not connected
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addresses, err := iface.Addrs()
		if err != nil {
			fmt.Printf("%sgetting addresses for interface %s: %v\n", ColorRed, iface.Name, err)
			continue
		}
		for _, addr := range addresses {
			// Check if the address is an IPv4 address
			ip, ok := addr.(*net.IPNet)
			if ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
				return ip.IP.String(), true
			}
		}
	}
	return "", false
}

func openBrowser(url string) {
	var err error
	switch os := runtime.GOOS; os {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		fmt.Printf("%s[js-rules dev-server] Unable to open browser on your platform.\n", ColorRed)
		return
	}
	if err != nil {
		fmt.Printf("%s[js-rules dev-server] Error opening browser: %v\n", ColorRed, err)
	}
}

func loadProxyConfig(path string) (*ProxyConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config ProxyConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func setupProxyRoutes(mux *http.ServeMux, config *ProxyConfig) {
	proxyRegex, err := regexp.Compile(config.Proxy)
	if err != nil {
		log.Fatalf("Invalid proxy regex: %v", err)
	}
	target := fmt.Sprintf("%s://%s:%d", config.Protocol, config.Host, config.Port)
	backendURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid backend URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(backendURL)

	proxy.ModifyResponse = func(resp *http.Response) error {
		for key, value := range config.Headers {
			resp.Header.Set(key, value)
		}
		return nil
	}
	
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Printf("%s[js-rules dev-server] [HPM] %shttp: error: %s%s\n",
			ColorGreen, ColorRed, err.Error(), ColorReset,
		)
		http.Error(w, "Service Unavailable", http.StatusBadGateway)
	}
	
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Adjust the Host header
		req.Host = backendURL.Host
		// Apply path rewrites
		originalPath := req.URL.Path
		for pattern, replacement := range config.PathRewrite {
			re, err := regexp.Compile(pattern)
			if err != nil {
				fmt.Printf("%s[js-rules dev-server] Invalid rewrite pattern: %v%s\n", ColorRed, err, ColorReset)
				continue
			}
			if re.MatchString(req.URL.Path) {
				req.URL.Path = re.ReplaceAllString(req.URL.Path, replacement)
				fmt.Printf("%s[js-rules dev-server] [HPM]%s Forwarding request: %s\"%s\"%s -> %s\"%s\"%s\n",
					ColorGreen, ColorReset,
					ColorCyan, originalPath, ColorReset,
					ColorCyan, req.URL.Path, ColorReset,
				)
				break
			}
		}
		for key, value := range config.Headers {
			req.Header.Set(key, value)
		}
	}

	proxyHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			for key, value := range config.Headers {
				w.Header().Set(key, value)
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		if !proxyRegex.MatchString(r.URL.Path) {
			http.NotFound(w, r)
			return
		}
		proxy.ServeHTTP(w, r)
	}

	// Register the proxy handler for both "/<proxy>" and "/<proxy>/"
	mux.HandleFunc(config.Proxy, proxyHandler)
	mux.HandleFunc(config.Proxy+"/", proxyHandler)

	fmt.Printf("%s[js-rules dev-server] [HPM]%s Proxy created: %s%s%s  ->  %s%s%s\n",
		ColorGreen, ColorReset,
		ColorCyan, config.Proxy, ColorReset,
		ColorCyan, target, ColorReset,
	)
	for pattern, replacement := range config.PathRewrite {
		fmt.Printf("%s[js-rules dev-server] [HPM]%s Proxy rewrite rule created: %s\"%s\"%s ~> %s\"%s\"%s\n",
			ColorGreen, ColorReset,
			ColorCyan, pattern, ColorReset,
			ColorCyan, replacement, ColorReset,
		)
	}
}
