/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/muqiuren/tsm/utils"
	"github.com/spf13/cobra"
	"log"
)

// goCmd represents the go command
var (
	goCmd = &cobra.Command{
		Use:     "go",
		Short:   `进入服务节点`,
		Long:    `go命令用于通过ssh账号密码登录方式进入服务节点`,
		Example: `tsm go [server id]`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			row := utils.Db.QueryRow("select `name`,`host`,`port`,`user`,`pass` from `main`.`servers` where id = ?", args[0])

			if row == nil {
				log.Fatalln("Can't find node, please enter node number")
			}
			var serverName, serverHost, serverPort, serverUser, serverPass string
			err := row.Scan(&serverName, &serverHost, &serverPort, &serverUser, &serverPass)
			utils.CheckErr(err)

			serverHost, _ = utils.DecryptByPrivate([]byte(serverHost), utils.PriKey)
			serverPort, _ = utils.DecryptByPrivate([]byte(serverPort), utils.PriKey)
			serverUser, _ = utils.DecryptByPrivate([]byte(serverUser), utils.PriKey)
			serverPass, _ = utils.DecryptByPrivate([]byte(serverPass), utils.PriKey)

			terminal := &utils.Terminal{
				Name: serverName,
				Host: serverHost,
				Port: serverPort,
				User: serverUser,
				Pass: serverPass,
			}

			terminal.Connect()
			terminal.NewSession()
		},
	}
)

func init() {
	rootCmd.AddCommand(goCmd)
}
