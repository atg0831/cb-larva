package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/cloud-barista/cb-larva/poc-cb-net/internal/app"
	cbnet "github.com/cloud-barista/cb-larva/poc-cb-net/internal/cb-network"
	model "github.com/cloud-barista/cb-larva/poc-cb-net/internal/cb-network/model"
	etcdkey "github.com/cloud-barista/cb-larva/poc-cb-net/internal/etcd-key"
	"github.com/cloud-barista/cb-larva/poc-cb-net/internal/file"
	cblog "github.com/cloud-barista/cb-log"
	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// CBNet represents a network for the multi-cloud.
var CBNet *cbnet.CBNetwork
var channel chan bool
var mutex = &sync.Mutex{}

// CBLogger represents a logger to show execution processes according to the logging level.
var CBLogger *logrus.Logger
var config model.Config

func init() {
	fmt.Println("Start......... init() of agent.go")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	fmt.Printf("exePath: %v\n", exePath)

	// Load cb-log config from the current directory (usually for the production)
	logConfPath := filepath.Join(exePath, "config", "log_conf.yaml")
	fmt.Printf("logConfPath: %v\n", logConfPath)
	if !file.Exists(logConfPath) {
		// Load cb-log config from the project directory (usually for development)
		path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			panic(err)
		}
		projectPath := strings.TrimSpace(string(path))
		logConfPath = filepath.Join(projectPath, "poc-cb-net", "config", "log_conf.yaml")
	}
	CBLogger = cblog.GetLoggerWithConfigPath("cb-network", logConfPath)
	CBLogger.Debugf("Load %v", logConfPath)

	// Load cb-network config from the current directory (usually for the production)
	configPath := filepath.Join(exePath, "config", "config.yaml")
	fmt.Printf("configPath: %v\n", configPath)
	if !file.Exists(configPath) {
		// Load cb-network config from the project directory (usually for the development)
		path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			panic(err)
		}
		projectPath := strings.TrimSpace(string(path))
		configPath = filepath.Join(projectPath, "poc-cb-net", "config", "config.yaml")
	}
	config, _ = model.LoadConfig(configPath)
	CBLogger.Debugf("Load %v", configPath)
	fmt.Println("End......... init() of agent.go")
}

func decodeAndSetNetworkingRule(key string, value []byte, hostID string) {
	mutex.Lock()
	CBLogger.Debug("Start.........")
	slicedKeys := strings.Split(key, "/")
	parsedHostID := slicedKeys[len(slicedKeys)-1]
	CBLogger.Tracef("ParsedHostID: %v", parsedHostID)

	var networkingRule model.NetworkingRule

	err := json.Unmarshal(value, &networkingRule)
	if err != nil {
		CBLogger.Error(err)
	}

	prettyJSON, _ := json.MarshalIndent(networkingRule, "", "\t")
	CBLogger.Trace("Pretty JSON")
	CBLogger.Trace(string(prettyJSON))

	if networkingRule.Contain(hostID) {
		CBNet.SetNetworkingRules(networkingRule)
		if !CBNet.IsRunning() {
			_, err := CBNet.StartCBNetworking(channel)
			if err != nil {
				CBLogger.Error(err)
			}
		}
	}
	CBLogger.Debug("End.........")
	mutex.Unlock()
}

