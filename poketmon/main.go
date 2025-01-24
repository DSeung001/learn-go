package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"pokemon.com/factories"
	"pokemon.com/models"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	pokemonResitry := factories.NewPokemonFactoryRegistry()
	pokemonResitry.PokemonFactoryRegister()

	fmt.Print("Enter your name: ")
	trainerName, _ := reader.ReadString('\n')
	trainerName = strings.TrimSpace(trainerName)

	fmt.Print("Enter your rival's name: ")
	rivalName, _ := reader.ReadString('\n')
	rivalName = strings.TrimSpace(rivalName)

	// 구조체 메모리 직접 참조
	trainer := &models.Trainer{Name: trainerName}
	rival := &models.Trainer{Name: rivalName}

	CatchAllEntryPokemon(pokemonResitry, trainer)
	CatchAllEntryPokemon(pokemonResitry, rival)

	trainer.PrintPokemons()
	rival.PrintPokemons()
}

func CatchAllEntryPokemon(pokemonResitry *factories.PokemonFactoryRegistry, trainer *models.Trainer) {
	for i := 0; i <= models.MaxEntryCount; i++ {
		pokemonFactory, err := pokemonResitry.GetRandomPokemonFactory()
		pokemon := pokemonFactory.NewPokemon(rand.Intn(50) + 1)

		if err != nil {
			fmt.Println("Error", err)
			break
		}
		trainer.CatchPokemon(pokemon)
	}
}
