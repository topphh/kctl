package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/topphh/kctl/pkg/metrics"
	"github.com/topphh/kctl/utils/display"
	"github.com/topphh/kctl/utils/format"
)

var (
	humanReadable bool
	sortBy        string
)

// topServiceCmd represents the 'top service' command
var topServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Show CPU and memory usage for services",
	Long: `The 'top service' command displays current CPU and memory usage 
for each Kubernetes service in the selected namespace.

Example usage:

  kctl top service --human-readable
  kctl top service --sort-by cpu
  kctl top service -H -s memory

This helps monitor service-level performance and identify resource hogs.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return topService()
	},
}

func topService() error {
	validSortKeys := map[string]bool{
		"pod":    true,
		"cpu":    true,
		"memory": true,
	}
	if sortBy != "" && !validSortKeys[sortBy] {
		return fmt.Errorf("invalid --sort-by value: %s (valid: name, pod, cpu, memory)", sortBy)
	}

	serviceInfo, err := metrics.GetKubeServiceTops()
	if err != nil {
		return fmt.Errorf("Unable to get kube metrics: %v", err)
	}

	switch sortBy {
	case "pod":
		sort.Slice(serviceInfo.Services, func(i, j int) bool {
			return serviceInfo.Services[i].PodCount > serviceInfo.Services[j].PodCount
		})
	case "cpu":
		sort.Slice(serviceInfo.Services, func(i, j int) bool {
			return serviceInfo.Services[i].Cpu > serviceInfo.Services[j].Cpu
		})
	case "memory":
		sort.Slice(serviceInfo.Services, func(i, j int) bool {
			return serviceInfo.Services[i].Memory > serviceInfo.Services[j].Memory
		})
	default:
		sort.Slice(serviceInfo.Services, func(i, j int) bool {
			return serviceInfo.Services[i].Name < serviceInfo.Services[j].Name
		})
	}

	fmt.Printf("TOTAL SERVICES: %v\tTOTAL POD:%v\n", len(serviceInfo.Services), serviceInfo.TotalPod)

	table := make([][]string, len(serviceInfo.Services)+1)
	table[0] = []string{"name", "pod", "cpu", "memory"}
	for i, service := range serviceInfo.Services {
		var podCount, cpu, memory string

		podCount = fmt.Sprint(service.PodCount)
		if humanReadable {
			cpu = format.CpuInfo(service.Cpu)
			memory = format.Bytes(service.Memory)
		} else {
			cpu = fmt.Sprintf("%vm", service.Cpu)
			memory = fmt.Sprintf("%v bytes", service.Memory)
		}

		table[i+1] = []string{service.Name, podCount, cpu, memory}
	}
	display.PrintTable(table)
	return nil
}

func init() {
	topCmd.AddCommand(topServiceCmd)

	topServiceCmd.Flags().BoolVarP(&humanReadable, "human-readable", "H", false, "Show values in human-readable units")
	topServiceCmd.Flags().StringVarP(&sortBy, "sort-by", "s", "", "Sort output by: name, pod, cpu, memory")
}