func watchStatusTestSpecification(wg *sync.WaitGroup, etcdClient *clientv3.Client, cladnetID string, hostID string) {
	wg.Done()
	// Watch "/registry/cloud-adaptive-network/status/test-specification/{cladnet-id}" with version
	keyStatusTestSpecification := fmt.Sprint(etcdkey.StatusTestSpecification + "/" + cladnetID)
	CBLogger.Debugf("Start to watch \"%v\"", keyStatusTestSpecification)
	watchChan1 := etcdClient.Watch(context.Background(), keyStatusTestSpecification)
	for watchResponse := range watchChan1 {
		for _, event := range watchResponse.Events {
			CBLogger.Tracef("Watch - %s %q : %q", event.Type, event.Kv.Key, event.Kv.Value)

			// Get the trial count
			var testSpecification model.TestSpecification
			errUnmarshalEvalSpec := json.Unmarshal(event.Kv.Value, &testSpecification)
			if errUnmarshalEvalSpec != nil {
				CBLogger.Error(errUnmarshalEvalSpec)
			}
			trialCount := testSpecification.TrialCount

			// Check status of a CLADNet
			list := CBNet.NetworkingRules
			idx := list.GetIndexOfPublicIP(CBNet.MyPublicIP)

			// Perform a ping test to the host behind this host (in other words, behind idx)
			listLen := len(list.HostIPAddress)
			outSize := listLen - idx - 1
			var testwg sync.WaitGroup
			out := make([]model.InterHostNetworkStatus, outSize)

			for i := 0; i < len(out); i++ {
				testwg.Add(1)
				j := idx + i + 1
				out[i].SourceID = list.HostID[idx]
				out[i].SourceIP = list.HostIPAddress[idx]
				out[i].DestinationID = list.HostID[j]
				out[i].DestinationIP = list.HostIPAddress[j]
				go pingTest(&out[i], &testwg, trialCount)
			}
			testwg.Wait()

			// Gather the evaluation results
			var networkStatus model.NetworkStatus
			for i := 0; i < len(out); i++ {
				networkStatus.InterHostNetworkStatus = append(networkStatus.InterHostNetworkStatus, out[i])
			}

			if networkStatus.InterHostNetworkStatus == nil {
				networkStatus.InterHostNetworkStatus = make([]model.InterHostNetworkStatus, 0)
			}

			// Put the network status of the CLADNet to the etcd
			// Key: /registry/cloud-adaptive-network/status/information/{cladnet-id}/{host-id}
			keyStatusInformation := fmt.Sprint(etcdkey.StatusInformation + "/" + cladnetID + "/" + hostID)
			strNetworkStatus, _ := json.Marshal(networkStatus)
			_, err := etcdClient.Put(context.Background(), keyStatusInformation, string(strNetworkStatus))
			if err != nil {
				CBLogger.Error(err)
			}
		}
	}
	CBLogger.Debugf("End to watch \"%v\"", keyStatusTestSpecification)
}

func pingTest(outVal *model.InterHostNetworkStatus, wg *sync.WaitGroup, trialCount int) {
	CBLogger.Debug("Start.........")
	defer wg.Done()

	pinger, err := ping.NewPinger(outVal.DestinationIP)
	pinger.SetPrivileged(true)
	if err != nil {
		CBLogger.Error(err)
	}
	//pinger.OnRecv = func(pkt *ping.Packet) {
	//	fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
	//		pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	//}

	pinger.Count = trialCount
	pinger.Run() // blocks until finished

	stats := pinger.Statistics() // get send/receive/rtt stats
	outVal.MininumRTT = stats.MinRtt.Seconds()
	outVal.AverageRTT = stats.AvgRtt.Seconds()
	outVal.MaximumRTT = stats.MaxRtt.Seconds()
	outVal.StdDevRTT = stats.StdDevRtt.Seconds()
	outVal.PacketsReceive = stats.PacketsRecv
	outVal.PacketsLoss = stats.PacketsSent - stats.PacketsRecv
	outVal.BytesReceived = stats.PacketsRecv * 24

	CBLogger.Tracef("round-trip min/avg/max/stddev/dupl_recv = %v/%v/%v/%v/%v bytes",
		stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt, stats.PacketsRecv*24)
	CBLogger.Debug("End.........")
}

