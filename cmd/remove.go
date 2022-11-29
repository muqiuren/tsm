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
		Run: func(cmd *cobra.Command, args []string) {
			stmt, err := utils.Db.Prepare("delete from `main`.`servers` where id = ?")
			utils.CheckErr(err)

			res, err := stmt.Exec(name)
			utils.CheckErr(err)

			affect, err := res.RowsAffected()
			utils.CheckErr(err)

			if affect > 0 {
				log.Printf("移除节点[%s]成功\n", name)
			}
		},
	}
)

func init() {
	removeCmd.Flags().Int32VarP(&id, "id", "i", 0, "节点id")
	removeCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(removeCmd)
}
