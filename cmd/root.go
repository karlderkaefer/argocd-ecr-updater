package cmd

import (
	app "github.com/karlderkaefer/argocd-ecr-updater/pkg/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type EcrUpdaterConfig struct {
	KubeConfig string
	Namespace  string
	Interval   string // 11h0m0s
}

var cfg EcrUpdaterConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "argocd-ecr-updater",
	Short: "update password for AWS ECR helm repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := app.Run(cmd.Context(), cfg.KubeConfig, cfg.Namespace, cfg.Interval)
		if err != nil {
			return err

		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().String("kubeconfig", "", "kubernetes config file")
	rootCmd.Flags().String("namespace", "argocd", "kubernetes namespace")
	rootCmd.Flags().String("interval", "6h", "interval to refresh token")
	err := bindFlags()
	if err != nil {
		logrus.Errorln(err)
		return
	}
}

func bindFlags() error {
	bindings := []string{"kubeconfig", "namespace", "interval"}
	for _, binding := range bindings {
		err := viper.BindPFlag(binding, rootCmd.Flags().Lookup(binding))
		if err != nil {
			return err
		}
	}
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ARGOCD_ECR_UPDATER")
	cfg = EcrUpdaterConfig{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Errorln(err)
		return
	}
}
