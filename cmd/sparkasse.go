package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(spkCmd)
	spkCmd.AddCommand(debitCmd)
}

var spkCmd = &cobra.Command{
	Use:   "spk",
	Short: "commands for sparkasse bank operations",
}

var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "process sparkasse debit account statement",
	Run:   processSpkDebitStatement,
}

func processSpkDebitStatement(cmd *cobra.Command, args []string) {

}
