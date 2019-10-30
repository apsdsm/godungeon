module github.com/apsdsm/godungeon

go 1.12

replace github.com/apsdsm/canvas => ../canvas

require (
	github.com/apsdsm/canvas v0.0.0-00010101000000-000000000000
	github.com/apsdsm/imposter v0.0.0-20170924074901-8fbde9df20ca
	github.com/apsdsm/mapmaker v0.0.0-20190816092324-98f9bf7a56c1 // indirect
	github.com/gdamore/tcell v1.2.0
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20181102032728-5e5cf60278f6 // indirect
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/arch v0.0.0-20190815191158-8a70ba74b3a1 // indirect
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/tools v0.0.0-20190820205717-547ecf7b1ef1 // indirect
)
