package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type SeederPerformer interface {
	Name() string
	Seed() bool
}

func PerformSeed(s SeederPerformer) bool {
	return s.Seed()
}

type AdminSeeder struct {
	SeederName string
}

func (a *AdminSeeder) Seed() bool {
	fmt.Println("seeding started for ", a.SeederName)
	return true
}

func (a *AdminSeeder) Name() string {
	return a.SeederName
}

var RegisteredSeeders = []SeederPerformer{}

func GetRegisteredSeeders() []string {
	args := []string{}
	for _, v := range RegisteredSeeders {
		args = append(args, v.Name())
	}
	return args
}

func GetSeeder(n string) (SeederPerformer, error) {
	var err error
	for _, v := range RegisteredSeeders {
		if v.Name() == n {
			return v, nil
		}
	}
	return &AdminSeeder{}, err //empty admin seeder
}

func init() {
	RegisteredSeeders = append(RegisteredSeeders, &AdminSeeder{"admin_seeder"})
	rootCmd.AddCommand(seeder)
}

var seeder = &cobra.Command{
	Use:       "seed",
	Short:     "Run seeders",
	ValidArgs: GetRegisteredSeeders(),
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("please provide a valid arguments")
			return
		}
		switch args[0] {
		case strings.Join(GetRegisteredSeeders(), ","):
			{
				seeder, _ := GetSeeder(args[0])
				PerformSeed(seeder)
			}
		default:
			fmt.Println("Please select a valid seeder, Valid options are: ", strings.Join(GetRegisteredSeeders(), ","))
		}
	},
}
