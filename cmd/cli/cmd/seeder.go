package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	s "github.com/ahmed-saleh/playbook/services"
)

type SeederPerformer interface {
	Name() string
	Seed() error
}

func PerformSeed(s SeederPerformer) error {
	return s.Seed()
}

type AdminSeeder struct {
	SeederName string
}

func (a *AdminSeeder) Seed() error {
	fmt.Println("seeding started Admin user")
	service := s.User{
		Display_name: "Admin",
		Email:        "admin@playbook.com",
		Password:     "changeMe",
	}

	return service.AddUser()
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
			fmt.Println("Please select a valid seeder, Valid options are: ", strings.Join(GetRegisteredSeeders(), ","))
			return
		}
		switch args[0] {
		case strings.Join(GetRegisteredSeeders(), ","):
			{
				seeder, _ := GetSeeder(args[0])
				err := PerformSeed(seeder)
				if err != nil {
					fmt.Println("there was an error ", err)
				}
				fmt.Println("seeding completed")
			}
		default:
			fmt.Println("Please select a valid seeder, Valid options are: ", strings.Join(GetRegisteredSeeders(), ","))
		}
	},
}
