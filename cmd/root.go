/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	id        int32
	name      string
	host      string
	port      string
	user      string
	pass      string
	isShowRaw bool
	page      int32
	perPage   int32
	rootCmd   = &cobra.Command{
		Use:   "tsm [create|list|remove|go]",
		Short: `Terminal Server Manager`,
		Long:  `tsm（Terminal Server Manager）终端服务管理器，你可以在这里存储调用进入终端.`,
		Example: `
1.创建节点
	tsm create -n server1 -s 127.0.0.1 -p 22 -u root -a admin123
2.列出节点
	tsm list
3.进入节点
	tsm exec server1
4.移除节点
	tsm remove server1`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Version = "0.1.0"
	rootCmd.SetVersionTemplate(`当前版本: {{.Version}}`)
	rootCmd.SetUsageTemplate(`使用:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

别名:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

示例:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

命令:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

标志:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
{{end}}`)
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help [command]",
		Short: "显示命令的具体用法",
		Long:  `你可以使用tsm help [command] 来查询改命令的具体用法`,
		ValidArgsFunction: func(c *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var completions []string
			cmd, _, e := c.Root().Find(args)
			if e != nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			if cmd == nil {
				// Root help command.
				cmd = c.Root()
			}
			for _, subCmd := range cmd.Commands() {
				if subCmd.IsAvailableCommand() {
					if strings.HasPrefix(subCmd.Name(), toComplete) {
						completions = append(completions, fmt.Sprintf("%s\t%s", subCmd.Name(), subCmd.Short))
					}
				}
			}
			return completions, cobra.ShellCompDirectiveNoFileComp
		},
		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Printf("Unknown help topic %#q\n", args)
				cobra.CheckErr(c.Root().Usage())
			} else {
				cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
				cobra.CheckErr(cmd.Help())
			}
		},
	})
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
