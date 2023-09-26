package domain

const (
	TablespaceType = iota + 1
	DirectoryType
	DbLinkType
	UserType
	ContextType
	SequencesType
	TypeType
	TableType
	TableFKType
	MLogType
	MViewType
	PackageType
	ViewType
	TriggerType
	VTaskType
	RowType
	RoleType
	FunctionType
	VClogType
	SynonymType
	ScriptsBeforeType
	ScriptsAfterType
	ScriptsMigrationType
)

type ChangelogItem struct {
	Name       string
	Server     string
	Schema     string
	Type       byte
	Script     string
	CommitHash string
}
