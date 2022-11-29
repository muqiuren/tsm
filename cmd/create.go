/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/muqiuren/tsm/utils"
	"log"
	"time"

	_ "github.com/muqiuren/tsm/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   `创建服务节点`,
	Long:    `tsm（Terminal Server Manager）终端服务管理器，你可以在这里存储调用进入终端.`,
	Example: `tsm create -s 127.0.0.1 -p 22 -u root -a admin123`,
	Run: func(cmd *cobra.Command, args []string) {
		stmt, err := utils.Db.Prepare("INSERT INTO `main`.`servers`(`name`, `host`, `port`, `user`, `pass`, `created_at`, `updated_at`) values(?,?,?,?,?,?,?)")
		utils.CheckErr(err)

		now := time.Now().Unix()
		host, _ := utils.EncryptByPublic(host, utils.PubKey)
		port, _ := utils.EncryptByPublic(port, utils.PubKey)
		user, _ := utils.EncryptByPublic(user, utils.PubKey)
		pass, _ := utils.EncryptByPublic(pass, utils.PubKey)

		_, err = stmt.Exec(name, host, port, user, pass, now, now)
		utils.CheckErr(err)

		log.Printf("添加[%s]节点成功\n", name)
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "节点名称 (required)")
	createCmd.Flags().StringVarP(&host, "host", "s", "", "节点host (required)")
	createCmd.Flags().StringVarP(&port, "port", "p", "22", "节点端口")
	createCmd.Flags().StringVarP(&user, "user", "u", "root", "节点用户")
	createCmd.Flags().StringVarP(&pass, "pass", "a", "", "节点密码 (required)")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("host")
	createCmd.MarkFlagRequired("pass")
	rootCmd.AddCommand(createCmd)
}
