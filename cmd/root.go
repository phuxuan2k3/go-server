/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "darius",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		go startGRPC()
		startGateway()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Thiết lập để tự động đọc biến môi trường
	viper.AutomaticEnv()

	// // Cấu hình Viper để đọc tệp YAML
	// viper.SetConfigName("config") // Tên tệp (không bao gồm phần mở rộng)
	// viper.SetConfigType("yaml")   // Loại tệp
	// viper.AddConfigPath(".")      // Thư mục chứa tệp cấu hình

	// // Đọc tệp cấu hình
	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(fmt.Errorf("Lỗi khi đọc tệp cấu hình: %w", err))
	// }
}