func main() {
	CBLogger.Debug("Start.........")

	var cladnetID = config.CBNetwork.CLADNetID
	var hostID string
	if config.CBNetwork.HostID == "" {
		name, err := os.Hostname()
		if err != nil {
			CBLogger.Error(err)
		}
		hostID = name
	} else {
		hostID = config.CBNetwork.HostID
	}

	// Wait for multiple goroutines to complete
	var wg sync.WaitGroup

	keyHostNetworkInformation := fmt.Sprint(etcdkey.HostNetworkInformation + "/" + cladnetID + "/" + hostID)
	keyNetworkingRule := fmt.Sprint(etcdkey.NetworkingRule + "/" + cladnetID)

	// etcd Section
	// Connect to the etcd cluster
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   config.ETCD.Endpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		CBLogger.Fatal(err)
	}

	defer func() {
		errClose := etcdClient.Close()
		if errClose != nil {
			CBLogger.Fatal("Can't close the etcd client", errClose)
		}
	}()

	CBLogger.Infoln("The etcdClient is connected.")

	channel = make(chan bool)

	// Create CBNetwork instance with port, which is tunneling port
	CBNet = cbnet.NewCBNetwork("cbnet0", 20000)

	// Start RunTunneling and blocked by channel until setting up the cb-network
	wg.Add(1)
	go CBNet.RunTunneling(&wg, channel)
	time.Sleep(3 * time.Second)

	if config.DemoApp.IsRun {
		// Start RunTunneling and blocked by channel until setting up the cb-network
		wg.Add(1)
		go app.PitcherAndCatcher(&wg, CBNet, channel)
	}

	wg.Add(1)
	go watchStatusTestSpecification(&wg, etcdClient, cladnetID, hostID)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		wg.Done()
		// Watch "/registry/cloud-adaptive-network/networking-rule/{cladnet-id}" with version
		CBLogger.Debugf("Start to watch \"%v\"", keyNetworkingRule)
		watchChan1 := etcdClient.Watch(context.Background(), keyNetworkingRule)
		for watchResponse := range watchChan1 {
			for _, event := range watchResponse.Events {
				CBLogger.Tracef("Watch - %s %q : %q", event.Type, event.Kv.Key, event.Kv.Value)
				decodeAndSetNetworkingRule(string(event.Kv.Key), event.Kv.Value, hostID)
			}
		}
		CBLogger.Debugf("End to watch \"%v\"", keyNetworkingRule)
	}(&wg)
	time.Sleep(3 * time.Second)

	// Try Compare-And-Swap (CAS) host-network-information by cladnetID and hostId
	CBLogger.Debug("Get the host network information")
	temp := CBNet.GetHostNetworkInformation()
	currentHostNetworkInformationBytes, _ := json.Marshal(temp)
	currentHostNetworkInformation := string(currentHostNetworkInformationBytes)
	CBLogger.Trace(currentHostNetworkInformation)

	CBLogger.Debug("Compare-And-Swap (CAS) the host network information")
	// NOTICE: "!=" doesn't work..... It might be a temporal issue.
	txnResp, err := etcdClient.Txn(context.Background()).
		If(clientv3.Compare(clientv3.Value(keyHostNetworkInformation), "=", currentHostNetworkInformation)).
		Then(clientv3.OpGet(keyNetworkingRule)).
		Else(clientv3.OpPut(keyHostNetworkInformation, currentHostNetworkInformation)).
		Commit()

	if err != nil {
		CBLogger.Error(err)
	}
	CBLogger.Tracef("Transaction Response: %v", txnResp)

	// The CAS would be succeeded if the prev host network information and current host network information are same.
	// Then the networking rule will be returned. (The above "watch" will not be performed.)
	// If not, the host tries to put the current host network information.
	if txnResp.Succeeded {
		// Set the networking rule to the host
		if len(txnResp.Responses[0].GetResponseRange().Kvs) != 0 {
			respKey := txnResp.Responses[0].GetResponseRange().Kvs[0].Key
			respValue := txnResp.Responses[0].GetResponseRange().Kvs[0].Value
			CBLogger.Tracef("The networking rule: %v", respValue)
			CBLogger.Debug("Set the networking rule")
			decodeAndSetNetworkingRule(string(respKey), respValue, hostID)
		}
	}

	// Waiting for all goroutines to finish
	CBLogger.Info("Waiting for all goroutines to finish")
	wg.Wait()

	CBLogger.Debug("End cb-network agent .........")
}
