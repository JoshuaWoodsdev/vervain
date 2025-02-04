package cmd

func Run() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
