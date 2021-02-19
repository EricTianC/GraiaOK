package environment

const (
	REPOURL = "iTXTech/mirai-console-loader"
	MCL_ZIP = "mcl.zip"
)

var mclse = &SimEnv{
	Name:     "mcl",
	BasePath: "mcl",
	ExecName: "mcl.jar",
}

func (es *EnvSpace) CheckMcl() {

}
