package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var eatWG sync.WaitGroup

type ChopStick struct {
	sync.Mutex
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopStick
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func callHost(philos []*Philo, seatsNum int, eatTimes int) {
	eatWG.Add(1)
	go host(philos, seatsNum, eatTimes)
	eatWG.Wait()
}

func host(philos []*Philo, seatsNum int, eatTimes int) {
	defer eatWG.Done()
	eating := make(chan *Philo, 2)

	for i := 0; i < seatsNum; i++ {
		eatWG.Add(1)
		go philos[i].eat(eatTimes, eating)
	}
}

func genTwoRandInd() (int, int) {
	randI1 := 0 // Random index of first chop stick to pick
	randI2 := 0 // Random index of second chop stick to pick
	randI1 = rand.Intn(2)
	if randI1 == 0 {
		randI2 = 1
	} else {
		randI2 = 0
	}
	return randI1, randI2
}

func (philo *Philo) eat(eatTimes int, eating chan *Philo) {
	defer eatWG.Done()
	sliceCS := make([]*ChopStick, 2)
	sliceCS[0] = philo.leftCS
	sliceCS[1] = philo.rightCS
	randI1, randI2 := genTwoRandInd() // Random index of first and second chop sticks to pick

	for i := 0; i < eatTimes; i++ {
		select {
		case eating <- philo:
			//fmt.Printf("Host allows philosopher with index - %d to eat\n", philo.id)
			//randomPause(3)
			sliceCS[randI1].Lock()
			sliceCS[randI2].Lock()

			fmt.Printf("Starting to eat philospher with index - %d\n", philo.id)
			randomPause(3)
			fmt.Printf("Finishing to eat philospher with index - %d\n", philo.id)
			<-eating

			sliceCS[randI2].Unlock()
			sliceCS[randI1].Unlock()
		default:
			//fmt.Printf("Host denied philosopher with index - %d to eat\n", philo.id)
			//randomPause(3)
			i--
		}
	}
}

func main() {
	seatsNum, err := strconv.Atoi(os.Args[1]) // Number of philosophers and chop sticks as well
	if err != nil {
		return
	}
	eatTimes, err := strconv.Atoi(os.Args[2]) // Number of times every philosopher eats
	if err != nil {
		return
	}
	// Initialize chop sticks
	chopSticks := make([]*ChopStick, seatsNum)
	for i := 0; i < seatsNum; i++ {
		chopSticks[i] = new(ChopStick)
	}
	// Initialize chop philosophers
	philos := make([]*Philo, seatsNum)
	for i := 0; i < seatsNum; i++ {
		philos[i] = &Philo{i + 1, chopSticks[i], chopSticks[(i+1)%seatsNum]}
	}

	//eatWG.Add(1)
	//go host(philos, seatsNum, eatTimes)
	//eatWG.Wait()
	callHost(philos, seatsNum, eatTimes)
}
