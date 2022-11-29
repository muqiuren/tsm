/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/muqiuren/tsm/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   `列出服务节点`,
	Long:    `list命令用于列出可用的服务节点`,
	Example: `tsm list`,
	Run: func(cmd *cobra.Command, args []string) {
		var offset int32 = 0
		if page > 1 {
			offset = (page - 1) * perPage
		}

		rows, err := utils.Db.Query(fmt.Sprintf("SELECT `id`, `name`,`host` FROM `main`.`servers` LIMIT %v,%v", offset, perPage))
		utils.CheckErr(err)

		headers := []string{"ID", "公网IP", "节点名称"}
		// 输出表格
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers)
		table.SetBorder(true)
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgYellowColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgYellowColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgYellowColor, tablewriter.Bold},
		)
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
			tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
			tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
		)

		for rows.Next() {
			var serverId string
			var serverName string
			var serverHost string
			err = rows.Scan(&serverId, &serverName, &serverHost)
			utils.CheckErr(err)
			serverHost, _ = utils.DecryptByPrivate([]byte(serverHost), utils.PriKey)
			if isShowRaw {
				serverHost = ipHidden(serverHost)
			}
			table.Append([]string{serverId, serverHost, serverName})
		}

		table.Render()
	},
}

func init() {
	listCmd.Flags().Int32VarP(&page, "page", "n", 1, "页码")
	listCmd.Flags().Int32VarP(&perPage, "per_page", "c", 20, "每页数量")
	listCmd.Flags().BoolVar(&isShowRaw, "r", true, "显示原始信息")
	rootCmd.AddCommand(listCmd)
}

// ipHidden 隐藏ip
func ipHidden(ip string) string {
	frags := strings.Split(ip, ".")
	return fmt.Sprintf("*.*.*.%s", frags[len(frags)-1])
}
