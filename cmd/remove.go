/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/muqiuren/tsm/utils"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var (
	removeCmd = &cobra.Command{
		Use:     "remove",
		Short:   `移除服务节点`,
		Long:    `remove命令用于移除没有用到的服务节点`,
		Example: `tsm remove [server id]`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			stmt, err := utils.Db.Prepare("delete from `main`.`servers` where id = ?")
			utils.CheckErr(err)

			res, err := stmt.Exec(args[0])
			utils.CheckErr(err)

			affect, err := res.RowsAffected()
			utils.CheckErr(err)

			if affect > 0 {
				log.Println("移除节点成功")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
