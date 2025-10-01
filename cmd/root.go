/*
Copyright © 2024 Michael Lacore mclacore@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mclacore/passh/pkg/auth"
	"github.com/mclacore/passh/pkg/store"
	"github.com/spf13/cobra"
)

var rootCmdLong = `
 ______  ______   ______   ______   __  __    
/\  == \/\  __ \ /\  ___\ /\  ___\ /\ \_\ \   
\ \  _-/\ \  __ \\ \___  \\ \___  \\ \  __ \  
 \ \_\   \ \_\ \_\\/\_____\\/\_____\\ \_\ \_\ 
  \/_/    \/_/\/_/ \/_____/ \/_____/ \/_/\/_/ 
                                              

CLI-based password manager, because why not?
`

var dbPath string

var rootCmd = &cobra.Command{
	Use:               "passh",
	Short:             "CLI-based password manager",
	Long:              rootCmdLong,
	DisableAutoGenTag: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		_, err := store.Open(dbPath)
		if err != nil {
			return err
		}

		return auth.Migrate(store.DB())
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "auth:status",
	Short: "Show authentication metadata status",
	RunE: func(cmd *cobra.Command, args []string) error {
		ok, err := auth.Exists(cmd.Context(), store.DB())
		if err != nil {
			return err
		}
		if !ok {
			fmt.Println("auth: uninitialized")
			return nil
		}
		rec, err := auth.Get(cmd.Context(), store.DB())
		if err != nil {
			return err
		}
		fmt.Printf("auth: initialized v%d, salt=%dB, kdfParams=%dB\n",
			rec.Version, len(rec.Salt), len(rec.KDFParams))
		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(NewCmdPass())
	rootCmd.AddCommand(NewCmdLogin())
	rootCmd.AddCommand(NewCmdCollection())
	rootCmd.AddCommand(authStatusCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", "passh.db", "path to sqlite db")
}
