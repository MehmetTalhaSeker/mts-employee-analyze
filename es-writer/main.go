package main

func main() {
	esc := newElasticSearchClient()

	k := newKafkaService(esc)
	k.startProcessing()
}
