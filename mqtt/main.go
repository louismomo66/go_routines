Inn0vexug@2024!
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var broker = "localhost"
// 	var port = 1883
// 	fmt.Println("Starting application")

// 	// Initialize MQTT service
// 	mqttService, err := NewMQTTClient(port, broker, "", "")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Topic to publish and subscribe to
// 	topic := "sensor_data/temperature"
// 	type publishMessage struct {
// 		Message string `json:"message"`
// 	}

// 	wg := sync.WaitGroup{}
// 	num := 10
// 	wg.Add(num)

// 	// Channel for receiving subscribed messages
// 	respDataChan := make(chan []byte)
// 	quit := make(chan bool, 1)

// 	// Start the subscription goroutine
// 	go func() {
// 		if err := mqttService.Subscribe(topic, respDataChan); err != nil {
// 			fmt.Printf("Error subscribing: %v\n", err)
// 			quit <- true
// 		}
// 	}()

// 	// Start the publishing goroutine
// 	go func() {
// 		for i := 0; i < num; i++ {
// 			fmt.Println("Publishing to topic")
// 			message, _ := json.Marshal(&publishMessage{Message: fmt.Sprintf("message %d", i)})
// 			if err := mqttService.Publish(topic, message); err != nil {
// 				fmt.Println(err)
// 			}
// 			time.Sleep(time.Second)
// 		}
// 		wg.Done()
// 	}()

// 	// Start a goroutine to monitor completion of the publishing process
// 	go func() {
// 		wg.Wait()
// 		quit <- true
// 		close(respDataChan)
// 	}()

//		// Main loop to receive responses or handle program exit
//		for {
//			select {
//			case respData, ok := <-respDataChan:
//				if !ok {
//					return
//				}
//				wg.Done()
//				fmt.Printf("Got response: %s\n", string(respData))
//			case <-quit:
//				fmt.Println("Program exited")
//				return
//			}
//		}
//	}
package main

import (
	"fmt"
)

func main() {
	var broker = "localhost"
	var port = 1883
	fmt.Println("Starting application")

	// Initialize MQTT service
	mqttService, err := NewMQTTClient(port, broker, "", "")
	if err != nil {
		panic(err)
	}

	// Topic to publish and subscribe to
	topic := "switch/message"
	// type publishMessage struct {
	// 	Message string `json:"message"`
	// }

	// Channel for receiving subscribed messages
	respDataChan := make(chan []byte)
	quit := make(chan bool, 1)

	// // Set up a WaitGroup for managing goroutines
	// var wg sync.WaitGroup

	// Start the subscription goroutine
	go func() {
		if err := mqttService.Subscribe(topic, respDataChan); err != nil {
			fmt.Printf("Error subscribing: %v\n", err)
			quit <- true
		}
	}()

	// Start a goroutine to handle user input and publish messages
	// go func() {
	// 	for {
	// 		var input string
	// 		fmt.Print("Enter message to publish (or type 'exit' to quit): ")
	// 		fmt.Scanf("%s", &input)

	// 		// Exit the loop if the user types 'exit'
	// 		if input == "exit" {
	// 			quit <- true
	// 			close(respDataChan)
	// 			break
	// 		}

	// 		// Publish the user's message to the MQTT topic
	// 		message, _ := json.Marshal(&publishMessage{Message: input})
	// 		if err := mqttService.Publish(topic, message); err != nil {
	// 			fmt.Println("Error publishing:", err)
	// 		}
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// Main loop to handle received messages and program exit
	for {
		select {
		case respData, ok := <-respDataChan:
			if !ok {
				return
			}
			fmt.Printf("Received message: %s\n", string(respData))
		case <-quit:
			fmt.Println("Program exiting...")
			return
		}
	}
}
